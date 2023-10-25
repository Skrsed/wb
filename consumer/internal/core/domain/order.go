package domain

import (
	"time"
)

type Order struct {
	Uid               string    `json:"order_uid" example:"b563feb7b2b84b6test" validate:"required"` // 19 sumbols uuid? // maybe type of uuid
	TrackNumber       string    `json:"track_number" example:"WBILMTESTTRACK" validate:"max=100"`    // 14 sub
	Entry             string    `json:"entry" example:"WBIL" validate:"required,max=50"`
	Delivery          Delivery  `json:"delivery" validate:"required"`
	Payment           Payment   `json:"payment" validate:"required"`
	Items             []*Item   `json:"items" validate:"required"`
	Locale            string    `json:"locale" example:"en" validate:"required"`
	InternalSignature string    `json:"internal_signature" validate:"required"`
	CustomerId        string    `json:"customer_id" example:"test" validate:"required"`
	DeliveryService   string    `json:"delivery_service" example:"meest" validate:"max=100"`
	Shardkey          string    `json:"shardkey" example:"9" validate:"required,max=50"`
	SmId              int       `json:"sm_id" example:"99" validate:"required"`
	DateCreated       time.Time `json:"date_created" example:"2021-11-26T06:22:19Z" validate:"required"`
	OofShard          string    `json:"oof_shard" example:"1" validate:"required"`
}
