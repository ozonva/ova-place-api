package repo

import (
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/ozonva/ova-place-api/internal/models"
)

type Repo interface {
	TotalCount() (uint64, error)
	AddEntity(entity models.Place) (uint64, error)
	AddEntities(entities []models.Place) error
	ListEntities(limit, offset uint64) ([]models.Place, error)
	DescribeEntity(entityID uint64) (*models.Place, error)
	UpdateEntity(entityID uint64, entity models.Place) error
	RemoveEntity(entityID uint64) error
}

type repo struct {
	db sqlx.DB
}

func NewRepo(db *sqlx.DB) Repo {
	return &repo{db: *db}
}

func (r *repo) TotalCount() (uint64, error) {
	var count uint64
	err := r.db.Get(&count, "SELECT count(1) FROM places")
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *repo) AddEntity(entity models.Place) (uint64, error) {
	var id uint64
	query, err := r.db.PrepareNamed(`INSERT INTO places (user_id,memo,seat) VALUES (:user_id,:memo,:seat) RETURNING id`)

	if err != nil {
		return 0, err
	}

	err = query.Get(&id, entity)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) AddEntities(entities []models.Place) error {
	_, err := r.db.NamedExec(`INSERT INTO person (user_id, memo, seat)
        VALUES (:user_id, :memo, :seat)`, entities)

	return err
}

func (r *repo) ListEntities(limit, offset uint64) ([]models.Place, error) {
	places := make([]models.Place, 0, limit)
	err := r.db.Select(&places, "SELECT id, user_id, memo, seat FROM places ORDER BY id ASC LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return nil, err
	}

	return places, nil
}

func (r *repo) DescribeEntity(entityID uint64) (*models.Place, error) {
	place := models.Place{}
	err := r.db.Get(&place, "SELECT user_id, memo, seat FROM places WHERE id=$1", entityID)
	if err != nil {
		return nil, mapErrors(err)
	}

	return &place, nil
}

func (r *repo) UpdateEntity(entityID uint64, entity models.Place) error {
	res, err := r.db.NamedExec(`UPDATE places SET user_id=:user_id, memo=:memo, seat=:seat, updated_at=:updated_at where id=:id`,
		map[string]interface{}{
			"user_id":    entity.UserID,
			"memo":       entity.Memo,
			"seat":       entity.Seat,
			"id":         entityID,
			"updated_at": time.Now(),
		})

	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return &NotFound{}
	}

	return nil
}

func (r *repo) RemoveEntity(entityID uint64) error {
	res, err := r.db.NamedExec(`DELETE from places where id = :id`, map[string]interface{}{
		"id": entityID,
	})

	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return &NotFound{}
	}

	return nil
}

func mapErrors(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return &NotFound{}
	}

	return err
}
