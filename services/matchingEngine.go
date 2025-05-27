package services

import (
	"math"

	"golangOrderMatchingSystem/db"
	"golangOrderMatchingSystem/models"
)

func MatchIncomingOrder(newOrder models.Order) {
	var oppositeSide string
	var bestPriceCmp func(float64, float64) bool
	var matchPrice func(models.Order) float64

	switch newOrder.Side {
	case "buy":
		oppositeSide = "sell"
		bestPriceCmp = func(a, b float64) bool { return a <= b }
		matchPrice = func(o models.Order) float64 { return o.Price }
	case "sell" :
		oppositeSide = "buy"
		bestPriceCmp = func(a , b  float64) bool { return a >= b}
		matchPrice = func(o models.Order) float64 { return o.Price }
	default:
		return
	}

	existingOrders := db.GetOpenOrders(newOrder.Symbol , oppositeSide)

	for i := 0 ; i < len(existingOrders) ; i++ {
		existing := existingOrders[i]

		if newOrder.RemainingQuantity <= 0 {
			break
		}

		if !bestPriceCmp(matchPrice(existing), newOrder.Price){
			continue
		}

		matchedQty := math.Min(newOrder.RemainingQuantity , existing.RemainingQuantity)
		buyID, sellID := chooseBuySellIDs(newOrder, existing)

		db.LogTrade(
			buyID,
			sellID,
			newOrder.Symbol,
			existing.Price,
			matchedQty,
		)

		newOrder.RemainingQuantity -= matchedQty
		existing.RemainingQuantity -= matchedQty

		if existing.RemainingQuantity <= 0 {
			existing.Status = "filled"
		}else {
			existing.Status = "open"
		}
		db.UpdateOrder(existing)	
	}

	if newOrder.RemainingQuantity > 0 {
		if newOrder.Type == "limit" {
			newOrder.Status = "open"
		} else {
			newOrder.Status = "filled"
		}
	}else{
		newOrder.Status = "filled"
	}
}

func chooseBuySellIDs(newOrder , existing models.Order) (string , string ) {
	if newOrder.Side == "buy" {
		return newOrder.ID , existing.ID
	}
	return existing.ID , newOrder.ID
}
