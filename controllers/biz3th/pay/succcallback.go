package pay

import (
	paydata "auth-demo-backend/data/pay"
	"auth-demo-backend/lib/httpserver"
	"auth-demo-backend/lib/rsp"
	paymodel "auth-demo-backend/model/pay"
	"auth-demo-backend/resource"
)

// SuccCallBack 支付成功回到
// https://dianshang.baidu.com/platform/doclist/index.html#!/docnuomiplus_1_guide/mini_program_cashier/standard_interface/push_notice.md
// count=1&dealId=2500207776&giftCardMoney=0&hbBalanceMoney=0&hbMoney=0&orderId=2645985458&partnerId=0&payMoney=1&payTime=1542124311&payType=1117&promoDetail=&promoMoney=0&returnData=&rsaSign=oOFlK3oOlVSbfLfgL%2FbgqO2rDIxt%2FVFQdxYeoZlj9WHXMrAW4hzyGwlplNJ%2FZzip9h0MRUSDwLQxDDhkFF1s59%2FzY%2BInhB6e%2FNyR%2Bpaurl2Sj1%2Fdj4Q0rbkuqMdKhU1A4A4jmG3fxTTx7FqExOcBjSxz6peLtmXyHDKhshoYYKc%3D&status=2&totalMoney=1&tpOrderId=18607895241&unitPrice=1&userId=3269234686470
func SuccCallBack(ctx *httpserver.Context) interface{} {

	req := &paydata.SuccCallBackParam{
		UserID:         ctx.FormString("userId"),
		OrderID:        ctx.FormString("orderId"),
		UnitPrice:      ctx.FormInt64("unitPrice"),
		Count:          ctx.FormInt64("count"),
		TotalMoney:     ctx.FormInt64("totalMoney"),
		PayMoney:       ctx.FormInt64("payMoney"),
		PromoMoney:     ctx.FormInt64("promoMoney"),
		HbMoney:        ctx.FormInt64("hbMoney"),
		HbBalanceMoney: ctx.FormInt64("hbBalanceMoney"),
		GiftCardMoney:  ctx.FormInt64("giftCardMoney"),
		DealID:         ctx.FormString("dealId"),
		PayTime:        ctx.FormInt64("payTime"),
		PromoDetail:    ctx.FormString("promoDetail"),
		PayType:        ctx.FormInt64("payType"),
		PartnerID:      ctx.FormInt64("partnerId"),
		Status:         ctx.FormInt64("status"),
		TpOrderID:      ctx.FormString("tpOrderId"),
		ReturnData:     ctx.FormString("returnData"),
		RsaSign:        ctx.FormString("rsaSign"),
	}

	form := ctx.Req.PostForm
	sign := form.Get("rsaSign")
	form.Del("rsaSign")

	err := paymodel.CheckSign(form.Encode(), sign, resource.C.PlatformRsaPubKey)
	if err != nil {
		return err
	}

	err = paymodel.SuccCallBack(req)
	if err != nil {
		return err
	}

	return &rsp.Raw{
		Data: map[string]interface{}{
			"errno": 0,
			"msg":   "success",
			"data": map[string]int{
				"isConsumed": 2,
			},
		},
	}
}
