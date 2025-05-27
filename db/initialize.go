package db

import (
	"log"
)

const orderSchema = `
CREATE TABLE IF NOT EXISTS orders (
    id VARCHAR(36) PRIMARY KEY,
    symbol VARCHAR(10) NOT NULL,
    side ENUM('buy', 'sell') NOT NULL,
    type ENUM('limit', 'market') NOT NULL,
    price DECIMAL(15,2),
    initialQuantity DECIMAL(15,8) NOT NULL,
    remainingQuantity DECIMAL(15,8) NOT NULL,
    status ENUM('open', 'filled', 'canceled') NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME ON UPDATE CURRENT_TIMESTAMP
);
`

const tradeSchema = `
CREATE TABLE IF NOT EXISTS trades (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    buyOrderId VARCHAR(36) NOT NULL,
    sellOrderId VARCHAR(36) NOT NULL,
    symbol VARCHAR(10) NOT NULL,
    price DECIMAL(15,2) NOT NULL,
    quantity DECIMAL(15,8) NOT NULL,
    matched_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
`

func InitializeDatabase() {
	log.Println("Initializing database schema...")
	_, err := DB.Exec(orderSchema)
	if err != nil {
		log.Fatalf("Error initializing database schema: %v", err)
		return
	}

	_, err2 := DB.Exec(tradeSchema)
	if err2 != nil {
		log.Fatalf("Error initializing database schema: %v", err2)
	}

	log.Println("Database schema initialized.")
}
