package main

import (
	"context"
	"encoding/json"
	"fmt"
	"learning-gin/distributed-transation-mvp/internal/kafka"
	"learning-gin/distributed-transation-mvp/internal/models"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
)

var (
	inventory   = map[string]int{"product-1": 100, "product-2": 0, "product-3": 50}
	inventoryMu sync.RWMutex
	producer    *kafka.Producer
)

func main() {
	producer = kafka.NewProducer([]string{"localhost:9092"}, "saga-events")
	defer producer.Close()

	consumer := kafka.NewConsumer([]string{"localhost:9092"}, "saga-events", "inventory-service")
	defer consumer.Close()

	log.Println("Inventory service started...")

	handler := func(data []byte) error {
		var event models.SagaEvent
		if err := json.Unmarshal(data, &event); err != nil {
			return err
		}

		if event.EventType == "ORDER_CREATED" {
			handleOrderCreated(event)
		} else if event.EventType == "Payment_Failed" {
			handlePaymentFailed(event)
		}
		return nil
	}

	consumer.Consume(context.Background(), handler)
}

func handleOrderCreated(event models.SagaEvent) {
	productID := event.Data["product_id"].(string)
	quantity := int(event.Data["quantity"].(float64))

	inventoryMu.Lock()
	defer inventoryMu.Unlock()

	currentStock, exist := inventory[productID]
	success := exist && currentStock >= quantity
	if success {
		inventory[productID] -= quantity
		log.Printf("Reserved %d unit of %s. Remaining is %d", quantity, productID, inventory[productID])
	} else {
		log.Printf("Infsufficient stock for %s. Requested: %d, Available: %d", productID, quantity, currentStock)
	}
	fmt.Printf("event %v\n", event)
	responseEvent := models.SagaEvent{
		EventID:   uuid.New().String(),
		EventType: map[bool]string{true: "INVENTORY_RESERVED", false: "INVENTORY_FAILED"}[success],
		OrderID:   event.OrderID,
		Success:   success,
		Data:      event.Data,
		Timestamp: time.Now(),
	}

	producer.SendEvent(context.Background(), event.OrderID, responseEvent)
}

func handlePaymentFailed(event models.SagaEvent) {
	productID := event.Data["product_id"].(string)
	quantity := int(event.Data["quantity"].(float64))
	inventoryMu.Lock()
	inventory[productID] += quantity
	inventoryMu.Unlock()

	log.Printf("Compensated: Released %d unit of %s\n", quantity, productID)
}
