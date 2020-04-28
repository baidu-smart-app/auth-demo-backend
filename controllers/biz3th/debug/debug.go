package debug

import (
	"auth-demo-backend/data/auth"
	"auth-demo-backend/data/pay"
	"auth-demo-backend/lib/httpserver"
)

func Debug(ctx *httpserver.Context) interface{} {
	user, _ := auth.DefaultAuth.GetAllUser()
	order, _ := pay.DefaulPay.GetAllOrders()

	return map[string]interface{}{
		"user":  user,
		"order": order,
	}
}
