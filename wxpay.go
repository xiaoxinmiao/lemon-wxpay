package wxpay

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/relax-space/go-kitt/random"

	"github.com/relax-space/go-kit/data"
	"github.com/relax-space/go-kit/httpreq"

	"github.com/relax-space/go-kit/base"
	"github.com/relax-space/go-kit/sign"
)

func LoopQuery(reqDto *ReqQueryDto, custDto *ReqCustomerDto, limit, interval int) (statusCode int, code string, queryResult map[string]interface{}, err error) {
	count := limit / interval
	waitTime := time.Duration(interval) * time.Second //2s
	for index := 0; index < count; index++ {
		statusCode, code, queryResult, err = Query(reqDto, custDto)
		if err == nil { // 1. request wechat query api success
			tradeStatusObj, ok := queryResult["trade_state"]
			if !ok { //1.1 wechat query api response result is exception
				time.Sleep(waitTime)
				continue
			}
			tradeStatus := tradeStatusObj.(string)
			switch {
			case tradeStatus == "SUCCESS": //1.2 pay success
				code = SUC
				return
			case tradeStatus == "CLOSED" || tradeStatus == "REFUND" || tradeStatus == "REVOKED" || tradeStatus == "NOTPAY" || tradeStatus == "PAYERROR":
				err = fmt.Errorf("pay status is %v", tradeStatus)
				code = E03
				return //1.3 pay failure
			case tradeStatus == "USERPAYING": //1.4 pay unknown
				if index < count {
					time.Sleep(waitTime)
					continue
				}
			}
		} else { //2. request wechat query api failure
			time.Sleep(waitTime)
			continue
		}
	}
	err = errors.New(MESSAGE_OVERTIME)
	code = E03
	return
}

func Pay(reqDto *ReqPayDto, custDto *ReqCustomerDto) (statusCode int, code string, result map[string]interface{}, err error) {
	wxPayData := BuildCommonparam(reqDto.ReqBaseDto)

	outTradeNo := reqDto.OutTradeNo
	if len(outTradeNo) == 0 {
		outTradeNo = random.Uuid(PRE_OUTTRADENO)
	}
	SetValue(wxPayData, "out_trade_no", outTradeNo)
	SetValue(wxPayData, "body", reqDto.Body)
	SetValue(wxPayData, "total_fee", reqDto.TotalFee)
	SetValue(wxPayData, "auth_code", reqDto.AuthCode)
	SetValue(wxPayData, "device_info", reqDto.DeviceInfo)

	SetValue(wxPayData, "detail", reqDto.Detail)
	SetValue(wxPayData, "attach", reqDto.Attach)
	SetValue(wxPayData, "fee_type", reqDto.FeeType)
	SetValue(wxPayData, "goods_tag", reqDto.GoodsTag)
	SetValue(wxPayData, "limit_pay", reqDto.LimitPay)
	if len(strings.TrimSpace(reqDto.SpbillCreateIp)) == 0 {
		SetValue(wxPayData, "spbill_create_ip", "8.8.8.8")
	} else {
		SetValue(wxPayData, "spbill_create_ip", reqDto.SpbillCreateIp)
	}
	signStr := base.JoinMapObject(wxPayData.DataAttr)
	SetValue(wxPayData, "sign", sign.MakeMd5Sign(signStr, custDto.Key))
	resp, body, err := httpreq.NewPost(URLPAY, []byte(wxPayData.ToXml()),
		&httpreq.Header{ContentType: httpreq.MIMEApplicationXMLCharsetUTF8}, nil)
	if err != nil {
		code = E01
		return
	}
	statusCode = resp.StatusCode
	code, result, err = RespParse(body, custDto.Key)
	if err != nil {
		if err.Error() == MESSAGE_PAYING {
			result = map[string]interface{}{"out_trade_no": outTradeNo}
		}
		return
	}
	return
}
func Query(reqDto *ReqQueryDto, custDto *ReqCustomerDto) (statusCode int, code string, result map[string]interface{}, err error) {
	wxPayData := BuildCommonparam(reqDto.ReqBaseDto)

	SetValue(wxPayData, "transaction_id", reqDto.TransactionId)
	SetValue(wxPayData, "out_trade_no", reqDto.OutTradeNo)

	signStr := base.JoinMapObject(wxPayData.DataAttr)
	SetValue(wxPayData, "sign", sign.MakeMd5Sign(signStr, custDto.Key))
	resp, body, err := httpreq.NewPost(URLQUERY, []byte(wxPayData.ToXml()),
		&httpreq.Header{ContentType: httpreq.MIMEApplicationXMLCharsetUTF8}, nil)
	if err != nil {
		code = E01
		return
	}
	statusCode = resp.StatusCode
	code, result, err = RespParse(body, custDto.Key)
	return
}

func Refund(reqDto *ReqRefundDto, custDto *ReqCustomerDto) (statusCode int, code string, result map[string]interface{}, err error) {
	wxPayData := BuildCommonparam(reqDto.ReqBaseDto)
	//query
	queryDto := ReqQueryDto{ReqBaseDto: reqDto.ReqBaseDto, OutTradeNo: reqDto.OutTradeNo}
	statusCode, code, queryResult, err := Query(&queryDto, custDto)
	if err != nil {
		err = errors.New("refund failure,please re-try")
		code = E01
		return
	}
	totalFee, ok := queryResult["total_fee"]
	if !ok {
		err = errors.New("refund failure,please re-try")
		code = E01
		return
	}

	outRefundNo := reqDto.OutRefundNo
	if len(outRefundNo) == 0 {
		outRefundNo = random.Uuid(PRE_OUTREFUNDNO)
	}
	SetValue(wxPayData, "out_refund_no", outRefundNo)
	SetValue(wxPayData, "transaction_id", reqDto.TransactionId)
	SetValue(wxPayData, "total_fee", totalFee)
	SetValue(wxPayData, "refund_fee", reqDto.RefundFee)

	SetValue(wxPayData, "refund_fee_type", reqDto.RefundFeeType)
	SetValue(wxPayData, "refund_account", reqDto.RefundAccount)
	SetValue(wxPayData, "out_trade_no", reqDto.OutTradeNo)
	signStr := base.JoinMapObject(wxPayData.DataAttr)
	SetValue(wxPayData, "sign", sign.MakeMd5Sign(signStr, custDto.Key))

	tr, err := httpreq.CertTransport(custDto.CertPathName, custDto.CertPathKey, custDto.RootCa)
	if err != nil {
		err = errors.New("Certificate verification failed")
		code = E02
		return
	}
	resp, body, err := httpreq.NewPost(URLREFUND, []byte(wxPayData.ToXml()),
		&httpreq.Header{ContentType: httpreq.MIMEApplicationXMLCharsetUTF8}, tr)
	if err != nil {
		code = E01
		return
	}
	statusCode = resp.StatusCode
	code, result, err = RespParse(body, custDto.Key)
	return
}

func Reverse(reqDto *ReqReverseDto, custDto *ReqCustomerDto, count int, interval int) (statusCode int, code string, result map[string]interface{}, err error) {
	if count <= 0 {
		err = errors.New("reverse failure")
		code = E01
		return
	}
	wxPayData := BuildCommonparam(reqDto.ReqBaseDto)

	SetValue(wxPayData, "transaction_id", reqDto.TransactionId)
	SetValue(wxPayData, "out_trade_no", reqDto.OutTradeNo)

	signStr := base.JoinMapObject(wxPayData.DataAttr)
	SetValue(wxPayData, "sign", sign.MakeMd5Sign(signStr, custDto.Key))
	tr, err := httpreq.CertTransport(custDto.CertPathName, custDto.CertPathKey, custDto.RootCa)
	if err != nil {
		err = errors.New("Certificate verification failed")
		code = E02
		return
	}
	resp, body, err := httpreq.NewPost(URLREVERSE, []byte(wxPayData.ToXml()),
		&httpreq.Header{ContentType: httpreq.MIMEApplicationXMLCharsetUTF8}, tr)
	if err != nil {
		code = E01
		return
	}
	statusCode = resp.StatusCode
	code, result, err = RespParse(body, custDto.Key)
	if err != nil {
		return
	}
	recallObj, ok := result["recall"]
	if !ok {
		err = errors.New("reverse failure,please re-try")
		code = E01
		return
	}

	recall := recallObj.(string)
	if recall == "Y" {
		time.Sleep(time.Duration(interval) * time.Second) //10s
		count = count - 1
		return Reverse(reqDto, custDto, count, interval)
	} else if recall != "N" {
		err = errors.New("reverse failure,please re-try")
		code = E01
		return
	}
	return
}

func RefundQuery(reqDto *ReqRefundQueryDto, custDto *ReqCustomerDto) (statusCode int, code string, result map[string]interface{}, err error) {
	wxPayData := BuildCommonparam(reqDto.ReqBaseDto)

	SetValue(wxPayData, "transaction_id", reqDto.TransactionId)
	SetValue(wxPayData, "out_trade_no", reqDto.OutTradeNo)
	SetValue(wxPayData, "out_refund_no", reqDto.OutRefundNo)
	SetValue(wxPayData, "refund_id", reqDto.RefundId)
	SetValue(wxPayData, "offset", reqDto.Offset)

	signStr := base.JoinMapObject(wxPayData.DataAttr)
	SetValue(wxPayData, "sign", sign.MakeMd5Sign(signStr, custDto.Key))
	resp, body, err := httpreq.NewPost(URLREFUNDQUERY, []byte(wxPayData.ToXml()),
		&httpreq.Header{ContentType: httpreq.MIMEApplicationXMLCharsetUTF8}, nil)
	if err != nil {
		code = E01
		return
	}
	statusCode = resp.StatusCode
	code, result, err = RespParse(body, custDto.Key)
	return
}

func Prepay(reqDto *ReqPrepayDto, custDto *ReqCustomerDto) (statusCode int, code string, result map[string]interface{}, err error) {
	wxPayData := BuildCommonparam(reqDto.ReqBaseDto)
	if len(reqDto.OutTradeNo) == 0 {
		SetValue(wxPayData, "out_trade_no", random.Uuid(PRE_PREOUTTRADENO))
	} else {
		SetValue(wxPayData, "out_trade_no", reqDto.OutTradeNo)
	}

	if len(reqDto.SpbillCreateIp) == 0 {
		SetValue(wxPayData, "spbill_create_ip", reqDto.SpbillCreateIp)
	} else {
		SetValue(wxPayData, "spbill_create_ip", "8.8.8.8")
	}

	SetValue(wxPayData, "sign_type", reqDto.SignType)
	SetValue(wxPayData, "body", reqDto.Body)
	SetValue(wxPayData, "detail", reqDto.Detail)
	SetValue(wxPayData, "attach", reqDto.Attach)
	SetValue(wxPayData, "device_info", reqDto.DeviceInfo)
	SetValue(wxPayData, "fee_type", reqDto.FeeType)

	SetValue(wxPayData, "total_fee", reqDto.TotalFee)

	SetValue(wxPayData, "time_start", reqDto.TimeStart)
	SetValue(wxPayData, "time_expire", reqDto.TimeExpire)
	SetValue(wxPayData, "goods_tag", reqDto.GoodsTag)

	if len(reqDto.NotifyUrl) != 0 {
		SetValue(wxPayData, "notify_url", reqDto.NotifyUrl)
	} else {
		SetValue(wxPayData, "notify_url", custDto.UnifiedNotifyUrl)
	}
	SetValue(wxPayData, "trade_type", reqDto.TradeType)
	SetValue(wxPayData, "product_id", reqDto.ProductId)
	SetValue(wxPayData, "limit_pay", reqDto.LimitPay)
	SetValue(wxPayData, "openid", reqDto.OpenId)
	SetValue(wxPayData, "sub_openid", reqDto.SubOpenId)

	SetValue(wxPayData, "scene_info", reqDto.SceneInfo)
	signStr := base.JoinMapObject(wxPayData.DataAttr)
	SetValue(wxPayData, "sign", sign.MakeMd5Sign(signStr, custDto.Key))
	resp, body, err := httpreq.NewPost(URLPREPAY, []byte(wxPayData.ToXml()), &httpreq.Header{ContentType: httpreq.MIMEApplicationXMLCharsetUTF8}, nil)
	if err != nil {
		code = E01
		return
	}
	statusCode = resp.StatusCode
	code, result, err = RespParse(body, custDto.Key)
	return
}

func Bill(reqDto *ReqBillDto, custDto *ReqCustomerDto) (statusCode int, code string, result []byte, err error) {
	wxPayData := BuildCommonparam(reqDto.ReqBaseDto)

	SetValue(wxPayData, "bill_date", reqDto.BillDate)
	SetValue(wxPayData, "bill_type", reqDto.BillType)
	SetValue(wxPayData, "tar_type", reqDto.TarType)

	signStr := base.JoinMapObject(wxPayData.DataAttr)
	SetValue(wxPayData, "sign", sign.MakeMd5Sign(signStr, custDto.Key))
	resp, body, err := httpreq.NewPost(URLBILL, []byte(wxPayData.ToXml()),
		&httpreq.Header{ContentType: httpreq.MIMEApplicationXMLCharsetUTF8}, nil)
	if err != nil {
		code = E01
		return
	}
	if resp == nil {
		code = E01
		return
	}

	statusCode = resp.StatusCode
	contentType := resp.Header.Get("Content-Type")
	if contentType == "application/x-gzip" { //success
		b, err := GzipDecode(body)
		if err != nil {
			err = errors.New("unpacking file error")
		}
		result = b
	} else if contentType == "text/plain; charset=utf-8" { //failure
		data := data.New()
		err = data.FromXml(string(body))
		if err != nil {
			code = E04
			return
		}
		if !data.IsSet("return_code") {
			err = errors.New("return_code is missing")
			code = E04
			return
		}
		if strings.ToUpper(data.GetValue("return_code").(string)) != "SUCCESS" {
			err = fmt.Errorf("%v:%v", MESSAGE_WECHAT, data.GetValue("return_msg"))
			code = E03
			return
		}
		err = errors.New("return_code is wrong")
		code = E03
		return
	}
	return
}

//sub_notify_url maybe exist in attach,
//if sub_notify_url exist,then redirect to sub_notify_url
func Notify(xmlBody string) (result NotifyDto, err error) {
	err = xml.Unmarshal([]byte(xmlBody), &result)
	if err != nil {
		err = fmt.Errorf("%v:%v", MESSAGE_WECHAT, err)
		return
	}

	if len(result.Attach) == 0 {
		return
	} else {
		var attachObj struct {
			SubNotifyUrl string `json:"sub_notify_url"`
		}
		err = json.Unmarshal([]byte(result.Attach), &attachObj)
		if err != nil {
			return
		}

		if len(attachObj.SubNotifyUrl) != 0 {
			_, err = httpreq.POST("", attachObj.SubNotifyUrl, result, nil)
		}
	}
	return
}
