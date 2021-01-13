package pay

import (
	"encoding/json"
	"net/url"
)

type Config struct {
	AppKey    string `json:"app_key"`
	AppID     string `json:"app_id"`
	DealID    string `json:"deal_id"`
	DealTitle string `json:"deal_title"`
}

/*
BizInfo 业务数据 糯米支付接口入参
https://dianshang.baidu.com/platform/doclist/index.html#!/doc/nuomiplus_1_guide/mini_program_cashier/parameter.md
{
	"tpData": {
		"appKey": "MMMabc",
		"dealId": "470193086",
		"tpOrderId": "3028903626",
		"rsaSign": "",
		"totalAmount": "11300",
		"payResultUrl": "",
		"returnData": {
			"bizKey1": "第三方的字段1取值",
			"bizKey2": "第三方的字段2取值"
		},
		"displayData": {
			"cashierTopBlock": [
				[{
					"leftCol": "订单名称",
					"rightCol": "爱鲜蜂"
				}, {
					"leftCol": "数量",
					"rightCol": "1"
				}, {
					"leftCol": "小计",
					"rightCol": "113"
				}],
				[{
					"leftCol": "服务地址",
					"rightCol": "北京市海淀区中关村南大街5号百度大厦"
				}, {
					"leftCol": "服务时间",
					"rightCol": "2015/05/20 13:30-14:00"
				}, {
					"leftCol": "服务人员",
					"rightCol": "娜娜"
				}]
			]
		},
		"dealTitle": "爱鲜蜂",
		"dealSubTitle": "爱鲜蜂副标题",
		"dealThumbView": "http://i1.sinaimg.cn/lx/news/2008-12-30/U1043P622T1D29849F15DT20081230120718.jpg"
	},
	"orderDetailData": {
		"displayData": {
			"tpOrderInfoBlock": {
				"title": "服务明细",
				"content": [{
					"leftCol": "服务地址",
					"rightCol": "北京市海淀区上地东北旺西路10号百度科技园"
				}, {
					"leftCol": "服务时间",
					"rightCol": "2015/06/06 13:00-15:00"
				}]
			}
		}
	}
}
*/
type BizInfo struct {
	TpData          *BizInfoTpData         `json:"tpData"`
	OrderDetailData *TpDataOrderDetailData `json:"orderDetailData"`
}

type BizInfoTpData struct {
	AppKey         string             `json:"appKey"`
	DealID         string             `json:"dealID"`
	TpOrderID      string             `json:"tpOrderID"`
	RsaSign        string             `json:"rsaSign"`
	TotalAmount    string             `json:"totalAmount"`
	PayResultUrl   string             `json:"payResultUrl"`
	ReturnData     map[string]string  `json:"returnData"`
	DealTitle      string             `json:"dealTitle"`
	DetailSubTitle string             `json:"detailSubTitle"`
	DealTumbView   string             `json:"dealTumbView"`
	DisplayData    *TpDataDisplayData `json:"displayData"`
}

type TpDataDisplayData struct {
	CashierTopBlocks []Rows `json:"cashierTopBlock"`
}

type TpDataOrderDetailData struct {
	DispalyData map[string]*Rows `json:"displayData"`
}

type Rows []*Row

type DispalyBlock struct {
	Title   string `json:"title"`
	Content Rows   `json:"content"`
}

type Row struct {
	LeftCol  string `json:"leftCol"`
	RightCol string `json:"rightCol"`
}

func NewOrderInfo(config *Config) *OrderInfo {
	return &OrderInfo{
		DealID:    config.DealID,
		AppKey:    config.AppKey,
		DealTitle: config.DealTitle,
		BizInfo: &BizInfo{
			TpData: &BizInfoTpData{
				DealID:    config.DealID,
				AppKey:    config.AppKey,
				DealTitle: config.DealTitle,
			},
		},
	}
}

/*
OrderInfo 订单数据 小程序接口数据
https://smartprogram.baidu.com/docs/develop/api/open_payment/#requestPolymerPayment

{
    "dealId": "470193086",
	"appKey": "MMMabc",
	"totalAmount": "11300",
	"tpOrderId": "3028903626",
	"dealTitle": "爱鲜蜂",
	"rsaSign": '',
	"bizInfo": ''
}
*/
type OrderInfo struct {
	DealID      string
	AppKey      string
	TotalAmount string
	TpOrderID   string
	DealTitle   string
	RsaSign     string
	BizInfo     *BizInfo
}

// Render 序列化订单信息
func (oi *OrderInfo) Render(rsaPrivateKey []byte) (map[string]string, error) {
	kvs := url.Values{
		"dealId":    []string{oi.DealID},
		"appKey":    []string{oi.AppKey},
		"tpOrderId": []string{oi.TpOrderID},
	}

	cipherText, err := Sign(kvs.Encode(), rsaPrivateKey)
	if err != nil {
		return nil, err
	}

	bizInfo := ""

	if oi.BizInfo != nil {
		oi.BizInfo.TpData.RsaSign = string(cipherText)

		bs, err := json.Marshal(oi.BizInfo)
		if err != nil {
			return nil, err
		}
		bizInfo = string(bs)
	}

	return map[string]string{
		"dealId":          oi.DealID,
		"appKey":          oi.AppKey,
		"totalAmount":     oi.TotalAmount,
		"tpOrderId":       oi.TpOrderID,
		"dealTitle":       oi.DealTitle,
		"rsaSign":         string(cipherText),
		"bizInfo":         bizInfo,
		"signFieldsRange": "1",
	}, nil
}
