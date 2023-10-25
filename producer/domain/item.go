package domain

type Item struct {
	ID          int     `json:"-"`
	ChrtId      int     `json:"chrt_id" example:"9934930" fake:"{number:1, 2147483647}"`
	TrackNumber string  `json:"track_number" example:"WBILMTESTTRACK"` // get it from order
	Price       float64 `json:"price" example:"453"`
	Rid         string  `json:"rid" example:"ab4219087a764ae0btest"`
	Name        string  `json:"name" example:"Mascaras"`
	Sale        float64 `json:"sale" example:"30"`
	Size        string  `json:"size" example:"0"`
	TotalPrice  float64 `json:"total_price" example:"317"`
	NmId        int     `json:"nm_id" example:"2389212" fake:"{number:1,99999999999999}"`
	Brand       string  `json:"brand" example:"Vivienne Sabo"`
	Status      int16   `json:"status" example:"202"`
	OrderUid    string  `json:"order_uid" example:"b563feb7b2b84b6test"`
}
