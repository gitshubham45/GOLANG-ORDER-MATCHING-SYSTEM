package db

import (
	"database/sql"
	"golangOrderMatchingSystem/models"
	"log"
)

func SaveOrder(order models.Order) {
	query := `
		INSERT INTO orders(
			id , symbol, side ,type , price , initialQuantity , remainingQuantity , status
		) VALUES (? , ? , ? , ? , ? , ? , ? , ? )
	`

	_, err := DB.Exec(query,
		order.ID,
		order.Symbol,
		order.Side,
		order.Type,
		order.Price,
		order.InitialQuantity,
		order.RemainingQuantity,
		order.Status,
	)

	if err != nil {
		log.Println("Error saving order ", err)
	}
}

func UpdateOrder(order models.Order) {
	query := `
		UPDATE orders SET 
			remainingQuantity = ? , status = ?
		WHERE id = ?
	`

	_, err := DB.Exec(query, order.RemainingQuantity, order.Status, order.ID)
	if err != nil {
		log.Println("Error updating order: ", err)
	}
}

func GetOpenOrders(symbol string, side string) []models.Order {
	rows, err := DB.Query(`
		SELECT id , symbol , type , price , initialQuantity , remainingQuantity 
			FROM orders 
		WHERE symbol = ? AND side = ? AND status = 'open'
		ORDER BY price DESC , id ASC 
	`, symbol, side)

	if err != nil {
		log.Println("Error fetching orders:", err)
		return nil
	}

	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var o models.Order
		if err := rows.Scan(
			&o.ID, &o.Symbol, &o.Type, &o.Price,
			&o.InitialQuantity, &o.RemainingQuantity,
		); err != nil {
			log.Println("Scan error:", err)
			continue
		}
		orders = append(orders, o)
	}
	return orders
}

func LogTrade(buyID, sellID string, symbol string, price, quantity float64) {
	query := `
		INSERT INTO trades (buyOrderId, sellOrderId , symbol , price , quantity)
		VALUES ( ? , ? , ? , ? , ?)
	`
	_, err := DB.Exec(query, buyID, sellID, symbol, price, quantity)

	if err != nil {
		log.Println("Error logging trade:", err)
	}
}

func GetTradesBySymbol(symbol string) ([]models.Trade, error) {
	rows, err := DB.Query(`
        SELECT buyOrderId, sellOrderId, symbol, price, quantity
        FROM trades
        WHERE symbol = ?
        ORDER BY id DESC
    `, symbol)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var trades []models.Trade
	for rows.Next() {
		var t models.Trade
		if err := rows.Scan(&t.BuyOrderId, &t.SellOrderId, &t.Symbol, &t.Price, &t.Quantity); err != nil {
			return nil, err
		}
		trades = append(trades, t)
	}
	return trades, nil
}

func UpdateOrderStatus(id string, status string) error {
	result, err := DB.Exec("UPDATE orders SET status = ? WHERE id = ?", status, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
