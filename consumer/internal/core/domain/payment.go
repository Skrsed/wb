package domain

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type Payment struct {
	ID           int      `json:"-"`
	Transaction  string   `json:"transaction" example:"b563feb7b2b84b6test" validate:"required,max=50"`
	RequestId    *string  `json:"request_id" validate:"max=100"`
	Currency     string   `json:"currency" validate:"required"`
	Provider     string   `json:"provider" example:"wbpay" validate:"required,max=50"`
	Amount       float64  `json:"amount" example:"1817" validate:"required,max=999999999"`
	PaymentDt    UnixTime `json:"payment_dt" example:"1637907727" validate:"required"`
	Bank         string   `json:"bank" example:"alpha" validate:"required,max=50"`
	DeliveryCost float64  `json:"delivery_cost" example:"1500" validate:"required,max=999999999"`
	GoodsTotal   float64  `json:"goods_total" example:"317" validate:"required,max=999999999"`
	CustomFee    float64  `json:"custom_fee" example:"0" validate:"required,max=999999999"`
	OrderUid     string   `json:"-"`
}

// Maybe some of theese needs to be persist on other lvl,
// e.g. unmarshal is may be an app lvl and scan is rep lvl
// but we'll be keep it here for now
type UnixTime struct {
	time.Time
}

func (t UnixTime) Value() (driver.Value, error) {
	if t.Time.IsZero() {
		return nil, nil
	}
	return t.Time.UTC().Format(time.RFC3339), nil
}

func (u *UnixTime) UnmarshalJSON(b []byte) error {
	var timestamp int64
	err := json.Unmarshal(b, &timestamp)
	if err != nil {
		return err
	}
	u.Time = time.Unix(timestamp, 0).UTC()
	return nil
}

func (c *UnixTime) Scan(v interface{}) error {
	timeV, ok := v.(time.Time)
	if !ok {
		return errors.New("invalid time format")
	}
	c.Time = timeV
	_, offset := c.Time.Zone()
	c.Time = c.Time.Add(time.Second * time.Duration(offset))
	c.Time = c.Time.Round(time.Second).UTC()
	return nil
}
