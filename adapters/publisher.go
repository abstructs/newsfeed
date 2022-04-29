package adapters

import (
	"context"

	"github.com/abstructs/newsfeed/models"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

const (
	topic = "message-log"
)

type Publisher struct {
	logger *zap.Logger
	Writer *kafka.Writer
}

func NewWriter(logger *zap.Logger, brokers []string) models.IPublisher {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: brokers,
		Topic:   topic,
	})

	return &Publisher{
		Writer: w,
		logger: logger,
	}
}

func (p *Publisher) Publish(ctx context.Context, key string, value string) error {
	err := p.Writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(key),
		Value: []byte(value),
	})
	if err != nil {
		p.logger.Sugar().Errorf("Failed to publish message:", err)
		return err
	}

	return nil
}
