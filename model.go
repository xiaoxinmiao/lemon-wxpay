package wxpay

type ReqBaseDto struct {
	AppId    string `json:"appid,omitempty" query:"appid,omitempty"`
	SubAppId string `json:"sub_appid,omitempty" query:"sub_appid,omitempty"`
	MchId    string `json:"mch_id,omitempty" query:"mch_id,omitempty"`
	SubMchId string `json:"sub_mch_id,omitempty" query:"sub_mch_id,omitempty"`
	NonceStr string `json:"nonce_str,omitempty" query:"nonce_str,omitempty"`

	Sign string `json:"sign,omitempty" query:"sign,omitempty"`
}

type ReqPayDto struct {
	*ReqBaseDto
	DeviceInfo string `json:"device_info,omitempty" query:"device_info,omitempty"`
	Body       string `json:"body,omitempty" query:"body,omitempty"`
	Detail     string `json:"detail,omitempty" query:"detail,omitempty"`
	Attach     string `json:"attach,omitempty" query:"attach,omitempty"`
	OutTradeNo string `json:"out_trade_no,omitempty" query:"out_trade_no,omitempty"`

	TotalFee       int64  `json:"total_fee,omitempty" query:"total_fee,omitempty"` //int64
	FeeType        string `json:"fee_type,omitempty" query:"fee_type,omitempty"`
	SpbillCreateIp string `json:"spbill_create_ip,omitempty" query:"spbill_create_ip,omitempty"`
	GoodsTag       string `json:"goods_tag,omitempty" query:"goods_tag,omitempty"`
	LimitPay       string `json:"limit_pay,omitempty" query:"limit_pay,omitempty"`

	TimeStart  string       `json:"time_start,omitempty" query:"time_start,omitempty"`
	TimeExpire string       `json:"time_expire,omitempty" query:"time_expire,omitempty"`
	AuthCode   string       `json:"auth_code,omitempty" query:"auth_code,omitempty"` //10,11,12,13,14,15
	SceneInfo  SceneInfoDto `json:"scene_info,omitempty" query:"scene_info,omitempty"`
}

type ReqPrepayDto struct {
	*ReqBaseDto
	DeviceInfo string `json:"device_info,omitempty" query:"device_info"`
	SignType   string `json:"sign_type,omitempty" query:"sign_type"`
	Body       string `json:"body,omitempty" query:"body"`
	Detail     string `json:"detail,omitempty" query:"detail"`
	Attach     string `json:"attach,omitempty" query:"attach"`

	OutTradeNo     string `json:"out_trade_no,omitempty" query:"out_trade_no"`
	FeeType        string `json:"fee_type,omitempty" query:"fee_type"`
	TotalFee       int64  `json:"total_fee,omitempty" query:"total_fee"`
	SpbillCreateIp string `json:"spbill_create_ip,omitempty" query:"spbill_create_ip"`
	TimeStart      string `json:"time_start,omitempty" query:"time_start"`

	TimeExpire string `json:"time_expire,omitempty" query:"time_expire"`
	GoodsTag   string `json:"goods_tag,omitempty" query:"goods_tag"`
	NotifyUrl  string `json:"notify_url,omitempty" query:"notify_url"`
	TradeType  string `json:"trade_type,omitempty" query:"trade_type"`
	ProductId  string `json:"product_id,omitempty" query:"product_id"`

	LimitPay  string       `json:"limit_pay,omitempty" query:"limit_pay"`
	OpenId    string       `json:"openid,omitempty" query:"openid"`
	SubOpenId string       `json:"sub_openid,omitempty" query:"sub_openid"`
	SceneInfo SceneInfoDto `json:"scene_info,omitempty" query:"scene_info"`
}

type ReqQueryDto struct {
	*ReqBaseDto
	OutTradeNo    string `json:"out_trade_no,omitempty" query:"out_trade_no,omitempty"`
	TransactionId string `json:"transaction_id,omitempty" query:"transaction_id,omitempty"`
}

type ReqRefundDto struct {
	*ReqBaseDto
	TransactionId string `json:"transaction_id,omitempty" query:"transaction_id,omitempty"`
	OutTradeNo    string `json:"out_trade_no,omitempty" query:"out_trade_no,omitempty"`
	OutRefundNo   string `json:"out_refund_no,omitempty" query:"out_refund_no,omitempty"`
	TotalFee      int64  `json:"total_fee,omitempty" query:"total_fee,omitempty"`   //int64
	RefundFee     int64  `json:"refund_fee,omitempty" query:"refund_fee,omitempty"` //int64

	RefundFeeType string `json:"refund_fee_type,omitempty" query:"refund_fee_type,omitempty"`
	RefundDesc    string `json:"refund_desc,omitempty" query:"refund_desc,omitempty"`
	RefundAccount string `json:"refund_account,omitempty" query:"refund_account,omitempty"`
}
type ReqReverseDto struct {
	*ReqBaseDto
	TransactionId string `json:"transaction_id,omitempty" query:"transaction_id,omitempty"`
	OutTradeNo    string `json:"out_trade_no,omitempty" query:"out_trade_no,omitempty"`
}
type ReqRefundQueryDto struct {
	*ReqBaseDto
	TransactionId string `json:"transaction_id,omitempty" query:"transaction_id,omitempty"`
	OutTradeNo    string `json:"out_trade_no,omitempty" query:"out_trade_no,omitempty"`
	OutRefundNo   string `json:"out_refund_no,omitempty" query:"out_refund_no,omitempty"`
	RefundId      string `json:"refund_id,omitempty" query:"refund_id,omitempty"`
	Offset        int    `json:"offset,omitempty" query:"offset,omitempty"`
}

type ReqBillDto struct {
	*ReqBaseDto
	BillDate string `json:"bill_date,omitempty" query:"bill_date,omitempty"` //20181217
	BillType string `json:"bill_type,omitempty" query:"bill_type,omitempty"`
	TarType  string `json:"tar_type,omitempty" query:"tar_type,omitempty"`
}

type ReqCustomerDto struct {
	Key              string `json:"key,omitempty" query:"key,omitempty"`
	CertPathName     string `json:"cert_path_name,omitempty" query:"cert_path_name,omitempty"`
	CertPathKey      string `json:"cert_path_key,omitempty" query:"cert_path_key,omitempty"`
	RootCa           string `json:"root_ca,omitempty" query:"root_ca,omitempty"`
	UnifiedNotifyUrl string `json:"unified_notify_url,omitempty" query:"unified_notify_url,omitempty"`
}

//response

type RespDto struct {
	ReturnCode string `json:"return_code,omitempty" query:"return_code,omitempty"`
	ReturnMsg  string `json:"return_msg,omitempty" query:"return_msg,omitempty"`
	AppId      string `json:"appid,omitempty" query:"appid,omitempty"`
	SubAppId   string `json:"sub_appid,omitempty" query:"sub_appid,omitempty"`
	MchId      string `json:"mch_id,omitempty" query:"mch_id,omitempty"`

	SubMchId   string `json:"sub_mch_id,omitempty" query:"sub_mch_id,omitempty"`
	NonceStr   string `json:"nonce_str,omitempty" query:"nonce_str,omitempty"`
	Sign       string `json:"sign,omitempty" query:"sign,omitempty"`
	ResultCode string `json:"result_code,omitempty" query:"result_code,omitempty"`
	ErrCode    string `json:"err_code,omitempty" query:"err_code,omitempty"`

	ErrCodeDes string `json:"err_code_des,omitempty" query:"err_code_des,omitempty"`
}

type RespPayDto struct {
	DeviceInfo     string `json:"device_info,omitempty" query:"device_info,omitempty"`
	OpenId         string `json:"openid,omitempty" query:"openid,omitempty"`
	IsSubscribe    string `json:"is_subscribe,omitempty" query:"is_subscribe,omitempty"`
	SubOpenid      string `json:"sub_openid,omitempty" query:"sub_openid,omitempty"`
	SubIsSubscribe string `json:"sub_is_subscribe,omitempty" query:"sub_is_subscribe,omitempty"`

	TradeType   string `json:"trade_type,omitempty" query:"trade_type,omitempty"`
	BankType    string `json:"bank_type,omitempty" query:"bank_type,omitempty"`
	FeeType     string `json:"fee_type,omitempty" query:"fee_type,omitempty"`
	TotalFee    int64  `json:"total_fee,omitempty" query:"total_fee,omitempty"` //int64
	CashFeeType string `json:"cash_fee_type,omitempty" query:"cash_fee_type,omitempty"`

	CashFee            int64  `json:"cash_fee,omitempty" query:"cash_fee,omitempty"`                         //int64
	SettlementTotalFee int64  `json:"settlement_total_fee,omitempty" query:"settlement_total_fee,omitempty"` //int64
	CouponFee          int64  `json:"coupon_fee,omitempty" query:"coupon_fee,omitempty"`                     //int64
	TransactionId      string `json:"transaction_id,omitempty" query:"transaction_id,omitempty"`
	OutTradeNo         string `json:"out_trade_no,omitempty" query:"out_trade_no,omitempty"`

	Attach  string `json:"attach,omitempty" query:"attach,omitempty"`
	TimeEnd string `json:"time_end,omitempty" query:"time_end,omitempty"` //yyyyMMddHHmmss
}

type RespQueryDto struct {
	DeviceInfo     string `json:"device_info,omitempty" query:"device_info,omitempty"`
	OpenId         string `json:"openid,omitempty" query:"openid,omitempty"`
	IsSubscribe    string `json:"is_subscribe,omitempty" query:"is_subscribe,omitempty"`
	SubOpenid      string `json:"sub_openid,omitempty" query:"sub_openid,omitempty"`
	SubIsSubscribe string `json:"sub_is_subscribe,omitempty" query:"sub_is_subscribe,omitempty"`

	TradeType  string `json:"trade_type,omitempty" query:"trade_type,omitempty"`
	TradeState string `json:"trade_state,omitempty" query:"trade_state,omitempty"`
	BankType   string `json:"bank_type,omitempty" query:"bank_type,omitempty"`
	Detail     string `json:"detail,omitempty" query:"detail,omitempty"`       //serviceProvidor
	TotalFee   int64  `json:"total_fee,omitempty" query:"total_fee,omitempty"` //int64

	FeeType            string `json:"fee_type,omitempty" query:"fee_type,omitempty"`
	SettlementTotalFee int64  `json:"settlement_total_fee,omitempty" query:"settlement_total_fee,omitempty"` //int64
	CashFee            int64  `json:"cash_fee,omitempty" query:"cash_fee,omitempty"`                         //int64
	CashFeeType        string `json:"cash_fee_type,omitempty" query:"cash_fee_type,omitempty"`
	CouponFee          int64  `json:"coupon_fee,omitempty" query:"coupon_fee,omitempty"` //int64

	CouponCount int64 `json:"coupon_count,omitempty" query:"coupon_count,omitempty"` //int64

	CouponType0 string `json:"coupon_type_0,omitempty" query:"coupon_type_0,omitempty"`
	CouponId0   string `json:"coupon_id_0,omitempty" query:"coupon_id_0,omitempty"`
	CouponFee0  int64  `json:"coupon_fee_0,omitempty" query:"coupon_fee_0,omitempty"` //int64

	CouponType1 string `json:"coupon_type_1,omitempty" query:"coupon_type_1,omitempty"`
	CouponId1   string `json:"coupon_id_1,omitempty" query:"coupon_id_1,omitempty"`
	CouponFee1  int64  `json:"coupon_fee_1,omitempty" query:"coupon_fee_1,omitempty"` //int64

	CouponType2 string `json:"coupon_type_2,omitempty" query:"coupon_type_2,omitempty"`
	CouponId2   string `json:"coupon_id_2,omitempty" query:"coupon_id_2,omitempty"`
	CouponFee2  int64  `json:"coupon_fee_2,omitempty" query:"coupon_fee_2,omitempty"` //int64

	CouponType3 string `json:"coupon_type_3,omitempty" query:"coupon_type_3,omitempty"`
	CouponId3   string `json:"coupon_id_3,omitempty" query:"coupon_id_3,omitempty"`
	CouponFee3  int64  `json:"coupon_fee_3,omitempty" query:"coupon_fee_3,omitempty"` //int64

	CouponType4 string `json:"coupon_type_4,omitempty" query:"coupon_type_4,omitempty"`
	CouponId4   string `json:"coupon_id_4,omitempty" query:"coupon_id_4,omitempty"`
	CouponFee4  int64  `json:"coupon_fee_4,omitempty" query:"coupon_fee_4,omitempty"` //int64

	TransactionId  string `json:"transaction_id,omitempty" query:"transaction_id,omitempty"`
	OutTradeNo     string `json:"out_trade_no,omitempty" query:"out_trade_no,omitempty"`
	Attach         string `json:"attach,omitempty" query:"attach,omitempty"`
	TimeEnd        string `json:"time_end,omitempty" query:"time_end,omitempty"`
	TradeStateDesc string `json:"trade_state_desc,omitempty" query:"trade_state_desc,omitempty"`
}

type RespRefundDto struct {
	DeviceInfo    string `json:"device_info,omitempty" query:"device_info,omitempty"`
	TransactionId string `json:"transaction_id,omitempty" query:"transaction_id,omitempty"`
	OutRefundNo   string `json:"out_refund_no,omitempty" jsoqueryn:"out_refund_no,omitempty"`
	RefundId      string `json:"refund_id,omitempty" query:"refund_id,omitempty"`

	RefundFee            int64  `json:"refund_fee,omitempty" query:"refund_fee,omitempty"`                       //int64
	SettlementRefund_Fee int64  `json:"settlement_refund_fee,omitempty" query:"settlement_refund_fee,omitempty"` //int64
	TotalFee             int64  `json:"total_fee,omitempty" query:"total_fee,omitempty"`                         //int64
	SettlementTotalFee   int64  `json:"settlement_total_fee,omitempty" query:"settlement_total_fee,omitempty"`   //int64
	FeeType              string `json:"fee_type,omitempty" query:"fee_type,omitempty"`

	CashFee           int64  `json:"cash_fee,omitempty" query:"cash_fee,omitempty"`                       //int64
	CashRefundFee     int64  `json:"cash_refund_fee,omitempty" query:"cash_refund_fee,omitempty"`         //int64
	CouponRefundFee   int64  `json:"coupon_refund_fee,omitempty" query:"coupon_refund_fee,omitempty"`     //int64
	CouponRefundCount int64  `json:"coupon_refund_count,omitempty" query:"coupon_refund_count,omitempty"` //int64
	CouponType0       string `json:"coupon_type_0,omitempty" query:"coupon_type_0,omitempty"`
	CouponRefundId0   string `json:"coupon_refund_id_0,omitempty" query:"coupon_refund_id_0,omitempty"`
	CouponRefundFee0  int64  `json:"coupon_refund_fee_0,omitempty" query:"coupon_refund_fee_0,omitempty"` //int64

	CouponType1      string `json:"coupon_type_1,omitempty" query:"coupon_type_1,omitempty"`
	CouponRefundId1  string `json:"coupon_refund_id_1,omitempty" query:"coupon_refund_id_1,omitempty"`
	CouponRefundFee1 int64  `json:"coupon_refund_fee_1,omitempty" query:"coupon_refund_fee_1,omitempty"` //int64

	CouponType2      string `json:"coupon_type_2,omitempty" query:"coupon_type_2,omitempty"`
	CouponRefundId2  string `json:"coupon_refund_id_2,omitempty" query:"coupon_refund_id_2,omitempty"`
	CouponRefundFee2 int64  `json:"coupon_refund_fee_2,omitempty" query:"coupon_refund_fee_2,omitempty"` //int64

	CouponType3      string `json:"coupon_type_3,omitempty" query:"coupon_type_3,omitempty"`
	CouponRefundId3  string `json:"coupon_refund_id_3,omitempty" query:"coupon_refund_id_3,omitempty"`
	CouponRefundFee3 int64  `json:"coupon_refund_fee_3,omitempty" query:"coupon_refund_fee_3,omitempty"` //int64

	CouponType4      string `json:"coupon_type_4,omitempty" query:"coupon_type_4,omitempty"`
	CouponRefundId4  string `json:"coupon_refund_id_4,omitempty" query:"coupon_refund_id_4,omitempty"`
	CouponRefundFee4 int64  `json:"coupon_refund_fee_4,omitempty" query:"coupon_refund_fee_4,omitempty"` //int64

	OutTradeNo string `json:"out_trade_no,omitempty" query:"out_trade_no,omitempty"`
}
type RespReverseDto struct {
	Recall string `json:"recall,omitempty" query:"recall,omitempty"`
}

type RespPrepayDto struct {
	TradeType string `json:"trade_type,omitempty" query:"trade_type,omitempty"`
	PrepayId  string `json:"prepay_id,omitempty" query:"prepay_id,omitempty"`
	CodeUrl   string `json:"code_url,omitempty" query:"code_url,omitempty"`
}

type RespRefundQueryDto struct {
	TransactionId      string `json:"transaction_id;omitempty" query:"transaction_id;omitempty"`
	OutTradeNo         string `json:"out_trade_no;omitempty" query:"out_trade_no;omitempty"`
	TotalFee           int64  `json:"total_fee;omitempty" query:"total_fee;omitempty"`
	SettlementTotalFee int64  `json:"settlement_total_fee;omitempty" query:"settlement_total_fee;omitempty"`
	FeeType            string `json:"fee_type;omitempty" query:"fee_type;omitempty"`

	CashFee          int64 `json:"cash_fee;omitempty" query:"cash_fee;omitempty"`
	RefundCount      int64 `json:"refund_count;omitempty" query:"refund_count;omitempty"`
	TotalRefundCount int64 `json:"total_refund_count;omitempty" query:"total_refund_count;omitempty"`

	OutRefundNo1 string `json:"out_refund_no_1;omitempty" query:"out_refund_no_1;omitempty"`
	OutRefundNo2 string `json:"out_refund_no_2;omitempty" query:"out_refund_no_2;omitempty"`
	OutRefundNo3 string `json:"out_refund_no_3;omitempty" query:"out_refund_no_3;omitempty"`
	OutRefundNo4 string `json:"out_refund_no_4;omitempty" query:"out_refund_no_4;omitempty"`

	RefundId1 string `json:"refund_id_1;omitempty" query:"refund_id_1;omitempty"`
	RefundId2 string `json:"refund_id_2;omitempty" query:"refund_id_2;omitempty"`
	RefundId3 string `json:"refund_id_3;omitempty" query:"refund_id_3;omitempty"`
	RefundId4 string `json:"refund_id_4;omitempty" query:"refund_id_4;omitempty"`

	RefundChannel1 string `json:"refund_channel_1;omitempty" query:"refund_channel_1;omitempty"`
	RefundChannel2 string `json:"refund_channel_2;omitempty" query:"refund_channel_2;omitempty"`
	RefundChannel3 string `json:"refund_channel_3;omitempty" query:"refund_channel_3;omitempty"`
	RefundChannel4 string `json:"refund_channel_4;omitempty" query:"refund_channel_4;omitempty"`

	RefundFee1 int64 `json:"refund_fee_1;omitempty" query:"refund_fee_1;omitempty"`
	RefundFee2 int64 `json:"refund_fee_2;omitempty" query:"refund_fee_2;omitempty"`
	RefundFee3 int64 `json:"refund_fee_3;omitempty" query:"refund_fee_3;omitempty"`
	RefundFee4 int64 `json:"refund_fee_4;omitempty" query:"refund_fee_4;omitempty"`

	SettlementRefundFee1 int64 `json:"settlement_refund_fee_1;omitempty" query:"settlement_refund_fee_1;omitempty"`
	SettlementRefundFee2 int64 `json:"settlement_refund_fee_2;omitempty" query:"settlement_refund_fee_2;omitempty"`
	SettlementRefundFee3 int64 `json:"settlement_refund_fee_3;omitempty" query:"settlement_refund_fee_3;omitempty"`
	SettlementRefundFee4 int64 `json:"settlement_refund_fee_4;omitempty" query:"settlement_refund_fee_4;omitempty"`

	CouponType01 string `json:"coupon_type_0_1;omitempty" query:"coupon_type_0_1;omitempty"`
	CouponType02 string `json:"coupon_type_0_2;omitempty" query:"coupon_type_0_2;omitempty"`
	CouponType03 string `json:"coupon_type_0_3;omitempty" query:"coupon_type_0_3;omitempty"`
	CouponType04 string `json:"coupon_type_0_4;omitempty" query:"coupon_type_0_4;omitempty"`

	CouponRefundFee1 int64 `json:"coupon_refund_fee_1;omitempty" query:"coupon_refund_fee_1;omitempty"`
	CouponRefundFee2 int64 `json:"coupon_refund_fee_2;omitempty" query:"coupon_refund_fee_2;omitempty"`
	CouponRefundFee3 int64 `json:"coupon_refund_fee_3;omitempty" query:"coupon_refund_fee_3;omitempty"`
	CouponRefundFee4 int64 `json:"coupon_refund_fee_4;omitempty" query:"coupon_refund_fee_4;omitempty"`

	CouponRefundCount1 int64 `json:"coupon_refund_count_1;omitempty" query:"coupon_refund_count_1;omitempty"`
	CouponRefundCount2 int64 `json:"coupon_refund_count_2;omitempty" query:"coupon_refund_count_2;omitempty"`
	CouponRefundCount3 int64 `json:"coupon_refund_count_3;omitempty" query:"coupon_refund_count_3;omitempty"`
	CouponRefundCount4 int64 `json:"coupon_refund_count_4;omitempty" query:"coupon_refund_count_4;omitempty"`

	CouponRefundId01 string `json:"coupon_refund_id_0_1;omitempty" query:"coupon_refund_id_0_1;omitempty"`
	CouponRefundId02 string `json:"coupon_refund_id_0_2;omitempty" query:"coupon_refund_id_0_2;omitempty"`
	CouponRefundId03 string `json:"coupon_refund_id_0_3;omitempty" query:"coupon_refund_id_0_3;omitempty"`
	CouponRefundId04 string `json:"coupon_refund_id_0_4;omitempty" query:"coupon_refund_id_0_4;omitempty"`

	CouponRefundFee01 int64 `json:"coupon_refund_fee_0_1;omitempty" query:"coupon_refund_fee_0_1;omitempty"`
	CouponRefundFee02 int64 `json:"coupon_refund_fee_0_2;omitempty" query:"coupon_refund_fee_0_2;omitempty"`
	CouponRefundFee03 int64 `json:"coupon_refund_fee_0_3;omitempty" query:"coupon_refund_fee_0_3;omitempty"`
	CouponRefundFee04 int64 `json:"coupon_refund_fee_0_4;omitempty" query:"coupon_refund_fee_0_4;omitempty"`

	RefundStatus1 string `json:"refund_status_1;omitempty" query:"refund_status_1;omitempty"`
	RefundStatus2 string `json:"refund_status_2;omitempty" query:"refund_status_2;omitempty"`
	RefundStatus3 string `json:"refund_status_3;omitempty" query:"refund_status_3;omitempty"`
	RefundStatus4 string `json:"refund_status_4;omitempty" query:"refund_status_4;omitempty"`

	RefundAccount1 string `json:"refund_account_1;omitempty" query:"refund_account_1;omitempty"`
	RefundAccount2 string `json:"refund_account_2;omitempty" query:"refund_account_2;omitempty"`
	RefundAccount3 string `json:"refund_account_3;omitempty" query:"refund_account_3;omitempty"`
	RefundAccount4 string `json:"refund_account_4;omitempty" query:"refund_account_4;omitempty"`

	RefundRecvAccout1 string `json:"refund_recv_accout_1;omitempty" query:"refund_recv_accout_1;omitempty"`
	RefundRecvAccout2 string `json:"refund_recv_accout_2;omitempty" query:"refund_recv_accout_2;omitempty"`
	RefundRecvAccout3 string `json:"refund_recv_accout_3;omitempty" query:"refund_recv_accout_3;omitempty"`
	RefundRecvAccout4 string `json:"refund_recv_accout_4;omitempty" query:"refund_recv_accout_4;omitempty"`

	RefundSuccessTime1 string `json:"refund_success_time_1;omitempty" query:"refund_success_time_1;omitempty"`
	RefundSuccessTime2 string `json:"refund_success_time_2;omitempty" query:"refund_success_time_2;omitempty"`
	RefundSuccessTime3 string `json:"refund_success_time_3;omitempty" query:"refund_success_time_3;omitempty"`
	RefundSuccessTime4 string `json:"refund_success_time_4;omitempty" query:"refund_success_time_4;omitempty"`
}

type NotifyDto struct {
	ReturnCode string `xml:"return_code,omitempty" json:"return_code,omitempty"  query:"return_code,omitempty"`
	ReturnMsg  string `xml:"return_msg,omitempty" json:"return_msg,omitempty"  query:"return_msg,omitempty"`
	AppId      string `xml:"appid,omitempty" json:"appid,omitempty" query:"appid,omitempty"`
	MchId      string `xml:"mch_id,omitempty" json:"mch_id,omitempty" query:"mch_id,omitempty"`
	SubAppId   string `xml:"sub_appid,omitempty" json:"sub_appid,omitempty" query:"sub_appid,omitempty"`

	SubMchId   string `xml:"sub_mch_id,omitempty" json:"sub_mch_id,omitempty" query:"sub_mch_id,omitempty"`
	DeviceInfo string `xml:"device_info,omitempty" json:"device_info,omitempty" query:"device_info,omitempty"`
	NonceStr   string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty" query:"nonce_str,omitempty"`
	Sign       string `xml:"sign,omitempty" json:"sign,omitempty" query:"sign,omitempty"`
	ResultCode string `xml:"result_code,omitempty" json:"result_code,omitempty" query:"result_code,omitempty"`

	ErrCode     string `xml:"err_code,omitempty" json:"err_code,omitempty" query:"err_code,omitempty"`
	ErrCodeDes  string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty" query:"err_code_des,omitempty"`
	OpenId      string `xml:"openid,omitempty" json:"openid,omitempty" query:"openid,omitempty"`
	IsSubscribe string `xml:"is_subscribe,omitempty" json:"is_subscribe,omitempty"  query:"is_subscribe,omitempty"`
	SubOpenId   string `xml:"sub_openid,omitempty" json:"sub_openid,omitempty"  query:"sub_openid,omitempty"`

	SubIsSubscribe string `xml:"sub_is_subscribe,omitempty" json:"sub_is_subscribe,omitempty" query:"sub_is_subscribe,omitempty"`
	TradeType      string `xml:"trade_type,omitempty" json:"trade_type,omitempty" query:"trade_type,omitempty"`
	BankType       string `xml:"bank_type,omitempty" json:"bank_type,omitempty" query:"bank_type,omitempty"`
	TotalFee       int64  `xml:"total_fee,omitempty" json:"total_fee,omitempty" query:"total_fee,omitempty"` //int64
	FeeType        string `xml:"fee_type,omitempty" json:"fee_type,omitempty" query:"fee_type,omitempty"`

	CashFee            int64  `xml:"cash_fee,omitempty" json:"cash_fee,omitempty" query:"cash_fee,omitempty"` //int64
	CashFeeType        string `xml:"cash_fee_type,omitempty" json:"cash_fee_type,omitempty" query:"cash_fee_type,omitempty"`
	SettlementTotalFee int64  `xml:"settlement_total_fee,omitempty" json:"settlement_total_fee,omitempty" query:"settlement_total_fee,omitempty"` //int64
	CouponFee          int64  `xml:"coupon_fee,omitempty" json:"coupon_fee,omitempty" query:"coupon_fee,omitempty"`                               //int64
	CouponCount        int64  `xml:"coupon_count,omitempty" json:"coupon_count,omitempty" query:"coupon_count,omitempty"`                         //int64

	TransactionId string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty" query:"transaction_id,omitempty"`
	OutTradeNo    string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty" query:"out_trade_no,omitempty"`
	Attach        string `xml:"attach,omitempty" json:"attach,omitempty" query:"attach,omitempty"`
	TimeEnd       string `xml:"time_end,omitempty" json:"time_end,omitempty" query:"time_end,omitempty"`
}

type SceneInfoDto struct {
	Id       string `json:"id,omitempty" query:"id,omitempty"`
	Name     string `json:"name,omitempty" query:"name,omitempty"`
	AreaCode string `json:"area_code,omitempty" query:"area_code,omitempty"`
	Address  string `json:"address,omitempty" query:"address,omitempty"`
}
