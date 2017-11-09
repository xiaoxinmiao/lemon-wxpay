# lemon-wxpay
wechat pay go-sdk

## Installation
```
go get github.com/relax-space/lemon-wxpay
```

## Usage
> pay
> query
> refund
> reverse

### pay
```go
	reqDto := wxpay.ReqPayDto{
		ReqBaseDto: wxpay.ReqBaseDto{
			AppId: "****",
			MchId: "****",
		},
		AuthCode: "134511727606227397",
		Body:     "xiaoxinmiao test",
		TotalFee: 1,
	}
	customDto := wxpay.ReqCustomerDto{
		Key: "****",
	}
	result, err := wxpay.Pay(reqDto, customDto)
    fmt.Printf("%+v,%v", result,err)
```

### query
```go
	reqDto := wxpay.ReqQueryDto{
		ReqBaseDto: wxpay.ReqBaseDto{
			AppId: "****",
			MchId: "****",
		},
		OutTradeNo: "14201711085205823413229775520",
	}
	customDto := wxpay.ReqCustomerDto{
		Key: "****",
	}
	result, err := wxpay.Query(reqDto, customDto)
    fmt.Printf("%+v,%v", result,err)
```

### refund
```go
	reqDto := wxpay.ReqRefundDto{
		ReqBaseDto: wxpay.ReqBaseDto{
			AppId: "****",
			MchId: "****",
		},
		OutTradeNo: "14201711085205823413229775520",
        RefundFee:  1,
	}
	customDto := wxpay.ReqCustomerDto{
		Key: "****",
        CertPathName: "c:/cert/apiclient_cert.pem",
		CertPathKey:  "c:/cert/apiclient_key.pem",
		RootCa:       "c:/cert/rootca.pem"),
	}
	result, err := wxpay.Refund(reqDto, customDto)
    fmt.Printf("%+v,%v", result,err)
```

### reverse
```go
	reqDto := wxpay.ReqReverseDto{
		ReqBaseDto: wxpay.ReqBaseDto{
			AppId: "****",
			MchId: "****",
		},
		OutTradeNo: "14201711085205823413229775520",
	}
	customDto := wxpay.ReqCustomerDto{
		Key: "****",
        CertPathName: "c:/cert/apiclient_cert.pem",
		CertPathKey:  "c:/cert/apiclient_key.pem",
		RootCa:       "c:/cert/rootca.pem"),
	}
	result, err := wxpay.Reverse(reqDto, customDto)
    fmt.Printf("%+v,%v", result,err)
```

