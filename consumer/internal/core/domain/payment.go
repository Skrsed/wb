package domain

import "time"

type Payment struct {
	ID           int       `json:"-"`
	Transaction  string    `json:"transaction" example:"b563feb7b2b84b6test"`
	RequestId    *string   `json:"request_id"`
	Currency     string    `json:"currency"`
	Provider     string    `json:"provider" example:"wbpay"`
	Amount       float64   `json:"amount" example:"1817"`
	PaymentDt    time.Time `json:"payment_dt" example:"1637907727"`
	Bank         string    `json:"bank" example:"alpha"`
	DeliveryCost float64   `json:"delivery_cost" example:"1500"`
	GoodsTotal   float64   `json:"goods_total" example:"317"`
	CustomFee    float64   `json:"custom_fee" example:"0"`
	OrderUid     string    `json:"-"`
}
