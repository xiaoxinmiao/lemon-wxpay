package wxpay

import (
	"flag"
	"fmt"
	"kit/test"
	"os"
	"testing"
)

var (
	appId = flag.String("WXPAY_APPID", os.Getenv("WXPAY_APPID"), "WXPAY_APPID")
	key   = flag.String("WXPAY_KEY", os.Getenv("WXPAY_KEY"), "WXPAY_KEY")
	mchid = flag.String("WXPAY_MCHID", os.Getenv("WXPAY_MCHID"), "WXPAY_MCHID")
)

func Test_Pay(t *testing.T) {
	reqDto := reqPayDto{
		reqBaseDto: reqBaseDto{
			AppId: *appId,
			MchId: *mchid,
		},
		AuthCode: "134511727606227397",
		Body:     "xiaoxinmiao test",
		TotalFee: 1,
	}
	customDto := reqCustomerDto{
		Key: *key,
	}
	result, err := Pay(reqDto, customDto)
	if err != nil {
		if err.Error() == "MESSAGE_PAYING" {
			dto := reqQueryDto{
				reqBaseDto: reqDto.reqBaseDto,
				OutTradeNo: result["out_trade_no"].(string),
			}
			queryResult, err := LoopQuery(dto, customDto, 40, 2)
			fmt.Printf("%+v", queryResult)
			test.Ok(t, err)
			return
		}
		test.Ok(t, err)
	}
	fmt.Printf("%+v", result)
	test.Ok(t, err)

}

func Test_Query(t *testing.T) {
	reqDto := reqQueryDto{
		reqBaseDto: reqBaseDto{
			AppId: *appId,
			MchId: *mchid,
		},
		OutTradeNo: "14201711085205823413229775520",
	}
	custDto := reqCustomerDto{
		Key: *key,
	}
	result, err := Query(reqDto, custDto)
	fmt.Printf("%+v", result)
	test.Ok(t, err)
}
