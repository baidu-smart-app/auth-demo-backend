package pay

import (
	"encoding/json"
	"fmt"

	paydata "auth-demo-backend/data/pay"
	"auth-demo-backend/lib/httpserver"
	"auth-demo-backend/lib/httputil"
)

// Status 查询支付状态
// https://dianshang.baidu.com/platform/doclist/index.html#!/doc/nuomiplus_1_guide/mini_program_cashier/standard_interface/status_search.md
func Status(ctx *httpserver.Context, tpOrderID string, conf *Config, pem []byte) (*OrderQueryData, error) {
	paySuccInfo, err := paydata.DefaulPay.GetOrder(tpOrderID)
	if err != nil {
		ctx.Warning(err)
		return nil, err
	}

	if paySuccInfo == nil {
		return nil, nil
	}

	bs, _ := json.Marshal(paySuccInfo)
	fmt.Println(string(bs))

	cli := httputil.NewClient().
		SetPath("https://dianshang.baidu.com/platform/entity/openapi/queryorderdetail").
		SetGetParam("appId", conf.AppID).
		SetGetParam("appKey", conf.AppKey).
		SetGetParam("orderId", paySuccInfo.OrderID).
		SetGetParam("siteId", paySuccInfo.UserID)

	sign, err := Sign(cli.Req.Form.Encode(), pem)
	if err != nil {
		ctx.Warning(err)
		return nil, err
	}

	rsp := &OrderQueryRsp{}

	e := cli.SetGetParam("sign", sign).
		Do().
		RspJson(rsp, func(data interface{}) error {
			if data.(*OrderQueryRsp).ErrNO != 0 {
				return httputil.BadContent
			}
			return nil
		}).
		Error()

	if e != nil {
		return nil, e
	}

	return rsp.Data["data"], nil
}

type OrderQueryRsp struct {
	ErrNO       int64  `json:"errno"`
	ErrMsg      string `json:"errmsg"`
	Timestamp   int64  `json:"timestamp"`
	Cached      int8   `json:"cached"`
	ServerLogID string `json:"serverlogid"`

	Data map[string]*OrderQueryData `json:"data"`
}

type OrderQueryData map[string]*StatusInfo

type StatusInfo struct {
	StatusNum  int8   `json:"statusNum"`
	StatusDesc string `json:"statusDesc"`
}
