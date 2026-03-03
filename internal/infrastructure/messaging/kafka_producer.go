package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"ms-ga-user/pkg/config"
	"ms-ga-user/pkg/utils"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type EventType string

const (
	EventUserCreated     EventType = "user.created"
	EventUserUpdated     EventType = "user.updated"
	EventUserDeactivated EventType = "user.deactivated"
)

type KafkaEvent struct {
	EventType string      `json:"event_type"`
	Source    string      `json:"source"`
	Data      interface{} `json:"data"`
}

type KafkaProducer interface {
	PublishEvent(eventType EventType, data interface{}) error
	Close()
}

type kafkaProducerImpl struct {
	writer *kafka.Writer
	topic  string
}

func NewKafkaProducer(cfg *config.Config) (KafkaProducer, error) {
	w := &kafka.Writer{
		Addr:                   kafka.TCP(cfg.KafkaBrokers),
		Topic:                  cfg.KafkaTopicUser,
		Balancer:               &kafka.LeastBytes{},
		BatchTimeout:           10 * time.Millisecond,
		AllowAutoTopicCreation: true,
	}

	utils.Log.Info("Initialized Kafka Writer (Pure Go)", zap.String("brokers", cfg.KafkaBrokers), zap.String("topic", cfg.KafkaTopicUser))

	return &kafkaProducerImpl{
		writer: w,
		topic:  cfg.KafkaTopicUser,
	}, nil
}

func (k *kafkaProducerImpl) PublishEvent(eventType EventType, data interface{}) error {
	event := KafkaEvent{
		EventType: string(eventType),
		Source:    "ms-ga-user",
		Data:      data,
	}

	payload, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal kafka event: %w", err)
	}

	err = k.writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(string(eventType)),
			Value: payload,
		},
	)

	if err != nil {
		return fmt.Errorf("failed to write message: %w", err)
	}

	utils.Log.Debug("Successfully produced record to topic", zap.String("topic", k.topic), zap.String("eventType", string(eventType)))

	return nil
}

func (k *kafkaProducerImpl) Close() {
	if k.writer != nil {
		if err := k.writer.Close(); err != nil {
			utils.Log.Error("failed to close kafka writer", zap.Error(err))
		}
	}
}
