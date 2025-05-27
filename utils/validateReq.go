package utils

import (
    "fmt"
    "strings"
)

func ValidatePlaceOrderReq(symbol, side, orderType string, price, quantity float64) error {
    if symbol == "" {
        return fmt.Errorf("missing 'symbol' field")
    }

    if side == "" {
        return fmt.Errorf("missing 'side' field")
    }

    if !strings.EqualFold(side, "buy") && !strings.EqualFold(side, "sell") {
        return fmt.Errorf("side must be 'buy' or 'sell'")
    }

    if orderType == "" {
        return fmt.Errorf("missing 'type' field")
    }

    if !strings.EqualFold(orderType, "limit") && !strings.EqualFold(orderType, "market") {
        return fmt.Errorf("type must be 'limit' or 'market'")
    }

    if quantity <= 0 {
        return fmt.Errorf("quantity must be greater than zero")
    }

    if strings.EqualFold(orderType, "limit") {
        if price <= 0 {
            return fmt.Errorf("limit order requires a positive price")
        }
    } else if strings.EqualFold(orderType, "market") {
        if price != 0 {
            return fmt.Errorf("market order must not include a price")
        }
    }

    if len(symbol) > 10 {
        return fmt.Errorf("symbol too long (max 10 characters)")
    }

    return nil
}