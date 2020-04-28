package pay

import (
	"auth-demo-backend/lib/httpserver"
	paymodel "auth-demo-backend/model/pay"
	"auth-demo-backend/resource"
)

// Status 查询订单
// https://dianshang.baidu.com/platform/doclist/index.html#!/doc/nuomiplus_1_guide/mini_program_cashier/standard_interface/status_search.md
func Status(ctx *httpserver.Context) interface{} {
	// 智能小程序现有的封装没办法得到订单创建成功后 orderID-tpOrderID 的映射
	orderID := ctx.QueryString("tp_order_id")

	data, err := paymodel.Status(ctx, orderID, resource.C.Pay, resource.C.SelfRsaPrivKey)
	if err != nil {
		ctx.Warning(err)
		return err
	}

	return data
}
