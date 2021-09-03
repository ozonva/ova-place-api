package producer

import (
	"fmt"

	"github.com/Shopify/sarama"
)

// Producer is an interface for sending data to kafka
type Producer interface {
	Push(topic string, message []byte) error
	Close() error
}

// NewProducer returns Producer
func NewProducer(brokersURL []string) (Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	conn, err := sarama.NewSyncProducer(brokersURL, config)
	if err != nil {
		return nil, fmt.Errorf("cannot init the connection to kafka: %w", err)
	}
	return &producer{
		conn: conn,
	}, nil
}

// producer is a Producer implementation
type producer struct {
	conn sarama.SyncProducer
}

// Push sends byte message to kafka
func (p *producer) Push(topic string, message []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	_, _, err := p.conn.SendMessage(msg)
	if err != nil {
		return fmt.Errorf("cannot SendMessage: %w", err)
	}
	return nil
}

// Close closes the connection with kafka
func (p *producer) Close() error {
	err := p.conn.Close()
	if err != nil {
		return fmt.Errorf("cannot Close connection with kafka: %w", err)
	}

	return nil
}
