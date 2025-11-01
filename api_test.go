package sets_test

import (
	"maps"
	"slices"
	"testing"

	"github.com/mcvoid/sets"
)

func TestExamples(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {
		a := sets.OrderedUnion([]string{"a", "b"}, []string{"b", "c"})
		expected := []string{"a", "b", "c"}
		if !slices.Equal(a, expected) {
			t.Errorf("expected %v got %v", expected, a)
		}
	})

	t.Run("example 2", func(t *testing.T) {
		s1 := map[string]struct{}{"a": {}, "b": {}}
		s2 := map[string]struct{}{"c": {}, "b": {}}
		b := sets.UnorderedIntersect(s1, s2)
		expected := map[string]struct{}{"b": {}}
		if !maps.Equal(b, expected) {
			t.Errorf("expected %v got %v", expected, b)
		}
	})

	t.Run("example 3", func(t *testing.T) {
		type MySet []string
		s1 := MySet{"foo", "bar", "baz"}
		s2 := MySet{"baz"}
		c := sets.OrderedComplement(s1, s2)
		expected := MySet{"foo", "bar"}
		if !slices.Equal(c, expected) {
			t.Errorf("expected %v got %v", expected, c)
		}
	})

	t.Run("example 4", func(t *testing.T) {
		type MyUSet map[int]any
		s3 := MyUSet{1: nil, 2: nil, 3: nil, 4: nil}
		s4 := MyUSet{2: nil, 3: nil, 4: nil}
		d := sets.IsUnorderedSubsetOf(s3, s4)
		expected := true
		if d != expected {
			t.Errorf("expected %v got %v", expected, d)
		}
	})

	t.Run("append an item to a set", func(t *testing.T) {
		s1 := []string{"a", "b", "c"}
		s1 = sets.OrderedUnion(s1, []string{"item to add"})
		expected := []string{"a", "b", "c", "item to add"}
		if !slices.Equal(s1, expected) {
			t.Errorf("expected %v got %v", expected, s1)
		}
	})

	t.Run("prepend an item to a set", func(t *testing.T) {
		s1 := []string{"a", "b", "c"}
		s1 = sets.OrderedUnion([]string{"item to add"}, s1)
		expected := []string{"item to add", "a", "b", "c"}
		if !slices.Equal(s1, expected) {
			t.Errorf("expected %v got %v", expected, s1)
		}
	})

	t.Run("remove an item from a set", func(t *testing.T) {
		s1 := []string{"a", "b", "c", "item to remove"}
		s1 = sets.OrderedComplement(s1, []string{"item to remove"})
		expected := []string{"a", "b", "c"}
		if !slices.Equal(s1, expected) {
			t.Errorf("expected %v got %v", expected, s1)
		}
	})

	t.Run("check if a set contains an item", func(t *testing.T) {
		s1 := []string{"a", "b", "c"}
		containsItem := sets.IsOrderedSubsetOf(s1, []string{"item to find"})
		expected := false
		if containsItem != expected {
			t.Errorf("expected %v got %v", expected, s1)
		}
	})

	t.Run("clone a set", func(t *testing.T) {
		s1 := []string{"a", "b", "c", "item to remove"}
		s2 := sets.OrderedUnion(nil, s1)
		if !slices.Equal(s1, s2) {
			t.Errorf("expected %v got %v", s2, s1)
		}
	})

	t.Run("deduplicate a slice", func(t *testing.T) {
		s1 := []string{"a", "b", "c", "a"}
		s1 = sets.OrderedUnion(s1)
		expected := []string{"a", "b", "c"}
		if !slices.Equal(s1, expected) {
			t.Errorf("expected %v got %v", expected, s1)
		}
	})
}
