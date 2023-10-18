package model

type Delivery struct {
	Name string `json:name example:Test Testov`
    Phone string `json:phone example:+9720000000` //TODO validate
    Zip string `json:zip example:2639809`
    City string `json:city example:Kiryat Mozkin`
    Address string `json:address example:Ploshad Mira 15`
    Region string `json:region example:Kraiot`
    Email string `json:email example:test@gmail.com` //TODO validate
}