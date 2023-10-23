package domain

type Item struct {
	ID          int32   `json:"-"`
	ChrtId      int32   `json:"chrt_id" example:"9934930"`
	TrackNumber string  `json:"track_number" example:"WBILMTESTTRACK"` // get it from order
	Price       float64 `json:"price" example:"453"`
	Rid         string  `json:"rid" example:"ab4219087a764ae0btest"`
	Name        string  `json:"name" example:"Mascaras"`
	Sale        float64 `json:"sale" example:"30"`
	Size        string  `json:"size" example:"0"`
	TotalPrice  float64 `json:"total_price" example:"317"`
	NmId        int32   `json:"nm_id" example:"2389212"`
	Brand       string  `json:"brand" example:"Vivienne Sabo"`
	Status      int16   `json:"status" example:"202"`
	OrderUid    string  `json:"order_uid" example:"b563feb7b2b84b6test"`
}
