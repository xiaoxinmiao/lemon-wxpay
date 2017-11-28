package wxpay

type ReqBaseDto struct {
	AppId    string `json:"appid,omitempty"`
	SubAppId string `json:"sub_appid,omitempty"`
	MchId    string `json:"mch_id,omitempty"`
	SubMchId string `json:"sub_mch_id,omitempty"`
	NonceStr string `json:"nonce_str,omitempty"`

	Sign string `json:"sign,omitempty"`
}

type ReqPayDto struct {
	*ReqBaseDto
	DeviceInfo string `json:"device_info,omitempty"`
	Body       string `json:"body,omitempty"`
	Detail     string `json:"detail,omitempty"`
	Attach     string `json:"attach,omitempty"`
	TotalFee   int64  `json:"total_fee,omitempty"` //int64

	FeeType        string `json:"fee_type,omitempty"`
	SpbillCreateIp string `json:"spbill_create_ip,omitempty"`
	GoodsTag       string `json:"goods_tag,omitempty"`
	LimitPay       string `json:"limit_pay,omitempty"`
	AuthCode       string `json:"auth_code,omitempty"` //10,11,12,13,14,15

	OutTradeNo string       `json:"out_trade_no,omitempty"`
	SceneInfo  SceneInfoDto `json:"scene_info,omitempty"`
}

type ReqPrepayDto struct {
	*ReqBaseDto
	SignType   string `json:"sign_type,omitempty" query:"sign_type"`
	Body       string `json:"body,omitempty" query:"body"`
	Detail     string `json:"detail,omitempty" query:"detail"`
	Attach     string `json:"attach,omitempty" query:"attach"`
	DeviceInfo string `json:"device_info,omitempty" query:"device_info"`

	FeeType        string `json:"fee_type,omitempty" query:"fee_type"`
	TotalFee       int64  `json:"total_fee,omitempty" query:"total_fee"`
	SpbillCreateIp string `json:"spbill_create_ip,omitempty" query:"spbill_create_ip"`
	TimeStart      string `json:"time_start,omitempty" query:"time_start"`
	TimeExpire     string `json:"time_expire,omitempty" query:"time_expire"`

	GoodsTag  string `json:"goods_tag,omitempty" query:"goods_tag"`
	NotifyUrl string `json:"notify_url,omitempty" query:"notify_url"`
	TradeType string `json:"trade_type,omitempty" query:"trade_type"`
	ProductId string `json:"product_id,omitempty" query:"product_id"`
	LimitPay  string `json:"limit_pay,omitempty" query:"limit_pay"`

	OpenId     string       `json:"openid,omitempty" query:"openid"`
	SubOpenId  string       `json:"sub_openid,omitempty" query:"sub_openid"`
	OutTradeNo string       `json:"out_trade_no,omitempty" query:"out_trade_no"`
	SceneInfo  SceneInfoDto `json:"scene_info,omitempty" query:"scene_info"`
}

type SceneInfoDto struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	AreaCode string `json:"area_code,omitempty"`
	Address  string `json:"address,omitempty"`
}

type ReqQueryDto struct {
	*ReqBaseDto
	OutTradeNo    string `json:"out_trade_no,omitempty"`
	TransactionId string `json:"transaction_id,omitempty"`
}

type ReqRefundDto struct {
	*ReqBaseDto
	OutRefundNo   string `json:"out_refund_no,omitempty"`
	DeviceInfo    string `json:"device_info,omitempty"`
	TransactionId string `json:"transaction_id,omitempty"`
	TotalFee      int64  `json:"total_fee,omitempty"`  //int64
	RefundFee     int64  `json:"refund_fee,omitempty"` //int64

	RefundFeeType string `json:"refund_fee_type,omitempty"`
	RefundAccount string `json:"refund_account,omitempty"`
	OutTradeNo    string `json:"out_trade_no,omitempty"`
}
type ReqReverseDto struct {
	*ReqBaseDto
	TransactionId string `json:"transaction_id,omitempty"`
	OutTradeNo    string `json:"out_trade_no,omitempty"`
}
type ReqRefundQueryDto struct {
	*ReqBaseDto
	TransactionId string `json:"transaction_id,omitempty"`
	OutTradeNo    string `json:"out_trade_no,omitempty"`
	OutRefundNo   string `json:"out_refund_no,omitempty"`
	RefundId      string `json:"refund_id,omitempty"`
	Offset        int    `json:"offset,omitempty"`
}

type ReqCustomerDto struct {
	Key              string `json:"key,omitempty"`
	CertPathName     string `json:"cert_path_name,omitempty"`
	CertPathKey      string `json:"cert_path_key,omitempty"`
	RootCa           string `json:"root_ca,omitempty"`
	UnifiedNotifyUrl string `json:"unified_notify_url,omitempty"`
}

//response

type RespDto struct {
	ReturnCode string `json:"return_code,omitempty"`
	ReturnMsg  string `json:"return_msg,omitempty"`
	AppId      string `json:"appid,omitempty"`
	SubAppId   string `json:"sub_appid,omitempty"`
	MchId      string `json:"mch_id,omitempty"`

	SubMchId   string `json:"sub_mch_id,omitempty"`
	NonceStr   string `json:"nonce_str,omitempty"`
	Sign       string `json:"sign,omitempty"`
	ResultCode string `json:"result_code,omitempty"`
	ErrCode    string `json:"err_code,omitempty"`

	ErrCodeDes string `json:"err_code_des,omitempty"`
}

type RespPayDto struct {
	DeviceInfo  string `json:"device_info,omitempty"`
	OpenId      string `json:"openid,omitempty"`
	IsSubscribe string `json:"is_subscribe,omitempty"`
	TradeType   string `json:"trade_type,omitempty"`
	BankType    string `json:"bank_type,omitempty"`

	FeeType             string `json:"fee_type,omitempty"`
	TotalFee            int64  `json:"total_fee,omitempty"`            //int64
	SettlementTotal_Fee int64  `json:"settlement_total_fee,omitempty"` //int64
	CouponFee           int64  `json:"coupon_fee,omitempty"`           //int64
	CashFeeType         string `json:"cash_fee_type,omitempty"`

	CashFee         int64  `json:"cash_fee,omitempty"` //int64
	TransactionId   string `json:"transaction_id,omitempty"`
	Attach          string `json:"attach,omitempty"`
	TimeEnd         string `json:"time_end,omitempty"` //yyyyMMddHHmmss
	PromotionDetail string `json:"promotion_detail,omitempty"`

	OutTradeNo string `json:"out_trade_no,omitempty"`
}

type RespQueryDto struct {
	DeviceInfo     string `json:"device_info,omitempty"`
	OpenId         string `json:"openid,omitempty"`
	IsSubscribe    string `json:"is_subscribe,omitempty"`
	SubOpenid      string `json:"sub_openid,omitempty"`
	SubIsSubscribe string `json:"sub_is_subscribe,omitempty"`

	TradeType  string `json:"trade_type,omitempty"`
	TradeState string `json:"trade_state,omitempty"`
	BankType   string `json:"bank_type,omitempty"`
	Detail     string `json:"detail,omitempty"`    //serviceProvidor
	TotalFee   int64  `json:"total_fee,omitempty"` //int64

	FeeType            string `json:"fee_type,omitempty"`
	SettlementTotalFee int64  `json:"settlement_total_fee,omitempty"` //int64
	CashFee            int64  `json:"cash_fee,omitempty"`             //int64
	CashFeeType        string `json:"cash_fee_type,omitempty"`
	CouponFee          int64  `json:"coupon_fee,omitempty"` //int64

	CouponCount int64 `json:"coupon_count,omitempty"` //int64

	CouponType0 string `json:"coupon_type_0,omitempty"`
	CouponId0   string `json:"coupon_id_0,omitempty"`
	CouponFee0  int64  `json:"coupon_fee_0,omitempty"` //int64

	CouponType1 string `json:"coupon_type_1,omitempty"`
	CouponId1   string `json:"coupon_id_1,omitempty"`
	CouponFee1  int64  `json:"coupon_fee_1,omitempty"` //int64

	CouponType2 string `json:"coupon_type_2,omitempty"`
	CouponId2   string `json:"coupon_id_2,omitempty"`
	CouponFee2  int64  `json:"coupon_fee_2,omitempty"` //int64

	CouponType3 string `json:"coupon_type_3,omitempty"`
	CouponId3   string `json:"coupon_id_3,omitempty"`
	CouponFee3  int64  `json:"coupon_fee_3,omitempty"` //int64

	CouponType4 string `json:"coupon_type_4,omitempty"`
	CouponId4   string `json:"coupon_id_4,omitempty"`
	CouponFee4  int64  `json:"coupon_fee_4,omitempty"` //int64

	TransactionId  string `json:"transaction_id,omitempty"`
	Attach         string `json:"attach,omitempty"`
	TimeEnd        string `json:"time_end,omitempty"`
	TradeStateDesc string `json:"trade_state_desc,omitempty"`
	OutTradeNo     string `json:"out_trade_no,omitempty"`
}

type RespRefundDto struct {
	DeviceInfo    string `json:"device_info,omitempty"`
	TransactionId string `json:"transaction_id,omitempty"`
	OutRefundNo   string `json:"out_refund_no,omitempty"`
	RefundId      string `json:"refund_id,omitempty"`

	RefundFee            int64  `json:"refund_fee,omitempty"`            //int64
	SettlementRefund_Fee int64  `json:"settlement_refund_fee,omitempty"` //int64
	TotalFee             int64  `json:"total_fee,omitempty"`             //int64
	SettlementTotalFee   int64  `json:"settlement_total_fee,omitempty"`  //int64
	FeeType              string `json:"fee_type,omitempty"`

	CashFee           int64  `json:"cash_fee,omitempty"`            //int64
	CashRefundFee     int64  `json:"cash_refund_fee,omitempty"`     //int64
	CouponRefundFee   int64  `json:"coupon_refund_fee,omitempty"`   //int64
	CouponRefundCount int64  `json:"coupon_refund_count,omitempty"` //int64
	CouponType0       string `json:"coupon_type_0,omitempty"`
	CouponRefundId0   string `json:"coupon_refund_id_0,omitempty"`
	CouponRefundFee0  int64  `json:"coupon_refund_fee_0,omitempty"` //int64

	CouponType1      string `json:"coupon_type_1,omitempty"`
	CouponRefundId1  string `json:"coupon_refund_id_1,omitempty"`
	CouponRefundFee1 int64  `json:"coupon_refund_fee_1,omitempty"` //int64

	CouponType2      string `json:"coupon_type_2,omitempty"`
	CouponRefundId2  string `json:"coupon_refund_id_2,omitempty"`
	CouponRefundFee2 int64  `json:"coupon_refund_fee_2,omitempty"` //int64

	CouponType3      string `json:"coupon_type_3,omitempty"`
	CouponRefundId3  string `json:"coupon_refund_id_3,omitempty"`
	CouponRefundFee3 int64  `json:"coupon_refund_fee_3,omitempty"` //int64

	CouponType4      string `json:"coupon_type_4,omitempty"`
	CouponRefundId4  string `json:"coupon_refund_id_4,omitempty"`
	CouponRefundFee4 int64  `json:"coupon_refund_fee_4,omitempty"` //int64

	OutTradeNo string `json:"out_trade_no,omitempty"`
}
type RespReverseDto struct {
	Recall string `json:"recall,omitempty"`
}

type RespPrepayDto struct {
	TradeType string `json:"trade_type,omitempty"`
	PrepayId  string `json:"prepay_id,omitempty"`
	CodeUrl   string `json:"code_url,omitempty"`
}

type NotifyDto struct {
	ReturnCode string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg  string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	AppId      string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchId      string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	SubAppId   string `xml:"sub_appid,omitempty" json:"sub_appid,omitempty"`

	SubMchId   string `xml:"sub_mch_id,omitempty" json:"sub_mch_id,omitempty"`
	DeviceInfo string `xml:"device_info,omitempty" json:"device_info,omitempty"`
	NonceStr   string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign       string `xml:"sign,omitempty" json:"sign,omitempty"`
	ResultCode string `xml:"result_code,omitempty" json:"result_code,omitempty"`

	ErrCode     string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes  string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	OpenId      string `xml:"openid,omitempty" json:"openid,omitempty"`
	IsSubscribe string `xml:"is_subscribe,omitempty" json:"is_subscribe,omitempty"`
	SubOpenId   string `xml:"sub_openid,omitempty" json:"sub_openid,omitempty"`

	SubIsSubscribe string `xml:"sub_is_subscribe,omitempty" json:"sub_is_subscribe,omitempty"`
	TradeType      string `xml:"trade_type,omitempty" json:"trade_type,omitempty"`
	BankType       string `xml:"bank_type,omitempty" json:"bank_type,omitempty"`
	TotalFee       int64  `xml:"total_fee,omitempty" json:"total_fee,omitempty"` //int64
	FeeType        string `xml:"fee_type,omitempty" json:"fee_type,omitempty"`

	CashFee            int64  `xml:"cash_fee,omitempty" json:"cash_fee,omitempty"` //int64
	CashFeeType        string `xml:"cash_fee_type,omitempty" json:"cash_fee_type,omitempty"`
	SettlementTotalFee int64  `xml:"settlement_total_fee,omitempty" json:"settlement_total_fee,omitempty"` //int64
	CouponFee          int64  `xml:"coupon_fee,omitempty" json:"coupon_fee,omitempty"`                     //int64
	CouponCount        int64  `xml:"coupon_count,omitempty" json:"coupon_count,omitempty"`                 //int64

	TransactionId string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`
	OutTradeNo    string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
	Attach        string `xml:"attach,omitempty" json:"attach,omitempty"`
	TimeEnd       string `xml:"time_end,omitempty" json:"time_end,omitempty"`
}

type RespRefundQueryDto struct {
	TransactionId      string `json:"transaction_id;omitempty"`
	OutTradeNo         string `json:"out_trade_no;omitempty"`
	TotalFee           int64  `json:"total_fee;omitempty"`
	SettlementTotalFee int64  `json:"settlement_total_fee;omitempty"`
	FeeType            string `json:"fee_type;omitempty"`

	CashFee          string `json:"cash_fee;omitempty"`
	RefundCount      string `json:"refund_count;omitempty"`
	TotalRefundCount string `json:"total_refund_count;omitempty"`
}
