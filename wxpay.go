package wxpay

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/relax-space/go-kit/random"

	"github.com/relax-space/go-kit/httpreq"

	"github.com/relax-space/go-kit/base"
	"github.com/relax-space/go-kit/sign"
)

func LoopQuery(reqDto reqQueryDto, custDto reqCustomerDto, limit, interval int) (queryResult map[string]interface{}, err error) {
	count := limit / interval
	waitTime := time.Duration(interval) * time.Second //2s
	for index := 0; index < count; index++ {
		queryResult, err = Query(reqDto, custDto)
		fmt.Printf("%+v", queryResult)
		if err == nil { // 1. query success
			tradeStatusObj, ok := queryResult["trade_state"]
			if !ok {
				continue
			}
			tradeStatus := tradeStatusObj.(string)
			fmt.Println(tradeStatus)
			switch {
			case tradeStatus == "SUCCESS":
				return
			case tradeStatus == "USERPAYING":
				if index < count {
					time.Sleep(waitTime)
					continue
				}
			case tradeStatus == "CLOSED" || tradeStatus == "REFUND" || tradeStatus == "REVOKED" || tradeStatus == "NOTPAY" || tradeStatus == "PAYERROR":
				err = errors.New("wechat pay failure")
				return
			}
		}
	}
	err = fmt.Errorf("wechat pay overtime:%v(s)", limit)
	return
}

func Pay(reqDto reqPayDto, custDto reqCustomerDto) (result map[string]interface{}, err error) {
	wxPayData := BuildCommonparam(reqDto.reqBaseDto)

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
	_, body, err := httpreq.NewPost(URLPAY, []byte(wxPayData.ToXml()),
		&httpreq.Header{ContentType: httpreq.MIMEApplicationXMLCharsetUTF8}, nil)
	if err != nil {
		err = fmt.Errorf("%v:%v", MESSAGE_WECHAT, err)
		return
	}
	result, err = RespParse(body, custDto.Key)
	if err != nil {
		if err.Error() == MESSAGE_PAYING {
			result = map[string]interface{}{"out_trade_no": outTradeNo}
		}
		return
	}
	return
}
func Query(reqDto reqQueryDto, custDto reqCustomerDto) (result map[string]interface{}, err error) {
	wxPayData := BuildCommonparam(reqDto.reqBaseDto)

	SetValue(wxPayData, "transaction_id", reqDto.TransactionId)
	SetValue(wxPayData, "out_trade_no", reqDto.OutTradeNo)

	signStr := base.JoinMapObject(wxPayData.DataAttr)
	SetValue(wxPayData, "sign", sign.MakeMd5Sign(signStr, custDto.Key))
	_, body, err := httpreq.NewPost(URLQUERY, []byte(wxPayData.ToXml()),
		&httpreq.Header{ContentType: httpreq.MIMEApplicationXMLCharsetUTF8}, nil)
	if err != nil {
		return
	}
	result, err = RespParse(body, custDto.Key)
	if err != nil {
		return
	}
	return
}

func Refund(reqDto reqRefundDto, custDto reqCustomerDto) (result map[string]interface{}, err error) {
	wxPayData := BuildCommonparam(reqDto.reqBaseDto)
	//query
	queryDto := reqQueryDto{reqBaseDto: reqDto.reqBaseDto, OutTradeNo: reqDto.OutTradeNo}
	queryResult, err := Query(queryDto, custDto)
	if err != nil {
		err = errors.New("refund failure")
		return
	}
	totalFee, ok := queryResult["total_fee"]
	if !ok {
		err = errors.New("refund failure")
		return
	}

	outRefundNo := reqDto.OutRefundNo
	if len(outRefundNo) == 0 {
		outRefundNo = random.Uuid(PRE_OUTREFUNDNO)
	}
	SetValue(wxPayData, "out_refund_no", outRefundNo)
	SetValue(wxPayData, "device_info", reqDto.DeviceInfo)
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
		return
	}
	_, body, err := httpreq.NewPost(URLREFUND, []byte(wxPayData.ToXml()), &httpreq.Header{ContentType: httpreq.MIMEApplicationXMLCharsetUTF8}, tr)
	if err != nil {
		err = fmt.Errorf("%v:%v", MESSAGE_WECHAT, err)
		return
	}
	result, err = RespParse(body, custDto.Key)
	return
}
