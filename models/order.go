package models

type Order struct {
	ID 					string 		`json:"id"`
	Symbol				string		`json:"symbol"`
	Side 				string		`json:"side"` // buy/sell
	Type 				string 		`json:"type"` // limit/ market
	Price 				float64 	`json:"price,omitempty"`
	InitialQuantity 	float64		`json:"initialQuantity"`
	RemainigQuantity 	float64 	`json:"remainigQuantity"`
	Status 				float64 	`json:"satus"` // open/filled/canceled
}