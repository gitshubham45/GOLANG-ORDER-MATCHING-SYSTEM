package models

type Order struct {
	ID 					string 		`json:"id"`
	Symbol				string		`json:"symbol"`
	Side 				string		`json:"side"` // buy/sell
	Type 				string 		`json:"type"` // limit/ market
	Price 				float64 	`json:"price,omitempty"`
	InitialQuantity 	float64		`json:"initialQuantity"`
	RemainingQuantity 	float64 	`json:"remainigQuantity"`
	Status 				string 		`json:"status"` // open/filled/canceled
	CreatedAt 			string		`json:"createdAt"`
	UpdatedAt 			string 		`json:"updatedAt"`
}