package kafka

import (
	"context"
	"encoding/json"
	"time"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(brokers []string, topic string) *Producer {
	return &Producer{
		writer: &kafka.Writer{
			Addr:         kafka.TCP(brokers...),
			Topic:        topic,
			Balancer:     &kafka.LeastBytes{},
			RequiredAcks: kafka.RequireAll,
			MaxAttempts:  3,
		},
	}
}

func (p *Producer) SendEvent(ctx context.Context, key string, event interface{}) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	msg := &kafka.Message{
		Key:   []byte(key),
		Value: data,
		Time:  time.Now(),
	}

	return p.writer.WriteMessages(ctx, *msg)
}

func (p *Producer) Close() error {
	return p.writer.Close()
}
