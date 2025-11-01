package sets

import (
	"fmt"
	"maps"
	"slices"
	"testing"
)

func TestOrderedUnion(t *testing.T) {
	for _, test := range []struct {
		input    [][]int
		expected []int
	}{
		{
			input:    [][]int{},
			expected: nil,
		},
		{
			input: [][]int{
				{1, 2, 3},
			},
			expected: []int{1, 2, 3},
		},
		{
			input: [][]int{
				{1, 2, 3},
				nil,
			},
			expected: []int{1, 2, 3},
		},
		{
			input: [][]int{
				{1, 1, 2, 2, 3, 3},
			},
			expected: []int{1, 2, 3},
		},
		{
			input: [][]int{
				{1, 2, 3},
				{2, 3, 4},
				{3, 4, 5},
			},
			expected: []int{1, 2, 3, 4, 5},
		},
	} {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			actual := OrderedUnion(test.input...)
			if !slices.Equal(test.expected, actual) {
				t.Errorf("expected %v got %v", test.expected, actual)
			}
		})
	}
}

func TestUnorderedUnion(t *testing.T) {
	for _, test := range []struct {
		input    []map[int]int
		expected map[int]int
	}{
		{
			input:    []map[int]int{},
			expected: nil,
		},
		{
			input: []map[int]int{
				{1: 0, 2: 0, 3: 0},
			},
			expected: map[int]int{1: 0, 2: 0, 3: 0},
		},
		{
			input: []map[int]int{
				{1: 0, 2: 0, 3: 0},
				nil,
			},
			expected: map[int]int{1: 0, 2: 0, 3: 0},
		},
		{
			input: []map[int]int{
				{1: 1, 2: 1, 3: 1},
				{2: 2, 3: 2, 4: 2},
				{3: 3, 4: 3, 5: 3},
			},
			expected: map[int]int{1: 1, 2: 1, 3: 1, 4: 2, 5: 3},
		},
	} {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			actual := UnorderedUnion(test.input...)
			if !maps.Equal(test.expected, actual) {
				t.Errorf("expected %v got %v", test.expected, actual)
			}
		})
	}
}

func TestOrderedComplement(t *testing.T) {
	for _, test := range []struct {
		input    [][]int
		expected []int
	}{
		{
			input:    [][]int{},
			expected: nil,
		},
		{
			input: [][]int{
				{1, 2, 3},
			},
			expected: []int{1, 2, 3},
		},
		{
			input: [][]int{
				{1, 2, 3},
				nil,
			},
			expected: []int{1, 2, 3},
		},
		{
			input: [][]int{
				{1, 2, 3},
				{2, 3, 4},
				{3, 4, 5},
			},
			expected: []int{1},
		},
		{
			input: [][]int{
				{1, 2, 3},
				{1, 2, 3},
			},
			expected: nil,
		},
		{
			input: [][]int{
				{1, 1, 2, 2, 3, 3},
				{2, 2, 3, 3, 4, 4},
				{3, 3, 4, 4, 5, 5},
			},
			expected: []int{1},
		},
		{
			input: [][]int{
				{1, 2, 3},
				{1, 2, 3},
				{3, 4, 5},
			},
			expected: nil,
		},
	} {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			actual := OrderedComplement(test.input...)
			if !slices.Equal(test.expected, actual) {
				t.Errorf("expected %v got %v", test.expected, actual)
			}
		})
	}
}

func TestUnorderedComplement(t *testing.T) {
	for _, test := range []struct {
		input    []map[int]int
		expected map[int]int
	}{
		{
			input:    []map[int]int{},
			expected: nil,
		},
		{
			input: []map[int]int{
				{1: 0, 2: 0, 3: 0},
			},
			expected: map[int]int{1: 0, 2: 0, 3: 0},
		},
		{
			input: []map[int]int{
				{1: 0, 2: 0, 3: 0},
				nil,
			},
			expected: map[int]int{1: 0, 2: 0, 3: 0},
		},
		{
			input: []map[int]int{
				{1: 1, 2: 1, 3: 1},
				{2: 2, 3: 2, 4: 2},
				{3: 3, 4: 3, 5: 3},
			},
			expected: map[int]int{1: 1},
		},
		{
			input: []map[int]int{
				{1: 1, 2: 1, 3: 1},
				{1: 1, 2: 1, 3: 1},
			},
			expected: nil,
		},
		{
			input: []map[int]int{
				{1: 1, 2: 1, 3: 1},
				{1: 1, 2: 1, 3: 1},
				{3: 3, 4: 3, 5: 3},
			},
			expected: nil,
		},
	} {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			actual := UnorderedComplement(test.input...)
			if !maps.Equal(test.expected, actual) {
				t.Errorf("expected %v got %v", test.expected, actual)
			}
		})
	}
}

func TestOrderedIntersect(t *testing.T) {
	for _, test := range []struct {
		input    [][]int
		expected []int
	}{
		{
			input:    [][]int{},
			expected: nil,
		},
		{
			input: [][]int{
				{1, 2, 3},
			},
			expected: []int{1, 2, 3},
		},
		{
			input: [][]int{
				{1, 2, 3},
				nil,
			},
			expected: nil,
		},
		{
			input: [][]int{
				{1, 2, 3},
				{2, 3, 4},
				{3, 4, 5},
			},
			expected: []int{3},
		},
		{
			input: [][]int{
				{1, 2, 3},
				{1, 2, 3},
			},
			expected: []int{1, 2, 3},
		},
		{
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			expected: nil,
		},
		{
			input: [][]int{
				{1, 1, 2, 2, 3, 3},
				{2, 2, 3, 3, 4, 4},
				{3, 3, 4, 4, 5, 5},
			},
			expected: []int{3},
		},
		{
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{3, 4, 5},
			},
			expected: nil,
		},
	} {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			actual := OrderedIntersect(test.input...)
			if !slices.Equal(test.expected, actual) {
				t.Errorf("expected %v got %v", test.expected, actual)
			}
		})
	}
}

func TestUnorderedIntersect(t *testing.T) {
	for _, test := range []struct {
		input    []map[int]int
		expected map[int]int
	}{
		{
			input:    []map[int]int{},
			expected: nil,
		},
		{
			input: []map[int]int{
				{1: 0, 2: 0, 3: 0},
			},
			expected: map[int]int{1: 0, 2: 0, 3: 0},
		},
		{
			input: []map[int]int{
				{1: 0, 2: 0, 3: 0},
				nil,
			},
			expected: nil,
		},
		{
			input: []map[int]int{
				{1: 1, 2: 1, 3: 1},
				{2: 2, 3: 2, 4: 2},
				{3: 3, 4: 3, 5: 3},
			},
			expected: map[int]int{3: 1},
		},
		{
			input: []map[int]int{
				{1: 1, 2: 1, 3: 1},
				{4: 2, 5: 2, 6: 2},
			},
			expected: nil,
		},
		{
			input: []map[int]int{
				{1: 1, 2: 1, 3: 1},
				{4: 2, 5: 2, 6: 2},
				{3: 3, 4: 3, 5: 3},
			},
			expected: nil,
		},
	} {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			actual := UnorderedIntersect(test.input...)
			if !maps.Equal(test.expected, actual) {
				t.Errorf("expected %v got %v", test.expected, actual)
			}
		})
	}
}

func TestOrderedSunset(t *testing.T) {
	for _, test := range []struct {
		super    []int
		sub      []int
		expected bool
	}{
		{
			super:    []int{1, 2, 3, 4},
			sub:      []int{1, 2, 3, 4},
			expected: true,
		},
		{
			super:    []int{1, 2, 3, 4},
			sub:      []int{1, 2},
			expected: true,
		},
		{
			super:    []int{1, 2, 3, 4},
			sub:      []int{1, 2, 3, 4, 5},
			expected: false,
		},

		{
			super:    nil,
			sub:      []int{1, 2, 3, 4},
			expected: false,
		},
		{
			super:    []int{1, 2, 3, 4},
			sub:      nil,
			expected: true,
		},
		{
			super:    nil,
			sub:      nil,
			expected: true,
		},
	} {
		t.Run(fmt.Sprintf("%v <= %v", test.super, test.sub), func(t *testing.T) {
			actual := OrderedSubset(test.super, test.sub)
			if test.expected != actual {
				t.Errorf("expected %v got %v", test.expected, actual)
			}
		})
	}
}

func TestUnorderedSunset(t *testing.T) {
	for _, test := range []struct {
		super    map[int]int
		sub      map[int]int
		expected bool
	}{
		{
			super:    map[int]int{1: 1, 2: 1, 3: 1, 4: 1},
			sub:      map[int]int{1: 1, 2: 1, 3: 1, 4: 1},
			expected: true,
		},
		{
			super:    map[int]int{1: 1, 2: 1, 3: 1, 4: 1},
			sub:      map[int]int{1: 1, 2: 1},
			expected: true,
		},
		{
			super:    map[int]int{1: 1, 2: 1, 3: 1, 4: 1},
			sub:      map[int]int{1: 1, 2: 1, 3: 1, 4: 1, 5: 1},
			expected: false,
		},

		{
			super:    nil,
			sub:      map[int]int{1: 1, 2: 1, 3: 1, 4: 1},
			expected: false,
		},
		{
			super:    map[int]int{1: 1, 2: 1, 3: 1, 4: 1},
			sub:      nil,
			expected: true,
		},
		{
			super:    nil,
			sub:      nil,
			expected: true,
		},
	} {
		t.Run(fmt.Sprintf("%v <= %v", test.super, test.sub), func(t *testing.T) {
			actual := UnorderedSubset(test.super, test.sub)
			if test.expected != actual {
				t.Errorf("expected %v got %v", test.expected, actual)
			}
		})
	}
}
