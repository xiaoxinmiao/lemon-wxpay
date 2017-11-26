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
	reqDto := ReqPayDto{
		ReqBaseDto: &ReqBaseDto{
			AppId: *appId,
			MchId: *mchId,
		},
		AuthCode: "134511727606227397",
		Body:     "xiaoxinmiao test",
		TotalFee: 1,
	}
	customDto := ReqCustomerDto{
		Key: *key,
	}
	result, err := Pay(&reqDto, &customDto)
	if err != nil {
		if err.Error() == "MESSAGE_PAYING" {
			dto := ReqQueryDto{
				ReqBaseDto: reqDto.ReqBaseDto,
				OutTradeNo: result["out_trade_no"].(string),
			}
			queryResult, err := LoopQuery(&dto, &customDto, 40, 2)
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
	reqDto := ReqQueryDto{
		ReqBaseDto: &ReqBaseDto{
			AppId: *appId,
			MchId: *mchId,
		},
		OutTradeNo: "14201711085205823413229775520",
	}
	custDto := ReqCustomerDto{
		Key: *key,
	}
	result, err := Query(&reqDto, &custDto)
	fmt.Printf("%+v", result)
	test.Ok(t, err)
}

func Test_Refund(t *testing.T) {
	reqDto := ReqRefundDto{
		ReqBaseDto: &ReqBaseDto{
			AppId: *appId,
			MchId: *mchId,
		},
		OutTradeNo: "14201711085205823413229775520",
		RefundFee:  1,
	}
	custDto := ReqCustomerDto{
		Key:          *key,
		CertPathName: fmt.Sprintf("c:/cert/%v/apiclient_cert.pem", *mchId),
		CertPathKey:  fmt.Sprintf("c:/cert/%v/apiclient_key.pem", *mchId),
		RootCa:       fmt.Sprintf("c:/cert/%v/rootca.pem", *mchId),
	}
	result, err := Refund(&reqDto, &custDto)
	fmt.Printf("%+v", result)
	test.Ok(t, err)
}

func Test_Reverse(t *testing.T) {
	reqDto := ReqReverseDto{
		ReqBaseDto: &ReqBaseDto{
			AppId: *appId,
			MchId: *mchId,
		},
		OutTradeNo: "1417084862106336446985",
	}
	custDto := ReqCustomerDto{
		Key:          *key,
		CertPathName: fmt.Sprintf("c:/cert/%v/apiclient_cert.pem", *mchId),
		CertPathKey:  fmt.Sprintf("c:/cert/%v/apiclient_key.pem", *mchId),
		RootCa:       fmt.Sprintf("c:/cert/%v/rootca.pem", *mchId),
	}
	result, err := Reverse(&reqDto, &custDto, 10, 10)
	fmt.Printf("%+v", result)
	test.Ok(t, err)
}

func Test_RefundQuery(t *testing.T) {
	reqDto := ReqRefundQueryDto{
		ReqBaseDto: &ReqBaseDto{
			AppId: *appId,
			MchId: *mchId,
		},
		OutTradeNo: "144650782494807835413",
	}
	custDto := ReqCustomerDto{
		Key: *key,
	}
	result, err := RefundQuery(&reqDto, &custDto)
	fmt.Printf("%+v", result)
	test.Ok(t, err)
}

func Test_PrePay(t *testing.T) {
	reqDto := ReqPrePayDto{
		ReqBaseDto: &ReqBaseDto{
			AppId: *appId,
			MchId: *mchId,
		},
		Body:      "xiaomiao test",
		TotalFee:  1,
		TradeType: PREPAY_TYPE_JSAPI,
		NotifyUrl: "https://xiaomiao.com",
		OpenId:    "os2u9uPKLkCKL08FwCM6hQAQ_LtI",
	}
	custDto := ReqCustomerDto{
		Key: *key,
	}
	result, err := PrePay(&reqDto, &custDto)
	fmt.Printf("%+v", result)
	test.Ok(t, err)
}

func Test_Notify(t *testing.T) {
	xmlBody := `<xml>
	<appid><![CDATA[wx2421b1c4370ec43b]]></appid>
	<attach><![CDATA[{"sub_notify_url":"https://baidu.com"}]]></attach>
	<bank_type><![CDATA[CFT]]></bank_type>
	<fee_type><![CDATA[CNY]]></fee_type>
	<is_subscribe><![CDATA[Y]]></is_subscribe>
	<mch_id><![CDATA[10000100]]></mch_id>
	<nonce_str><![CDATA[5d2b6c2a8db53831f7eda20af46e531c]]></nonce_str>
	<openid><![CDATA[oUpF8uMEb4qRXf22hE3X68TekukE]]></openid>
	<out_trade_no><![CDATA[1409811653]]></out_trade_no>
	<result_code><![CDATA[SUCCESS]]></result_code>
	<return_code><![CDATA[SUCCESS]]></return_code>
	<sign><![CDATA[B552ED6B279343CB493C5DD0D78AB241]]></sign>
	<sub_mch_id><![CDATA[10000100]]></sub_mch_id>
	<time_end><![CDATA[20140903131540]]></time_end>
	<total_fee>2</total_fee>
	<trade_type><![CDATA[JSAPI]]></trade_type>
	<transaction_id><![CDATA[1004400740201409030005092168]]></transaction_id>
 </xml>`
	result, err := Notify(xmlBody)
	fmt.Printf("%+v,%v", result, err)
	test.Ok(t, err)
}
