package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
}

func NewConsumer(brokers []string, topic string, groupID string) *Consumer {
	return &Consumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:  brokers,
			Topic:    topic,
			GroupID:  groupID,
			MinBytes: 10e3,
			MaxBytes: 10e6,
		}),
	}
}

func (p *Consumer) Consume(ctx context.Context, handler func([]byte) error) error {
	for {
		msg, err := p.reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("Error reading message: %v", err)
			continue
		}

		if err := handler(msg.Value); err != nil {
			log.Printf("Error handling message: %v", err)
		}
	}
}

func (p *Consumer) Close() error {
	return p.reader.Close()
}
