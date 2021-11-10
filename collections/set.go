package collections

type (
	// SetItem represent a item present in a Set.
	SetItem interface{}
	// SetColl is a slice of items present in a Set.
	SetColl []SetItem
	// Set is just a set.
	Set map[SetItem]void

	void struct{}
)

// Insert adds a value to the set. If the set did not have this value present,
// true is returned. If the set did have this value present, false is returned.
func (s Set) Insert(v SetItem) (ok bool) {
	_, ok = s[v]
	s[v] = void{}

	return !ok
}

// Contains returns true if the set contains a value.
func (s Set) Contains(si SetItem) (ok bool) {
	_, ok = s[si]

	return
}

// Remove a value from the set. Returns true if the value was present in the
// set.
func (s Set) Remove(si SetItem) (ok bool) {
	_, ok = s[si]
	if ok {
		delete(s, si)
	}

	return
}

// Clear clears the set, removing all values.
func (s *Set) Clear() {
	*s = MakeSet()
}

// Collect returns a slice of all the elements present in the set in arbitrary
// order.
func (s Set) Collect() (vs SetColl) {
	vs = make([]SetItem, 0)

	for k := range s {
		vs = append(vs, k)
	}

	return
}

// Union returns the set of values that are in s or in s2, without duplicates.
func (s Set) Union(s2 Set) (set Set) {
	set = MakeSet()

	for k := range s {
		set[k] = void{}
	}

	for k := range s2 {
		set[k] = void{}
	}

	return
}

// Difference returns the set of values that are in s but not in s2.
func (s Set) Difference(s2 Set) (set Set) {
	set = MakeSet()

	for k := range s {
		if ok := s2.Contains(k); !ok {
			set[k] = void{}
		}
	}

	return
}

// SymmetricDifference returns the set of values that are in s or in s2 but not
// in both.
func (s Set) SymmetricDifference(s2 Set) (set Set) {
	set = MakeSet()

	for k := range s {
		if ok := s2.Contains(k); !ok {
			set[k] = void{}
		}
	}

	for k := range s2 {
		if ok := s.Contains(k); !ok {
			set[k] = void{}
		}
	}

	return
}

// Intersection returns the set of values that are both in s and s2.
func (s Set) Intersection(s2 Set) (set Set) {
	set = MakeSet()

	for k := range s {
		if ok := s2.Contains(k); ok {
			set[k] = void{}
		}
	}

	return
}

// IsEmpty returns true if the set contains no elements.
func (s Set) IsEmpty() bool {
	return len(s) == 0
}

// MakeSet creates a new Set.
func MakeSet(sx ...SetItem) (s Set) {
	s = make(map[SetItem]void)
	for _, si := range sx {
		s[si] = void{}
	}

	return
}
