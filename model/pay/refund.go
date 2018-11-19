package pay

import (
	paydata "icode.baidu.com/baidu/searchbox/ma-srvdemo/data/pay"
	"icode.baidu.com/baidu/searchbox/ma-srvdemo/lib/httpserver"
	"icode.baidu.com/baidu/searchbox/ma-srvdemo/lib/httputil"
)

// Refund 退款申请
// https://dianshang.baidu.com/platform/doclist/index.html#!/doc/nuomiplus_1_guide/mini_program_cashier/standard_interface/apply_refund.md
func Refund(ctx *httpserver.Context, tpOrderID string, conf *Config, pem []byte) (*RefundQueryData, error) {
	paySuccInfo, err := paydata.DefaulPay.GetOrder(tpOrderID)
	if err != nil {
		ctx.Warning(err)
		return nil, err
	}

	if paySuccInfo == nil {
		return nil, nil
	}

	cli := httputil.NewClient().
		SetPath("https://nop.nuomi.com/nop/server/rest").
		SetGetParam("method", "nuomi.cashier.applyorderrefund").
		SetGetParam("orderId", paySuccInfo.OrderID).
		SetGetParam("userId", paySuccInfo.UserID).
		SetGetParam("refundType", "1").
		SetGetParam("refundReason", "退款demo").
		SetGetParam("tpOrderId", paySuccInfo.TpOrderID).
		SetGetParam("appKey", conf.AppKey)

	sign, err := Sign(cli.Req.Form.Encode(), pem)
	if err != nil {
		ctx.Warning(err)
		return nil, err
	}

	rsp := &RefundQueryRsp{}

	e := cli.SetGetParam("sign", sign).
		Do().
		RspJson(rsp, func(data interface{}) error {
			if data.(*RefundQueryRsp).ErrNO != 0 {
				return httputil.BadContent
			}
			return nil
		}).
		Error()

	if e != nil {
		return nil, e
	}
	return rsp.Data, nil
}

type RefundQueryRsp struct {
	ErrNO int64  `json:"errno"`
	Msg   string `json:"msg"`

	Data *RefundQueryData `json:"data"`
}

type RefundQueryData struct {
	RefundBatchID  string `json:"refundBatchId"`
	RefundPayMoney string `json:"refundPayMoney"`
}
