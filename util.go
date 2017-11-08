package wxpay

import (
	"errors"
	"fmt"
	"strings"

	"github.com/relax-space/go-kit/base"
	"github.com/relax-space/go-kit/data"
	"github.com/relax-space/go-kit/random"
	"github.com/relax-space/go-kit/sign"
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
	URLPAY     = "https://api.mch.weixin.qq.com/pay/micropay"
	URLREVERSE = "https://api.mch.weixin.qq.com/secapi/pay/reverse"
	URLREFUND  = "https://api.mch.weixin.qq.com/secapi/pay/refund"
	URLQUERY   = "https://api.mch.weixin.qq.com/pay/orderquery"
	URLPREPAY  = "https://api.mch.weixin.qq.com/pay/unifiedorder"
)

func BuildCommonparam(baseDto reqBaseDto) *data.Data {
	data := data.NewDataObj()
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

func RespParse(bodyByte []byte, key string) (result map[string]interface{}, err error) {
	data := data.NewDataObj()
	err = data.FromXml(string(bodyByte))
	if err != nil {
		return
	}

	if !data.IsSet("return_code") || strings.ToUpper(data.GetValue("return_code").(string)) != "SUCCESS" {
		err = fmt.Errorf("%v:%v", MESSAGE_WECHAT, data.GetValue("return_msg"))
		return
	}
	if !(data.IsSet("result_code")) || strings.ToUpper(data.GetValue("result_code").(string)) != "SUCCESS" {
		errCode := data.GetValue("err_code")
		if errCode == SYSTEMERROR || errCode == BANKERROR || errCode == USERPAYING {
			err = errors.New(MESSAGE_PAYING)
			return
		} else {
			err = fmt.Errorf("%v:%v-%v", MESSAGE_WECHAT, errCode, data.GetValue("err_code_des"))
			return
		}
	}
	if !data.IsSet("sign") {
		err = errors.New("Signature verification failed")
		return
	}
	wxSign := data.DataAttr["sign"].(string)
	delete(data.DataAttr, "sign")
	if sign.CheckSign(base.JoinMapObject(data.DataAttr), key, wxSign) == false {
		err = errors.New("Signature verification failed")
		return
	}
	result = data.DataAttr
	return

}
