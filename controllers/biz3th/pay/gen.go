package pay

import (
	"strconv"
	"time"

	"icode.baidu.com/baidu/searchbox/ma-srvdemo/lib/httpserver"
	paymodel "icode.baidu.com/baidu/searchbox/ma-srvdemo/model/pay"
	"icode.baidu.com/baidu/searchbox/ma-srvdemo/resource"
)

// Gen 生成订单
func Gen(ctx *httpserver.Context) interface{} {
	orderInfo := paymodel.NewOrderInfo(resource.C.Pay)

	mockData(orderInfo)

	result, err := orderInfo.Render(resource.C.SelfRsaPrivKey)
	if err != nil {
		ctx.Warning(err)
		return err
	}

	return result
}

func mockData(order *paymodel.OrderInfo) {
	mockID := mockOrderID()
	order.TpOrderID, order.BizInfo.TpData.TpOrderID = mockID, mockID

	// 每个订单1分钱
	order.TotalAmount, order.BizInfo.TpData.TotalAmount = "1", "1"

	order.BizInfo.TpData.DisplayData = &paymodel.TpDataDisplayData{
		CashierTopBlocks: []paymodel.Rows{
			paymodel.Rows{
				&paymodel.Row{"订单名称", "智能小程序支付实例" + mockID},
				&paymodel.Row{"数量", "1"},
				&paymodel.Row{"订单金额", "0.01元"},
			},
			paymodel.Rows{
				&paymodel.Row{"服务地址", "北京市海淀区上地十街10号百度大厦"},
			},
		},
	}

	order.BizInfo.TpData.DetailSubTitle = "支付示例"
	order.DealTitle = "支付示例"
}

func mockOrderID() string {
	// tpOrderId生成规则：(当前UnixTime - 2018-04-12 15:00:00的UnixTime) * 1000 + 当前UnixTime的Millissecond
	t := time.Now()
	intTpOrderID := (t.Unix()-1523516400)*1000 + t.UnixNano()%1000

	return strconv.FormatInt(intTpOrderID, 10)
}
