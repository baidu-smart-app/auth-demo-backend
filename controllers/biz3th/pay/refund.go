package pay

import (
	"icode.baidu.com/baidu/searchbox/ma-srvdemo/lib/httpserver"
	"icode.baidu.com/baidu/searchbox/ma-srvdemo/lib/rsp"
	paymodel "icode.baidu.com/baidu/searchbox/ma-srvdemo/model/pay"
	"icode.baidu.com/baidu/searchbox/ma-srvdemo/resource"
)

// Refund 申请退款
// https://dianshang.baidu.com/platform/doclist/index.html#!/doc/nuomiplus_1_guide/mini_program_cashier/standard_interface/apply_refund.md
func Refund(ctx *httpserver.Context) interface{} {
	// 智能小程序现有的封装没办法得到订单创建成功后 orderID-tpOrderID 的映射
	tpOrderID := ctx.QueryString("tp_order_id")
	if tpOrderID == "" {
		return rsp.ParamIllegal
	}

	data, err := paymodel.Refund(ctx, tpOrderID, resource.C.Pay, resource.C.SelfRsaPrivKey)
	if err != nil {
		ctx.Warning(err.Error())
		return err
	}

	return data
}
