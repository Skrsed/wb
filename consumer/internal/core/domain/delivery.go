package domain

type Delivery struct {
	ID       int    `json:"-"`
	Name     string `json:"name" example:"Test Testov" validate:"required"`
	Phone    string `json:"phone" example:"+9720000000" validate:"required,e164"`
	Zip      int    `json:"zip" example:"2639809" validate:"required,min=5,max=9"`
	City     string `json:"city" example:"Kiryat Mozkin" validate:"required,max=50"`
	Address  string `json:"address" example:"Ploshad Mira 15" validate:"required,max=100"`
	Region   string `json:"region" example:"Kraiot" validate:"required,max=50"`
	Email    string `json:"email" example:"test@gmail.com" validate:"required,email"`
	OrderUid string `json:"order_uid"`
}
