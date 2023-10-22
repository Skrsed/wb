package domain

type Delivery struct {
	ID       string `json:"id" example:"123"`
	Name     string `json:"name" example:"Test Testov"`
	Phone    string `json:"phone" example:"+9720000000"` //TODO validate
	Zip      string `json:"zip" example:"2639809"`
	City     string `json:"city" example:"Kiryat Mozkin"`
	Address  string `json:"address" example:"Ploshad Mira 15"`
	Region   string `json:"region" example:"Kraiot"`
	Email    string `json:"email" example:"test@gmail.com"` //TODO validate
	OrderUid string `json:"order_uid" example:"b563feb7b2b84b6test"`
}
