# GOLANG ORDER MATCHING SYSTEM

A simple **Order Matching Engine** written in Go, simulating a stock exchange matching engine. Supports both **limit and market orders**, enforces **price-time priority**, and persists data in **MySQL**.

This system allows clients to place buy/sell orders via a REST API, matches incoming orders with existing ones based on price-time priority, and stores order and trade data in a MySQL database.

---

## 📦 Features

- ✅ RESTful API to place, cancel, and query orders  
- ✅ Order types: `limit`, `market`  
- ✅ Matching Engine with price-time priority (FIFO at same price)  
- ✅ Order book maintained in-memory  
- ✅ Trade history recorded on match  
- ✅ Data stored in MySQL (Dockerized for easy setup)  
- ✅ Configurable via `.env` file  
- ✅ REST APIs:
  - `POST /api/orders` – Place order  
  - `DELETE /api/orders/{id}` – Cancel open order  
  - `GET /api/orderbook?symbol={symbol}` – View order book  
  - `GET /api/trades?symbol={symbol}` – View trade history  
  - `GET /api/orders/{id}` – Get order status  

---


---

## 🔧 Dependencies

### Local Machine Setup

1. **Go**: v1.20 or newer  
2. **Docker** and **Docker Compose**  
3. **API-Testing**: `curl` or `postman` for testing APIs  

---

## 🚀 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/GOLANG-ORDER-MATCHING-SYSTEM.git 
cd GOLANG-ORDER-MATCHING-SYSTEM
```

### 2. Set Up Environment Variables
- Create a .env file:

```bash
    cp .env.example .env
```

### 3. Start MySQL and Go App with Docker Compose

```bash
    docker-compose up -d --build
```


🧪 Test the API

### 1. Place an Order( Buy order with limit)


```bash 
curl -X POST http://localhost:8080/api/orders \
     -H "Content-Type: application/json" \
     -d '{
           "symbol": "AAPL",
           "side": "buy",
           "type": "limit",
           "price": 100.0,
           "quantity": 5
         }'
```

### 2. sell order with limit

```bash 
curl -X POST http://localhost:8080/api/orders \
     -H "Content-Type: application/json" \
     -d '{
           "symbol": "AAPL",
           "side": "sell",
           "type": "limit",
           "price": 100.0,
           "quantity": 5
         }'
```

### 3. Normal buy/sell order at market price

```bash 
curl -X POST http://localhost:8080/api/orders \
     -H "Content-Type: application/json" \
     -d '{
           "symbol": "AAPL",
           "side": "sell", // buy / sell
           "price": 100.0,
           "quantity": 5
         }'
```

### 4. View Order Book 

```bash 
    curl "http://localhost:8080/api/orderbook?symbol=AAPL"
```