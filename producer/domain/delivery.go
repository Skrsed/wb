package domain

type Delivery struct {
	ID       int32  `json:"-"`
	Name     string `json:"name" example:"Test Testov"`
	Phone    string `json:"phone" example:"+9720000000" fake:"+7{phone}"`
	Zip      int    `json:"zip" example:"2639809" fake:"{number:5,9}"`
	City     string `json:"city" example:"Kiryat Mozkin"`
	Address  string `json:"address" example:"Ploshad Mira 15"`
	Region   string `json:"region" example:"Kraiot"`
	Email    string `json:"email" example:"test@gmail.com" fake:"{email}"`
	OrderUid string `json:"order_uid"`
}
