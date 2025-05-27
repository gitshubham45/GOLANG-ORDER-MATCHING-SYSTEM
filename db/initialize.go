package db

import (
    "log"
)

const schema = `
CREATE TABLE IF NOT EXISTS orders (
    id VARCHAR(36) PRIMARY KEY,
    symbol VARCHAR(10) NOT NULL,
    side ENUM('buy', 'sell') NOT NULL,
    type ENUM('limit', 'market') NOT NULL,
    price DECIMAL(15,2),
    initial_quantity DECIMAL(15,8) NOT NULL,
    remaining_quantity DECIMAL(15,8) NOT NULL,
    status ENUM('open', 'filled', 'canceled') NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME ON UPDATE CURRENT_TIMESTAMP
);

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
    _, err := DB.Exec(schema)
    if err != nil {
        log.Fatalf("Error initializing database schema: %v", err)
    }
    log.Println("Database schema initialized.")
}