package main

import (
	"learning-gin/distributed-transation-mvp/internal/kafka"
	"sync"
)

var (
	inventory   = map[string]int{"product-1": 100, "product-2": 0, "product-3": 50}
	inventoryMu sync.RWMutex
	producer    *kafka.Producer
)

func main() {

}
