package utils

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

type SplitTestCaseInt struct {
	expected     [][]int
	sliceToSplit []int
	batchSize    int
}

type FilterTestCaseInt struct {
	expected      []int
	sliceToFilter []int
	blackList     []int
}

type SplitTestCaseStr struct {
	expected     [][]string
	sliceToSplit []string
	batchSize    int
}

type FilterTestCaseStr struct {
	expected      []string
	sliceToFilter []string
	blackList     []string
}

func TestSplitInt(t *testing.T) {
	testCases := []SplitTestCaseInt{
		{
			[][]int{{1, 3, 2, 1, 5, 4, 7}, {8}}, []int{1, 3, 2, 1, 5, 4, 7, 8}, 7,
		},
		{
			[][]int{{1, 3, 2, 1, 5, 4}, {7, 8}}, []int{1, 3, 2, 1, 5, 4, 7, 8}, 6,
		},
		{
			[][]int{{1, 3, 2, 1, 5}, {4, 7, 8}}, []int{1, 3, 2, 1, 5, 4, 7, 8}, 5,
		},
		{
			[][]int{{1, 3, 2, 1}, {5, 4, 7, 8}}, []int{1, 3, 2, 1, 5, 4, 7, 8}, 4,
		},
		{
			[][]int{{1, 3, 2}, {1, 5, 4}, {7, 8}}, []int{1, 3, 2, 1, 5, 4, 7, 8}, 3,
		},
		{
			[][]int{{1, 3}, {2, 1}, {5, 4}, {7, 8}}, []int{1, 3, 2, 1, 5, 4, 7, 8}, 2,
		},
		{
			[][]int{{1}, {3}, {2}, {1}, {5}, {4}, {7}, {8}}, []int{1, 3, 2, 1, 5, 4, 7, 8}, 1,
		},
	}

	for _, testCase := range testCases {

		splitted := SplitInt(testCase.sliceToSplit, testCase.batchSize)

		if !cmp.Equal(splitted, testCase.expected) {
			t.Fatal("slices do not match", splitted, testCase.expected)
		}
	}
}

func TestSplitStr(t *testing.T) {
	testCases := []SplitTestCaseStr{
		{
			[][]string{{"a", "b", "c", "d", "e", "f", "g"}, {"q"}}, []string{"a", "b", "c", "d", "e", "f", "g", "q"}, 7,
		},
		{
			[][]string{{"a", "b", "c", "d", "e", "f"}, {"g", "q"}}, []string{"a", "b", "c", "d", "e", "f", "g", "q"}, 6,
		},
		{
			[][]string{{"a", "b", "c", "d", "e"}, {"f", "g", "q"}}, []string{"a", "b", "c", "d", "e", "f", "g", "q"}, 5,
		},
		{
			[][]string{{"a", "b", "c", "d"}, {"e", "f", "g", "q"}}, []string{"a", "b", "c", "d", "e", "f", "g", "q"}, 4,
		},
		{
			[][]string{{"a", "b", "c"}, {"d", "e", "f"}, {"g", "q"}}, []string{"a", "b", "c", "d", "e", "f", "g", "q"}, 3,
		},
		{
			[][]string{{"a", "b"}, {"c", "d"}, {"e", "f"}, {"g", "q"}}, []string{"a", "b", "c", "d", "e", "f", "g", "q"}, 2,
		},
		{
			[][]string{{"a"}, {"b"}, {"c"}, {"d"}, {"e"}, {"f"}, {"g"}, {"q"}}, []string{"a", "b", "c", "d", "e", "f", "g", "q"}, 1,
		},
	}

	for _, testCase := range testCases {

		splitted := SplitStr(testCase.sliceToSplit, testCase.batchSize)

		if !cmp.Equal(splitted, testCase.expected) {
			t.Fatal("slices do not match", splitted, testCase.expected)
		}
	}
}

func TestFilterByBlackListInt(t *testing.T) {
	testCases := []FilterTestCaseInt{
		{
			[]int{1, 3, 2, 1, 5}, []int{1, 3, 2, 1, 5, 4, 7, 8}, []int{4, 7, 8},
		},
	}

	for _, testCase := range testCases {

		filtered := FilterByBlackListInt(testCase.sliceToFilter, testCase.blackList)

		if !cmp.Equal(filtered, testCase.expected) {
			t.Fatal("slices do not match", filtered, testCase.expected)
		}
	}
}

func TestFilterByBlackListStr(t *testing.T) {
	testCases := []FilterTestCaseStr{
		{
			[]string{"a", "b", "c", "d", "e"}, []string{"a", "b", "c", "d", "e", "f", "g", "q"}, []string{"f", "g", "q"},
		},
	}

	for _, testCase := range testCases {

		filtered := FilterByBlackListStr(testCase.sliceToFilter, testCase.blackList)

		if !cmp.Equal(filtered, testCase.expected) {
			t.Fatal("slices do not match", filtered, testCase.expected)
		}
	}
}
