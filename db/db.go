package db

import (
	"database/sql"
	"fmt"
	"golangOrderMatchingSystem/models"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	fmt.Printf("dsn : %s", dsn)

	var errOpen error
	DB, errOpen = sql.Open("mysql", dsn)
	if errOpen != nil {
		log.Fatal(errOpen)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to the database")

	InitializeDatabase()
}

func GetOrderById(id string) (*models.Order, error) {
	row := DB.QueryRow(`
        SELECT id, symbol, side, type, price, initialQuantity, remainingQuantity, status
        FROM orders
        WHERE id = ?
    `, id)

	var order models.Order
	err := row.Scan(
		&order.ID,
		&order.Symbol,
		&order.Side,
		&order.Type,
		&order.Price,
		&order.InitialQuantity,
		&order.RemainingQuantity,
		&order.Status,
	)

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func GetOrderBookEntries(symbol string, side string) ([]map[string]interface{}, error) {
	rows, err := DB.Query(`
        SELECT id  , type , price, initialQuantity , remainingQuantity , SUM(remainingQuantity) as totalQuantity
        FROM orders
        WHERE symbol = ? AND side = ? AND status = 'open'
        GROUP BY price
        ORDER BY price DESC`, symbol, side)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []map[string]interface{}
	for rows.Next() {
		var price sql.NullFloat64
		var quantity float64

		if err := rows.Scan(&price, &quantity); err != nil {
			return nil, err
		}

		if !price.Valid {
			continue // skip market orders if any (no price)
		}

		entry := map[string]interface{}{
			"price":    price.Float64,
			"quantity": quantity,
		}
		entries = append(entries, entry)
	}

	// Reverse for sell orders to show ascending price
	if side == "sell" {
		for i, j := 0, len(entries)-1; i < j; i, j = i+1, j-1 {
			entries[i], entries[j] = entries[j], entries[i]
		}
	}

	return entries, nil
}
