package models

import "time"

type OrderStatus string
type PaymentStatus string
type InventoryStatus string

const (
	OrderPending   OrderStatus = "PENDING"
	OrderConfirmed OrderStatus = "CONFIRMED"
	OrderFailed    OrderStatus = "FAILED"

	PaymentPending PaymentStatus = "PENDING"
	PaymentSuccess PaymentStatus = "SUCCESS"
	PaymentFailed  PaymentStatus = "FAILED"

	InventoryReserved InventoryStatus = "RESERVED"
	InventoryReleased InventoryStatus = "RELEASED"
)

type Order struct {
	ID        string      `json:"id"`
	UserID    string      `json:"user_id"`
	ProductID string      `json:"product_id"`
	Quantity  int         `json:"quantity"`
	Amount    int         `json:"amount"`
	Status    OrderStatus `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
}
type SagaEvent struct {
	EventID   string                 `json:"event_id"`
	EventType string                 `json:"event_type"`
	OrderID   string                 `json:"order_id"`
	Success   bool                   `json:"success"`
	Data      map[string]interface{} `json:"data"`
	Timestamp time.Time              `json:"timestamp"`
}
