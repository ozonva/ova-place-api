package utils

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

type InvertTestCaseInt struct {
	expected    map[int]int
	mapToInvert map[int]int
}

type InvertTestCaseStr struct {
	expected    map[string]string
	mapToInvert map[string]string
}

func TestInvertMappingInt(t *testing.T) {
	testCases := []InvertTestCaseInt{
		{
			map[int]int{1: 3, 2: 1, 5: 4}, map[int]int{3: 1, 1: 2, 4: 5},
		},
	}

	for _, testCase := range testCases {
		if !cmp.Equal(InvertMappingInt(testCase.mapToInvert), testCase.expected) {
			t.Fatal("maps not match", testCase.mapToInvert, testCase.expected)
		}
	}
}

func TestInvertMappingIntPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	InvertMappingInt(map[int]int{3: 1, 1: 2, 4: 1})
}

func TestInvertMappingStr(t *testing.T) {
	testCases := []InvertTestCaseStr{
		{
			map[string]string{"a": "b", "c": "d", "e": "f"}, map[string]string{"b": "a", "d": "c", "f": "e"},
		},
	}

	for _, testCase := range testCases {
		if !cmp.Equal(InvertMappingStr(testCase.mapToInvert), testCase.expected) {
			t.Fatal("maps not match", testCase.mapToInvert, testCase.expected)
		}
	}
}

func TestInvertMappingStrPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	InvertMappingStr(map[string]string{"b": "a", "d": "c", "f": "a"})
}
