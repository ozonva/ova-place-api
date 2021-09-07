package health

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ozonva/ova-place-api/internal/producer"
)

// NewKafkaCheck returns Service.
func NewKafkaCheck(producer *producer.Producer) Service {
	return &kafka{
		producer: *producer,
	}
}

// kafka is a Service implementation.
type kafka struct {
	producer producer.Producer
}

// Health checks kafka status and returns an error if the kafka has some problems.
func (k *kafka) Health() error {
	result, err := json.Marshal("check")
	if err != nil {
		return fmt.Errorf("cannot Marshal: %w", err)
	}
	err = k.producer.Push(context.TODO(), "health_check", result)
	if err != nil {
		return fmt.Errorf("cannot Push: %w", err)
	}

	return nil
}
