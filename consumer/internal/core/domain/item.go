package domain

type Item struct {
	ID          int     `json:"-"`
	ChrtId      int     `json:"chrt_id" example:"9934930" validate:"required"`
	TrackNumber string  `json:"track_number" example:"WBILMTESTTRACK" validate:"max=100"` // get it from order
	Price       float64 `json:"price" example:"453" validate:"required,max=999999999"`
	Rid         string  `json:"rid" example:"ab4219087a764ae0btest" validate:"required,max=50"`
	Name        string  `json:"name" example:"Mascaras" validate:"required,max=50"`
	Sale        float64 `json:"sale" example:"30" validate:"required,min=0"`
	Size        string  `json:"size" example:"0" validate:"required,max=10"`
	TotalPrice  float64 `json:"total_price" example:"317" validate:"required,max=999999999"`
	NmId        int     `json:"nm_id" example:"2389212" validate:"required"`
	Brand       string  `json:"brand" example:"Vivienne Sabo" validate:"required,max=100"`
	Status      int16   `json:"status" example:"202" validate:"required,max=999"`
	OrderUid    string  `json:"order_uid" example:"b563feb7b2b84b6test"`
}
