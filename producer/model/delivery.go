package model

type Delivery struct {
	name string `json:name example:Test Testov`
    phone string `json:phone example:+9720000000` //TODO validate
    zip string `json:zip example:2639809`
    city string `json:city example:Kiryat Mozkin`
    address string `json:address example:Ploshad Mira 15`
    region string `json:region example:Kraiot`
    email string `json:email example:test@gmail.com` //TODO validate
}