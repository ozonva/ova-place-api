package utils

import (
	"errors"

	"github.com/ozonva/ova-place-api/internal/models"
)

// SplitPlacesToBatches splits slice to slice of slices and returns result slice.
func SplitPlacesToBatches(entities []models.Place, batchSize int) ([][]models.Place, error) {

	if batchSize == 0 {
		return nil, errors.New("the batchSize is zero")
	}

	sliceLen := len(entities)
	splittedLen := sliceLen / batchSize
	if sliceLen%batchSize != 0 {
		splittedLen++
	}
	splitted := make([][]models.Place, splittedLen)

	for i := 0; i < splittedLen; i++ {

		if i == 0 {
			splitted[i] = entities[i : i+batchSize]
			continue
		}

		if i*batchSize+batchSize > sliceLen {
			splitted[i] = entities[i*batchSize:]
			continue
		}

		splitted[i] = entities[i*batchSize : i*batchSize+batchSize]
	}

	return splitted, nil
}

// PlacesSliceToMap converts slice to map (key is a user_id) and returns result map.
func PlacesSliceToMap(entities []models.Place) (map[uint64]models.Place, error) {
	resultMap := make(map[uint64]models.Place, len(entities))

	for i := range entities {
		if _, ok := resultMap[entities[i].UserID]; ok {
			return nil, errors.New("duplicate keys")
		}
		resultMap[entities[i].UserID] = entities[i]
	}

	return resultMap, nil
}
