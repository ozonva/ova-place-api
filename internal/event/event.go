package event

import (
	"encoding/json"

	"github.com/ozonva/ova-place-api/internal/models"
)

// Event keeps its own type and place model
type Event struct {
	EventType string       `json:"event_type"`
	Place     models.Place `json:"place"`
}

// NewEvent returns Event in bytes
func NewEvent(eventType string, model models.Place) ([]byte, error) {
	event := Event{
		EventType: eventType,
		Place:     model,
	}

	modelInBytes, err := json.Marshal(event)

	if err != nil {
		return nil, err
	}

	return modelInBytes, nil
}
