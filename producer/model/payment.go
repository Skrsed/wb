package model

type Payment struct {
	Transaction string `json:transaction example:b563feb7b2b84b6test`
	RequestId string `json:request_id`
	currency string `json:currency example:USD`
	Provider string `json:provider example:wbpay`
	Amount int `json:amount example:1817`
	PaymentDt int `json:payment_dt example:1637907727`
	Bank string `json:bank example:alpha`
	DeliveryCost int `json:delivery_cost example:1500`
	GoodsTotal int `json:goods_total example:317`
	CustomFee int `json:custom_fee example:0`
}