package model

type Item struct {
	ChrtId int `json:chrt_id example:9934930`
	TrackNumber string `json:track_number example:WBILMTESTTRACK` //TODO DRY
	Price string `json:price example:453` // TODO decimal
	Rid string `json:rid example:ab4219087a764ae0btest`
	Name string `json:name example:Mascaras`
	Sale int `json:sale example:30`
	Size int `json:size example:0`
	TotalPrice int `json:total_price example:317`
	NmId int `json:nm_id example:2389212`
	Brand string `json:brand example:Vivienne Sabo`
	Status uint16 `json:status example:202`
}