package main

import (
	"context"
	"learning-gin/distributed-transation-mvp/internal/kafka"
	"learning-gin/distributed-transation-mvp/internal/models"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	orders   = make(map[string]*models.Order)
	orderMu  sync.RWMutex
	producer *kafka.Producer
)

func main() {
	producer = kafka.NewProducer([]string{"localhost:9092"}, "saga-events")
	defer producer.Close()

	// 启动消费者监听补偿事件
	go startCompensationListener()

	r := gin.Default()
	r.POST("/orders", createOrder)
	r.GET("/orders/:id", getOrder)

	log.Println("Listening on 8081")
	r.Run(":8081")
}

func startCompensationListener() {

}

func getOrder(c *gin.Context) {
	orderID := c.Param("id")
	orderMu.RLock()
	order, ok := orders[orderID]
	orderMu.RUnlock()
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"message": "order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

func createOrder(c *gin.Context) {
	var req struct {
		UserID    string  `json:"user_id"`
		ProductID string  `json:"product_id"`
		Quantity  int     `json:"quantity"`
		Amount    float64 `json:"amount"`
	}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := &models.Order{
		ID:        uuid.New().String(),
		UserID:    req.UserID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
		Amount:    req.Amount,
		Status:    models.OrderPending,
		CreatedAt: time.Now(),
	}

	orderMu.Lock()
	orders[order.ID] = order
	orderMu.Unlock()

	event := models.SagaEvent{
		EventID:   uuid.New().String(),
		EventType: "ORDER_CREATED",
		OrderID:   order.ID,
		Success:   true,
		Data: map[string]interface{}{
			"user_id":    order.UserID,
			"product_id": order.ProductID,
			"quantity":   order.Quantity,
			"amount":     order.Amount,
		},
		Timestamp: time.Now(),
	}

	if err := producer.SendEvent(context.Background(), order.ID, event); err != nil {
		log.Printf("Failed to send event: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order created"})
}
