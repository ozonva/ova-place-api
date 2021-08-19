package repo

import "github.com/ozonva/ova-place-api/internal/models"

type Repo interface {
	AddEntities(entities []models.Place) error
	ListEntities(limit, offset uint64) ([]models.Place, error)
	DescribeEntity(entityID uint64) (*models.Place, error)
}
