package models

type Trade struct {
	BuyOrderId 		string		`json:"buyOrderId"`
	SellOrderId		string		`json:"sellOrderId"`
	Symbol 			string  	`json:"symbol"`
	Price 			string 		`json:"price"`
	Quantity 		string 		`json:"quantity"`
}