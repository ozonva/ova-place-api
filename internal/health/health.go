package health

import (
	"fmt"
)

// Status is a structure for keeping app status.
type Status struct {
	App      string `json:"app"`
	Database string `json:"database"`
	Kafka    string `json:"kafka"`
}

// Service is an interface for check internal service status.
type Service interface {
	Health() error
}

// Health is an interface for check app status.
type Health interface {
	Check() Status
}

// health is a Health implementation.
type health struct {
	database Service
	kafka    Service
}

// NewHealth returns Health instance.
func NewHealth(database Service, kafka Service) Health {
	return &health{
		database: database,
		kafka:    kafka,
	}
}

// Check checks all internal services and returns Status.
func (s *health) Check() Status {
	appStatus := "OK"
	databaseStatus := "OK"
	kafkaStatus := "OK"

	if err := s.database.Health(); err != nil {
		databaseStatus = fmt.Sprintf("Database ERROR: %s", err)
		appStatus = "NOT OK"

	}

	if err := s.kafka.Health(); err != nil {
		kafkaStatus = fmt.Sprintf("Kafka ERROR: %s", err)
		appStatus = "NOT OK"
	}

	return Status{
		App:      appStatus,
		Database: databaseStatus,
		Kafka:    kafkaStatus,
	}
}
