package pay

import (
	paydata "auth-demo-backend/data/pay"
)

func SuccCallBack(params *paydata.SuccCallBackParam) error {
	return paydata.DefaulPay.PaySucc(params)
}
