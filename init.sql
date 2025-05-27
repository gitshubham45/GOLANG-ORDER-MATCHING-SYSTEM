USE order_matching;

-- Orders table
CREATE TABLE IF NOT EXISTS orders (
    id VARCHAR(36) PRIMARY KEY,
    symbol VARCHAR(10) NOT NULL,
    side ENUM('buy', 'sell') NOT NULL,
    type ENUM('limit', 'market') NOT NULL,
    price DECIMAL(15,2),
    initialQuantity DECIMAL(15,8) NOT NULL,
    remainingQuantity DECIMAL(15,8) NOT NULL,
    status ENUM('open', 'filled', 'canceled') NOT NULL,
    createdAt DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME ON UPDATE CURRENT_TIMESTAMP
);

-- Trades table
CREATE TABLE IF NOT EXISTS trades (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    buyOrderId VARCHAR(36) NOT NULL,
    sellOrderId VARCHAR(36) NOT NULL,
    symbol VARCHAR(10) NOT NULL,
    price DECIMAL(15,2) NOT NULL,
    quantity DECIMAL(15,8) NOT NULL,
    matchedAt DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);