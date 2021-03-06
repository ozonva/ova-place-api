package flusher

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"

	"github.com/ozonva/ova-place-api/internal/models"
	"github.com/ozonva/ova-place-api/internal/repo"
	"github.com/ozonva/ova-place-api/internal/utils"
)

// Flusher is an interface for dumping places to storage.
type Flusher interface {
	Flush(ctx context.Context, places []models.Place) []models.Place
}

// NewFlusher returns Flusher with batch saving support.
func NewFlusher(
	batchSize int,
	entityRepo repo.Repo,
) Flusher {
	return &flusher{
		batchSize:  batchSize,
		entityRepo: entityRepo,
	}
}

// flusher is a Flusher implementation.
type flusher struct {
	batchSize  int
	entityRepo repo.Repo
}

// Flush saves places in batches.
// It returns nil when all places have been successfully saved.
// It can return places, which have been not saved.
func (f *flusher) Flush(ctx context.Context, places []models.Place) []models.Place {

	batches, err := utils.SplitPlacesToBatches(places, f.batchSize)

	if err != nil {
		return places
	}

	notAdded := make([]models.Place, 0, len(places))

	for index := range batches {

		span, _ := opentracing.StartSpanFromContext(ctx, "batch_save")
		span.LogFields(log.Int("places_count", len(batches[index])))

		err := f.entityRepo.AddEntities(ctx, batches[index])

		if err != nil {
			notAdded = append(notAdded, batches[index]...)
		}

		span.Finish()
	}

	if len(notAdded) == 0 {
		return nil
	}

	return notAdded
}
