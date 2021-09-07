package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/ozonva/ova-place-api/internal/models"
)

// Repo is an interface for interacting with db through models.Place.
type Repo interface {
	TotalCount(ctx context.Context) (uint64, error)
	AddEntity(ctx context.Context, entity models.Place) (uint64, error)
	AddEntities(ctx context.Context, entities []models.Place) error
	ListEntities(ctx context.Context, limit, offset uint64) ([]models.Place, error)
	DescribeEntity(ctx context.Context, entityID uint64) (*models.Place, error)
	UpdateEntity(ctx context.Context, entityID uint64, entity models.Place) error
	RemoveEntity(ctx context.Context, entityID uint64) error
}

// repo is a Repo implementation.
type repo struct {
	db sqlx.DB
}

// NewRepo returns Repo.
func NewRepo(db *sqlx.DB) Repo {
	return &repo{db: *db}
}

// TotalCount returns total rows count from places table.
func (r *repo) TotalCount(ctx context.Context) (uint64, error) {
	var count uint64
	err := r.db.Get(&count, "SELECT count(1) FROM places")
	if err != nil {
		return 0, fmt.Errorf("cannot Get: %w", err)
	}

	return count, nil
}

// AddEntity inserts place in the table.
func (r *repo) AddEntity(ctx context.Context, entity models.Place) (uint64, error) {
	var id uint64
	query, err := r.db.PrepareNamed(`INSERT INTO places (user_id,memo,seat) VALUES (:user_id,:memo,:seat) RETURNING id`)

	if err != nil {
		return 0, fmt.Errorf("cannot PrepareNamed: %w", err)
	}

	err = query.Get(&id, entity)

	if err != nil {
		return 0, fmt.Errorf("cannot Get: %w", err)
	}

	return id, nil
}

// AddEntities inserts places in the table.
func (r *repo) AddEntities(ctx context.Context, entities []models.Place) error {
	_, err := r.db.NamedExec(`INSERT INTO places (user_id, memo, seat)
        VALUES (:user_id, :memo, :seat)`, entities)
	if err != nil {
		return fmt.Errorf("cannot NamedExec: %w", err)
	}

	return nil
}

// ListEntities returns places with a pagination.
func (r *repo) ListEntities(ctx context.Context, limit, offset uint64) ([]models.Place, error) {
	places := make([]models.Place, 0, limit)
	err := r.db.Select(&places, "SELECT id, user_id, memo, seat FROM places ORDER BY id ASC LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return nil, fmt.Errorf("cannot Select: %w", err)
	}

	return places, nil
}

// DescribeEntity returns place.
func (r *repo) DescribeEntity(ctx context.Context, entityID uint64) (*models.Place, error) {
	place := models.Place{}
	err := r.db.Get(&place, "SELECT user_id, memo, seat FROM places WHERE id=$1", entityID)
	if err != nil {
		return nil, mapErrors(fmt.Errorf("cannot Get: %w", err))
	}

	return &place, nil
}

// UpdateEntity updates the place.
func (r *repo) UpdateEntity(ctx context.Context, entityID uint64, entity models.Place) error {
	res, err := r.db.NamedExec(`UPDATE places SET user_id=:user_id, memo=:memo, seat=:seat, updated_at=:updated_at where id=:id`,
		map[string]interface{}{
			"user_id":    entity.UserID,
			"memo":       entity.Memo,
			"seat":       entity.Seat,
			"id":         entityID,
			"updated_at": time.Now(),
		})

	if err != nil {
		return fmt.Errorf("cannot NamedExec: %w", err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("cannot RowsAffected: %w", err)
	}

	if count == 0 {
		return &NotFound{}
	}

	return nil
}

// RemoveEntity deletes the place.
func (r *repo) RemoveEntity(ctx context.Context, entityID uint64) error {
	res, err := r.db.NamedExec(`DELETE from places where id = :id`, map[string]interface{}{
		"id": entityID,
	})

	if err != nil {
		return fmt.Errorf("cannot NamedExec: %w", err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("cannot RowsAffected: %w", err)
	}

	if count == 0 {
		return &NotFound{}
	}

	return nil
}

// mapErrors maps lib errors to internal error types.
func mapErrors(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return &NotFound{}
	}

	return err
}
