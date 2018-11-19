package pay

import (
	paydata "icode.baidu.com/baidu/searchbox/ma-srvdemo/data/pay"
)

func SuccCallBack(params *paydata.SuccCallBackParam) error {
	return paydata.DefaulPay.PaySucc(params)
}
