package utils

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ozonva/ova-place-api/internal/models"
)

type SplitPlacesToBatchesTestCase struct {
	expected     [][]models.Place
	sliceToSplit []models.Place
	batchSize    int
}

type PlacesSliceToMapTestCase struct {
	expected map[uint64]models.Place
	slice    []models.Place
}

func TestSplitToBatches(t *testing.T) {
	testCases := []SplitPlacesToBatchesTestCase{
		{
			[][]models.Place{{{UserId: 1}, {UserId: 2}}, {{UserId: 3}, {UserId: 4}}}, []models.Place{{UserId: 1}, {UserId: 2}, {UserId: 3}, {UserId: 4}}, 2,
		},
		{
			[][]models.Place{{{UserId: 1}, {UserId: 2}, {UserId: 3}}, {{UserId: 4}}}, []models.Place{{UserId: 1}, {UserId: 2}, {UserId: 3}, {UserId: 4}}, 3,
		},
	}

	for _, testCase := range testCases {

		splitted, err := SplitPlacesToBatches(testCase.sliceToSplit, testCase.batchSize)

		if err != nil {
			t.Fatal("An error has been occurred", err)
		}

		if !cmp.Equal(splitted, testCase.expected) {
			t.Fatal("The slices do not match", splitted, testCase.expected)
		}
	}
}

func TestSplitToBatchesError(t *testing.T) {
	_, err := SplitPlacesToBatches([]models.Place{{UserId: 1}, {UserId: 2}, {UserId: 3}, {UserId: 4}}, 0)

	if err == nil {
		t.Fatal("An error has not been occurred", err)
	}
}

func TestSliceToMap(t *testing.T) {
	place1 := models.Place{UserId: 1}
	place2 := models.Place{UserId: 2}
	testCases := []PlacesSliceToMapTestCase{
		{
			map[uint64]models.Place{1: place1, 2: place2}, []models.Place{place1, place2},
		},
	}

	for _, testCase := range testCases {

		converted, err := PlacesSliceToMap(testCase.slice)

		if err != nil {
			t.Fatal("An error has been occurred", err)
		}

		if !cmp.Equal(converted, testCase.expected) {
			t.Fatal("The maps do not match", converted, testCase.expected)
		}
	}
}

func TestSliceToMapError(t *testing.T) {
	place1 := models.Place{UserId: 1}
	place2 := models.Place{UserId: 1}

	_, err := PlacesSliceToMap([]models.Place{place1, place2})

	if err == nil {
		t.Fatal("An error has not been occurred", err)
	}
}
