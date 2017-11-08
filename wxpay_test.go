package wxpay

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/relax-space/go-kit/test"
)

var (
	appId = flag.String("WXPAY_APPID", os.Getenv("WXPAY_APPID"), "WXPAY_APPID")
	key   = flag.String("WXPAY_KEY", os.Getenv("WXPAY_KEY"), "WXPAY_KEY")
	mchId = flag.String("WXPAY_MCHID", os.Getenv("WXPAY_MCHID"), "WXPAY_MCHID")
)

func Test_Pay(t *testing.T) {
	reqDto := reqPayDto{
		reqBaseDto: reqBaseDto{
			AppId: *appId,
			MchId: *mchId,
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
			MchId: *mchId,
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

func Test_Refund(t *testing.T) {
	reqDto := reqRefundDto{
		reqBaseDto: reqBaseDto{
			AppId: *appId,
			MchId: *mchId,
		},
		OutTradeNo: "14201711085205823413229775520",
		RefundFee:  1,
	}
	custDto := reqCustomerDto{
		Key:          *key,
		CertPathName: fmt.Sprintf("c:/cert/%v/apiclient_cert.pem", *mchId),
		CertPathKey:  fmt.Sprintf("c:/cert/%v/apiclient_key.pem", *mchId),
		RootCa:       fmt.Sprintf("c:/cert/%v/rootca.pem", *mchId),
	}
	result, err := Refund(reqDto, custDto)
	fmt.Printf("%+v", result)
	test.Ok(t, err)
}

func Test_Reverse(t *testing.T) {
	reqDto := reqReverseDto{
		reqBaseDto: reqBaseDto{
			AppId: *appId,
			MchId: *mchId,
		},
		OutTradeNo: "1417084862106336446985",
	}
	custDto := reqCustomerDto{
		Key:          *key,
		CertPathName: fmt.Sprintf("c:/cert/%v/apiclient_cert.pem", *mchId),
		CertPathKey:  fmt.Sprintf("c:/cert/%v/apiclient_key.pem", *mchId),
		RootCa:       fmt.Sprintf("c:/cert/%v/rootca.pem", *mchId),
	}
	result, err := Reverse(reqDto, custDto, 10, 10)
	fmt.Printf("%+v", result)
	test.Ok(t, err)
}
