package wxpay

import (
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/relax-space/go-kit/base"
	"github.com/relax-space/go-kit/data"
	"github.com/relax-space/go-kit/sign"
	"github.com/relax-space/go-kitt/random"
)

//wechat const
const (
	SYSTEMERROR   = "SYSTEMERROR"
	ORDERNOTEXIST = "ORDERNOTEXIST"
	BANKERROR     = "BANKERROR"
	USERPAYING    = "USERPAYING"
)

//random prefixes const
const (
	PRE_OUTTRADENO    = "14"
	PRE_OUTREFUNDNO   = "15"
	PRE_PREOUTTRADENO = "16"
)

//message const
const (
	MESSAGE_PAYING   = "MESSAGE_PAYING"
	MESSAGE_WECHAT   = "MESSAGE_WECHAT"
	MESSAGE_OVERTIME = "MESSAGE_OVERTIME"
)

const (
	PREPAY_TYPE_JSAPI  = "JSAPI"
	PREPAY_TYPE_NATIVE = "NATIVE"
	PREPAY_TYPE_APP    = "APP"
)

const (
	URLPAY         = "https://api.mch.weixin.qq.com/pay/micropay"
	URLREVERSE     = "https://api.mch.weixin.qq.com/secapi/pay/reverse"
	URLREFUND      = "https://api.mch.weixin.qq.com/secapi/pay/refund"
	URLQUERY       = "https://api.mch.weixin.qq.com/pay/orderquery"
	URLPREPAY      = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	URLREFUNDQUERY = "https://api.mch.weixin.qq.com/pay/refundquery"
	URLBILL        = "https://api.mch.weixin.qq.com/pay/downloadbill"
)

const (
	SUC = "SUC" //success
	E01 = "E01" //system error,can re-try
	E02 = "E02" //bad request format
	E03 = "E03" //message from wechat
	E04 = "E04" //bad response
)

func BuildCommonparam(baseDto *ReqBaseDto) *data.Data {
	data := data.New()
	SetValue(data, "appid", baseDto.AppId)
	SetValue(data, "sub_appid", baseDto.SubAppId)
	SetValue(data, "mch_id", baseDto.MchId)
	SetValue(data, "sub_mch_id", baseDto.SubMchId)
	SetValue(data, "nonce_str", random.Uuid(""))
	return data
}

func SetValue(data *data.Data, key string, value interface{}) {
	if len(base.ToString(value)) != 0 {
		data.SetValue(key, value)
	}
}

func RespParse(bodyByte []byte, key string) (code string, result map[string]interface{}, err error) {
	data := data.New()
	err = data.FromXml(string(bodyByte))
	if err != nil {
		code = E04
		return
	}
	if !data.IsSet("return_code") {
		err = errors.New("wechat response:return_code is missing")
		code = E04
		return
	}

	if strings.ToUpper(data.GetValue("return_code").(string)) != "SUCCESS" {
		err = fmt.Errorf("%v:%v", MESSAGE_WECHAT, data.GetValue("return_msg"))
		code = E03
		return
	}
	if !data.IsSet("result_code") {
		err = errors.New("wechat response:result_code is missing")
		code = E04
		return
	}
	if strings.ToUpper(data.GetValue("result_code").(string)) != "SUCCESS" {
		errCode := data.GetValue("err_code")
		if errCode == SYSTEMERROR || errCode == BANKERROR || errCode == USERPAYING {
			err = errors.New(MESSAGE_PAYING)
			code = E03
			return
		} else {
			err = fmt.Errorf("%v:%v-%v", MESSAGE_WECHAT, errCode, data.GetValue("err_code_des"))
			code = E03
			return
		}
	}
	if !data.IsSet("sign") {
		err = errors.New("Signature verification failed")
		code = E04
		return
	}
	wxSign := data.DataAttr["sign"].(string)
	delete(data.DataAttr, "sign")
	if sign.CheckMd5Sign(base.JoinMapObject(data.DataAttr), key, wxSign) == false {
		err = errors.New("Signature verification failed")
		code = E04
		return
	}
	code = SUC
	result = data.DataAttr
	return

}

func GzipDecode(in []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(in))
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	return ioutil.ReadAll(reader)
}
