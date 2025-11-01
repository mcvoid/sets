# Sets

A dead-simple library to handle quick and dirty sets.

## Install

```
go get github.com/mcvoid/sets
```

## What's up with sets?

We've all improvised set types in Go. Using a map like `map[MyObj]struct{}`, it works okay.
But if we want some set operations like union, you have to roll your own. And it doesn't
preserve order. You need a slice for that, but then you have to roll your own deduplication.

I did that all for you. And it'll work on all your custom improvised set types. Now that that's
over with, you can get on with the task at hand.

## Ordered vs Unordered

There's basically two sets of set functions in this library: ordered set functions, and
unordered set functions.

Unordered sets are maps, like given above. Any map can be made to act
like a set, and these functions act on these. Any function that has "Unordered" in its name
works on all maps. If you created your own type that uses a map as underlying storage,
that works too. These are:

* `UnorderedUnion`
* `UnorderedComplement`
* `UnorderedIntersect`
* `IsUnorderedSubsetOf`

Ordered sets are slices. Any slice of a comparable type (if it can be used as a map key,
it's comparable), these functions can acts on those as well. You just use the functions with
"Ordered" in their name. These are:

* `OrderedUnion`
* `OrderedComplement`
* `OrderedIntersect`
* `IsOrderedSubsetOf`

## How to use

These functions work with plain slices and maps:

```
a := sets.OrderedUnion([]string{"a", "b"}, []string{"b", "c"})
fmt.Println(a) // prints ["a", "b", "c"]

s1 := map[string]struct{}{"a": {}, "b": {}}
s2 := map[string]struct{}{"c": {}, "b": {}}
b := sets.UnorderedIntersect(s1, s2)
fmt.Println(b) // prints {"b":{}}
```

They can also be used for user-defined types based on slices and maps:

```
type MySet []string
type MyUSet map[int]any

...
s1 := MySet{"foo", "bar", "baz"}
s2 := MySet{"baz"}
c := sets.OrderedComplement(s1, s2) // is type MySet
fmt.Println(c)                      // prints ["foo", "bar"]

s3 := MyUSet{1: nil, 2: nil, 3: nil, 4: nil}
s4 := MyUSet{2: nil, 3: nil, 4: nil}
d := sets.IsUnorderedSubsetOf(s3, s4)
fmt.Println(d) // prints true
```

## Tips and Tricks

```
// append an item to a set
s1 = sets.OrderedUnion(s1, []string{"item to add"})

// prepend an item to a set
s1 = sets.OrderedUnion([]string{"item to add"}, s1)

// remove an item from a set
s1 = sets.OrderedComplement(s1, []string{"item to remove"})

// check if a set contains an item
containsItem := sets.IsOrderedSubsetOf(s1, []string{"item to find"})

// clone a set
s2 := sets.OrderedUnion(nil, s1)

// deduplicate a slice
s1 = sets.OrderedUnion(s1)
```

## License

Released under the MIT license.
