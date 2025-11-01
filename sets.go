// A simple library for managing slices and maps as sets.
//
// All OrderedX functions (eg: OrderedUnion) work with slices
// or any type based on a slice as long as the slice elements are
// comparable. These slices are treated as ordered sets.
//
// For slices, all operations here automatically remove duplicates and
// preserve the order of the elements in the original slice.
// Whenever possible, they also seek to avoid allocations.
// So they preserve nilness, they don't make a new set out of a slice
// which already adheres to the set property, etc. All operations are
// variadic and applied left-to-right (index 0 and up).
//
// It's "naive" in that it has no tricks to speed up performance.
// It will deliberately sacrifice cycles to favor avoiding
// an allocation, for example.
//
// All UnorderedX functions (eg: UnorderedUnion) work with maps and
// any type based on a map. Since all map keys are already comparable,
// there is no restriction to their use. These maps are treated as unordered
// sets.
//
// For maps, no order is preserved and the operations are faster.
// Allocations are still avoided. Map keys are treated as the set elements,
// not their values. As such, values found in duplicate keys will be clobbered.
// This mimics the conventional use of maps as ad-hoc sets, so they should
// work in the same way. Whenever possible, the values for the first given
// set will be in the output.
//
// As "nil" is common nomenclature for the empty set ({}, or 0), nil
// is accepted as input to mean an empty slice or empty map, and
// empty set results such as the intersection of disjoint sets are
// similarly returned as nil.
//
// All sets are treated as if immutable - no modifications will occur
// on any of these operations, and instead copies will be returned.
package sets

import (
	"maps"
	"slices"
)

// Adds values to a set if they are not already contained therein.
// The output will be a new slice with all the elements of the
// original slices plus all of vals, without duplicates. Preserves order.
func setAppend[S ~[]E, E comparable](s S, vals ...E) S {
	set := slices.Clone(s)
	for _, val := range vals {
		if !slices.Contains(set, val) {
			set = append(set, val)
		}
	}
	return set
}

// Makes a slice enforce the set characteristic. The output will be
// a new slice with duplicates removed. Preserves order.
func makeSet[S ~[]E, E comparable](s S) S {
	// test the set characteristic by comparing every element
	// to every other element
	isSet := true
	for i, e1 := range s {
		for _, e2 := range s[i+1:] {
			if e1 == e2 {
				isSet = false
			}
		}
	}
	if isSet {
		return s
	}
	// there's some duplicates, so dedup by making a copy
	return setAppend(S(nil), s...)
}

// Produces the union of slices as sets. The output will
// be a new slice containing the elements of all input slices,
// without duplicates. Preserves order.
func OrderedUnion[S ~[]E, E comparable](sets ...S) S {
	var union S
	for _, set := range sets {
		union = setAppend(union, set...)
	}
	return union
}

// Produces the union of maps as sets, where
// the maps' keys are the set's members. The output will be a
// new map containing the keys of all the sets. Union is applied
// in argument order, with the values of repeated keys clobbering
// each other, so the last value applied for a key us what's in the
// set.
func UnorderedUnion[S ~map[K]V, K comparable, V any](sets ...S) S {
	var union map[K]V
	for i := range sets {
		// reverse order traversal makes it that
		// the first given arg overwrites the
		// previous values.
		set := sets[len(sets)-1-i]
		if set == nil {
			continue
		}
		if union == nil {
			// don't allocate until we need to
			// we'd prefer to return nil if no elements
			// end up in the set
			union = map[K]V{}
		}
		maps.Copy(union, set)
	}
	return union
}

// Produces the relative complements of slices as sets. The output
// will be all elements of the first argument which do not appear in any
// the the rest of the input slices, without suplicates.
// Preserves order.
//
// Each argument is getting complemented to the relative
// complement of the previous two arguments. Applies in
// left-to-right order, so OrderedComplement(a, b, c) is the
// set of all elements in a not in b, and also not in c.
func OrderedComplement[S ~[]E, E comparable](sets ...S) S {
	if len(sets) == 0 {
		return nil
	}
	if len(sets) == 1 {
		return makeSet(sets[0])
	}
	s1 := makeSet(sets[0])
	for _, s2 := range sets[1:] {
		// if running complement is empty, we can stop the loop early
		// since there's nothing left to remove.
		if len(s1) == 0 {
			return nil
		}
		var complement S
		for _, v := range s1 {
			if !slices.Contains(s2, v) {
				complement = setAppend(complement, v)
			}
		}
		s1 = complement
	}
	return s1
}

// Produces the relative complement of maps as sets, where
// the maps' keys are the set's members. The output will be a new map
// with the keys of the first argument which do not appear in the rest
// of the input maps.
func UnorderedComplement[S ~map[K]V, K comparable, V any](sets ...S) S {
	if len(sets) == 0 {
		return nil
	}
	if len(sets) == 1 {
		return sets[0]
	}
	complement := maps.Clone(sets[0])
	for _, s2 := range sets[1:] {
		// if running complement is empty, we can stop the loop early
		// since there's nothing left to remove.
		if len(complement) == 0 {
			return nil
		}
		for e := range s2 {
			delete(complement, e)
		}
	}
	if len(complement) == 0 {
		return nil
	}
	return complement
}

// Produces the intersection of slices as sets. The output
// will be a new slice containing the elements in all input
// slices, without duplicates. Preserves order.
//
// Each argument is getting intersected with the intersection
// of the previous two arguments.
func OrderedIntersect[S ~[]E, E comparable](sets ...S) S {
	if len(sets) == 0 {
		return nil
	}
	if len(sets) == 1 {
		return makeSet(sets[0])
	}
	s1 := makeSet(sets[0])
	for _, s2 := range sets[1:] {
		// if runing intersection is empty, we can stop the loop early
		// since there's nothing we can add to the set
		if len(s1) == 0 {
			return nil
		}
		var intersection S
		for _, v := range s1 {
			if slices.Contains(s2, v) {
				intersection = setAppend(intersection, v)
			}
		}
		s1 = intersection
	}
	return s1
}

// Produces the intersection of maps as sets, where
// the maps' keys are the set's members. The output will be
// a new map where the keys are the ones of the first given set which also
// appear in every other given set.
//
// The values are those of the first set for they keys still remaining
// after intersection.
func UnorderedIntersect[S ~map[K]V, K comparable, V any](sets ...S) S {
	if len(sets) == 0 {
		return nil
	}
	if len(sets) == 1 {
		return sets[0]
	}
	intersection := maps.Clone(sets[0])
	for _, s2 := range sets[1:] {
		// if runing intersection is empty, we can stop the loop early
		// since there's nothing we can add to the set
		if len(intersection) == 0 {
			return nil
		}
		if s2 == nil {
			return nil
		}
		for e := range intersection {
			if _, exists := s2[e]; !exists {
				delete(intersection, e)
			}
		}
	}
	if len(intersection) == 0 {
		return nil
	}
	return intersection
}

// Returns whether sub is a subset of super. It is a subset
// if and only if every element in sub is also in super.
//
// This version is meant to be for slices.
func OrderedSubset[S ~[]E, E comparable](super, sub S) bool {
	for _, e := range sub {
		if !slices.Contains(super, e) {
			return false
		}
	}
	return true
}

// Returns whether sub is a subset of super. It is a subset
// if and only if every key in sub is also in super.
//
// This vesrion is meant to be for maps.
func UnorderedSubset[S ~map[K]V, K comparable, V any](super, sub S) bool {
	for k := range sub {
		if _, exists := super[k]; !exists {
			return false
		}
	}
	return true
}
