package domain

import (
	"time"
)

type Order struct {
	Uid               string    `json:"order_uid" example:"b563feb7b2b84b6test" validate:"required"` // 19 sumbols uuid? // maybe type of uuid
	TrackNumber       string    `json:"track_number" example:"WBILMTESTTRACK"`                       // 14 sub
	Entry             string    `json:"entry" example:"WBIL"`
	Delivery          Delivery  `json:"delivery" validate:"required"`
	Payment           Payment   `json:"payment" validate:"required"`
	Items             []*Item   `json:"items" validate:"required"`
	Locale            string    `json:"locale" fake:"{randomstring:[en,ru,kz]}"`
	InternalSignature string    `json:"internal_signature"`
	CustomerId        string    `json:"customer_id" example:"test"`
	DeliveryService   string    `json:"delivery_service" example:"meest"`
	Shardkey          string    `json:"shardkey" example:"9"`
	SmId              int32     `json:"sm_id" example:"99" fake:"{number:1, 2147483647}"`
	DateCreated       time.Time `json:"date_created" fake:"{date}"`
	OofShard          string    `json:"oof_shard" example:"1"`
}
