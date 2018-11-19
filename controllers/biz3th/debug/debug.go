package debug

import (
	"icode.baidu.com/baidu/searchbox/ma-srvdemo/data/auth"
	"icode.baidu.com/baidu/searchbox/ma-srvdemo/data/pay"
	"icode.baidu.com/baidu/searchbox/ma-srvdemo/lib/httpserver"
)

func Debug(ctx *httpserver.Context) interface{} {
	user, _ := auth.DefaultAuth.GetAllUser()
	order, _ := pay.DefaulPay.GetAllOrders()

	return map[string]interface{}{
		"user":  user,
		"order": order,
	}
}
