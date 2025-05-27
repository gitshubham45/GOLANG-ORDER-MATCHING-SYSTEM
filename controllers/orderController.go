package controllers

import (
	"database/sql"
	"fmt"
	"golangOrderMatchingSystem/db"
	"golangOrderMatchingSystem/models"
	"golangOrderMatchingSystem/services"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getenvDefault(key, defaultValue string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultValue
}

func GetPort() string {
	return getenvDefault("PORT", "8080")
}

func PlaceOrder(c *gin.Context) {
	var req struct {
		Symbol   string  `json:"symbol"`
		Side     string  `json:"side"`
		Type     string  `json:"type"`
		Price    float64 `json:"price,omitempty"`
		Quantity float64 `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if req.Quantity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quantity"})
		return
	}

	rand.Seed(time.Now().UnixNano())
	orderId := generateUUID()

	newOrder := models.Order{
		ID:                orderId,
		Symbol:            req.Symbol,
		Side:              req.Side,
		Type:              req.Type,
		Price:             req.Price,
		InitialQuantity:   req.Quantity,
		RemainingQuantity: req.Quantity,
		Status:            "pending",
	}

	services.MatchIncomingOrder(newOrder)

	db.SaveOrder(newOrder)

	c.JSON(http.StatusOK, gin.H{
		"id":                newOrder.ID,
		"status":            newOrder.Status,
		"remainingQuantity": newOrder.RemainingQuantity,
	})
}

func generateUUID() string {
	return uuid.New().String()
}

func GetTradeHistory(c *gin.Context) {
	symbol := c.Query("symbol")
	if symbol == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing symbol parameter"})
		return
	}

	trades, err := db.GetTradesBySymbol(symbol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch trades"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"trades": trades})
}

func GetOrderStatus(c *gin.Context) {
	orderID := c.Param("orderId")
	if orderID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing order ID"})
		return
	}

	order, err := db.GetOrderById(orderID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}

func GetOrderBook(c *gin.Context) {
	symbol := c.Query("symbol")
	if symbol == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'symbol' parameter"})
		return
	}

	fmt.Printf("Symbol : %s", symbol)

	bids, err := db.GetOrderBookEntries(symbol, "buy")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bids"})
		fmt.Println(err)
		return
	}

	asks, err := db.GetOrderBookEntries(symbol, "sell")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch asks"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"bids": bids,
		"asks": asks,
	})
}
