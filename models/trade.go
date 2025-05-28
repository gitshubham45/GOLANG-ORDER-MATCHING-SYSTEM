package models

type Trade struct {
	BuyOrderId 		string		`json:"buyOrderId"`
	SellOrderId		string		`json:"sellOrderId"`
	Symbol 			string  	`json:"symbol"`
	Price 			float64 	`json:"price,omitempty"`
	Quantity 		string 		`json:"quantity"`
	MatchedAt		string 		`json:"matchedAt"`
}