package wxpay

import (
	"kit/util"
	"strings"

	"github.com/relax-space/go-kit/base"
	"github.com/relax-space/go-kit/data"
	"github.com/relax-space/go-kit/sign"
)

const (
	SYSTEMERROR   = "SYSTEMERROR"
	ORDERNOTEXIST = "ORDERNOTEXIST"
	BANKERROR     = "BANKERROR"
	USERPAYING    = "USERPAYING"
)

const (
	OUTTRADENO    = "14"
	OUTREFUNDNO   = "15"
	PREOUTTRADENO = "16"
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
	SetValue(data, "nonce_str", util.Uuid(""))
	return data
}

func SetValue(data *data.Data, key string, value interface{}) {
	if len(base.ToString(value)) != 0 {
		data.SetValue(key, value)
	}
}

func RespParse(bodyByte []byte, key string) (result map[string]interface{}, err error) {
	//wxResponse := new(WxPayDataResp)
	data := data.NewDataObj()
	err = data.FromXml(string(bodyByte))
	if err != nil {
		//to be
		//statusMessage = util.NewApiStatusMessage(http.StatusInternalServerError, 20004, "")
		return
	}
	//util.Info(wxResponse)

	if strings.ToUpper(data.GetValue("return_code").(string)) != "SUCCESS" {
		//statusMessage = util.NewApiStatusMessage(http.StatusInternalServerError, 20004, wxResponse.GetValue(respKeys.Base.ReturnMsg).(string))
		return
	}
	if !(data.IsSet("result_code")) || strings.ToUpper(data.GetValue("result_code").(string)) != "SUCCESS" {
		errCode := data.GetValue("err_code")
		if errCode == SYSTEMERROR || errCode == BANKERROR || errCode == USERPAYING {
			//statusMessage = util.NewApiStatusMessage(http.StatusOK, 20006, "")
			return
		} else {
			//errMsg := data.GetValue("err_code_des")
			//statusMessage = util.NewApiStatusMessage(http.StatusOK, 20005, "", fmt.Sprintf("errCode:%v,errMsg:%v", errCode, errMsg))
			return
		}
	}

	if !data.IsSet("sign") {
		//to be
		return
	}
	if sign.CheckSign(base.JoinMapObject(data.DataAttr), key, data.DataAttr["sign"].(string)) == false {
		//statusMessage = util.NewApiStatusMessage(http.StatusOK, 20002, "")
		return
	}
	result = data.DataAttr
	return

}
