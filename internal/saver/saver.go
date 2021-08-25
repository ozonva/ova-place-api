package saver

import (
	"errors"
	"github.com/ozonva/ova-place-api/internal/utils"
	"time"

	"github.com/ozonva/ova-place-api/internal/flusher"
	"github.com/ozonva/ova-place-api/internal/models"
)

// Saver is an interface for models.Place periodic saving.
type Saver interface {
	Save(entity models.Place) error
	Close() error
	init(tickDuration time.Duration)
}

// NewSaver returns Saver.
func NewSaver(
	capacity uint,
	tickDuration time.Duration,
	flusher flusher.Flusher,
) Saver {
	saver := &saver{
		done:     *utils.NewSyncChannel(),
		entities: make([]models.Place, 0, capacity),
		flusher:  flusher,
	}

	saver.init(tickDuration)

	return saver
}

// saver is a Saver implementation.
type saver struct {
	done     utils.SyncChannel
	entities []models.Place
	flusher  flusher.Flusher
}

// Save adds models.Place to the buffer.
// It returns nil when the models.Place has been successfully added in a buffer.
// It returns an error when the buffer capacity is exceeded.
func (s *saver) Save(entity models.Place) error {
	if len(s.entities) >= cap(s.entities) {
		return errors.New("capacity is exceeded")
	}

	s.entities = append(s.entities, entity)

	return nil
}

// Close flushes remaining entities in buffer and close the buffer.
// It returns nil when all entities have been successfully flushed.
// It returns an error and does not close the buffer when some entities have not been saved.
func (s *saver) Close() error {
	err := s.flush()
	if err != nil {
		return err
	}

	s.done.Once.Do(func() {
		s.done.C <- true
		close(s.done.C)
	})

	return nil
}

// init configures periodic saving.
func (s *saver) init(tickDuration time.Duration) {
	ticker := time.NewTicker(tickDuration)

	go func(s *saver, ticker *time.Ticker) {

		for {
			select {
			case <-s.done.C:
				ticker.Stop()

				return
			case <-ticker.C:
				s.flush() // todo log
			}
		}
	}(s, ticker)

}

// flush performs entities saving.
// It returns nil when all entities have been successfully flushed.
// It returns an error when some entities have not been saved.
// These entities will remain in the buffer for the next saving.
func (s *saver) flush() error {
	if len(s.entities) == 0 {
		return nil
	}

	unsaved := s.flusher.Flush(s.entities)

	if len(unsaved) > 0 {
		copy(s.entities, unsaved)
		return errors.New("there are unsaved entities")
	}

	s.entities = s.entities[:0]

	return nil
}
