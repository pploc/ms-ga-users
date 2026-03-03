package messaging

import (
	"encoding/json"
	"fmt"

	"ms-ga-user/pkg/config"
	"ms-ga-user/pkg/utils"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
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
	producer *kafka.Producer
	topic    string
}

func NewKafkaProducer(cfg *config.Config) (KafkaProducer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.KafkaBrokers,
	})
	if err != nil {
		return nil, err
	}

	utils.Log.Info("Connected to Kafka successfully", zap.String("brokers", cfg.KafkaBrokers))

	// Delivery confirmation go routine
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					utils.Log.Error("Failed to deliver message", zap.Error(ev.TopicPartition.Error))
				} else {
					utils.Log.Debug("Successfully produced record to topic",
						zap.String("topic", *ev.TopicPartition.Topic),
						zap.Int32("partition", ev.TopicPartition.Partition),
						zap.Any("offset", ev.TopicPartition.Offset))
				}
			}
		}
	}()

	return &kafkaProducerImpl{
		producer: p,
		topic:    cfg.KafkaTopicUser,
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

	topic := k.topic
	err = k.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          payload,
	}, nil)

	if err != nil {
		return fmt.Errorf("failed to enqueue message: %w", err)
	}

	return nil
}

func (k *kafkaProducerImpl) Close() {
	if k.producer != nil {
		// Wait for message deliveries before shutting down
		k.producer.Flush(15 * 1000)
		k.producer.Close()
	}
}
