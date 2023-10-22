package domain

import (
	"time"
)

type Order struct {
	Uid               string    `json:"order_uid" example:"b563feb7b2b84b6test" validate:"required"` // 19 sumbols uuid? // maybe type of uuid
	TrackNumber       string    `json:"track_number" example:"WBILMTESTTRACK"`                       // 14 sub
	Entry             string    `json:"entry" example:"WBIL"`
	DeliveryId        int       `json:"delivery_id" example:"123"`
	Delivery          Delivery  `json:"delivery" validate:"required"`
	PaymentId         int       `json:"payment_id" example:"123"`
	Payment           Payment   `json:"payment" validate:"required"`
	Items             []*Item   `json:"items" validate:"required"`
	Locale            string    `json:"locale" example:"en"`
	InternalSignature string    `json:"internal_signature"`
	CustomerId        string    `json:"customer_id" example:"test"`
	DeliveryService   string    `json:"delivery_service" example:"meest"`
	Shardkey          int       `json:"shardkey" example:"9"`
	SmId              int       `json:"sm_id" example:"99"`
	DateCreated       time.Time `json:"date_created" example:"2021-11-26T06:22:19Z"`
	OofShard          int       `json:"oof_shard" example:"1"`
}
