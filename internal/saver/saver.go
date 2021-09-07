package saver

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/ozonva/ova-place-api/internal/flusher"
	"github.com/ozonva/ova-place-api/internal/models"
	"github.com/ozonva/ova-place-api/internal/utils"
)

// Saver is an interface for models.Place periodic saving.
type Saver interface {
	Save(entity models.Place) error
	Close() error
	init(tickDuration time.Duration)
}

// NewSaver returns Saver.
func NewSaver(
	ctx context.Context,
	capacity uint,
	tickDuration time.Duration,
	flusher flusher.Flusher,
) Saver {
	saver := &saver{
		done:     *utils.NewSyncChannel(),
		entities: make([]models.Place, 0, capacity),
		flusher:  flusher,
		ctx:      ctx,
	}

	saver.init(tickDuration)

	return saver
}

// saver is a Saver implementation.
type saver struct {
	m        sync.Mutex
	entities []models.Place

	done    utils.SyncChannel
	flusher flusher.Flusher
	ctx     context.Context
}

// Save adds models.Place to the buffer.
// It returns nil when the models.Place has been successfully added in a buffer.
// It returns an error when the buffer capacity is exceeded.
func (s *saver) Save(entity models.Place) error {
	s.m.Lock()
	defer s.m.Unlock()

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
				err := s.flush()
				if err != nil {
					log.Err(fmt.Errorf("cannot flush: %w", err)).Msg("Error from saver")
				}
			}
		}
	}(s, ticker)

}

// flush performs entities saving.
// It returns nil when all entities have been successfully flushed.
// It returns an error when some entities have not been saved.
// These entities will remain in the buffer for the next saving.
func (s *saver) flush() error {
	s.m.Lock()
	defer s.m.Unlock()

	if len(s.entities) == 0 {
		return nil
	}

	unsaved := s.flusher.Flush(s.ctx, s.entities)

	if len(unsaved) > 0 {
		copy(s.entities, unsaved)
		return errors.New("there are unsaved entities")
	}

	s.entities = s.entities[:0]

	return nil
}
