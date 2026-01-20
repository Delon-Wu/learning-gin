package main

import (
	"context"
	"encoding/json"
	"learning-gin/distributed-transation-mvp/internal/kafka"
	"learning-gin/distributed-transation-mvp/internal/models"
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

var producer *kafka.Producer

func main() {
	producer = kafka.NewProducer([]string{"localhost:9092"}, "saga-events")
	defer producer.Close()

	consumer := kafka.NewConsumer([]string{"localhost:9092"}, "saga-events", "payment-service")
	defer consumer.Close()

	log.Printf("Payment Service Started")

	handler := func(data []byte) error {
		var event models.SagaEvent
		if err := json.Unmarshal(data, &event); err != nil {
			return err
		}

		if event.EventType == "INVENTORY_RESERVED" {
			handleInventoryReserved(event)
		}
		return nil
	}

	consumer.Consume(context.Background(), handler)
}

func handleInventoryReserved(event models.SagaEvent) {
	var amount = event.Data["amount"].(float64)

	time.Sleep(time.Millisecond * 500)
	success := rand.Float32() > 0.5

	if success {
		log.Printf("Payment Success for order %s. Amount: %.2f", event.OrderID, amount)
	} else {
		log.Printf("Payment Failure for order %s. Amount: %.2f", event.OrderID, amount)
	}

	responseEvent := models.SagaEvent{
		EventID:   uuid.New().String(),
		EventType: map[bool]string{true: "PAYMENT_SUCCESS", false: "PAYMENT_FAILED"}[success],
		OrderID:   event.OrderID,
		Success:   success,
		Data:      event.Data,
		Timestamp: time.Now(),
	}

	producer.SendEvent(context.Background(), event.EventID, responseEvent)
}
