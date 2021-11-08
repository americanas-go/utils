package collections

type (
	// Set ...
	Set map[interface{}]bool
)

// Insert adds a value to the set. If the set did not have this value present,
// true is returned. If the set did have this value present, false is returned.
func (s Set) Insert(v interface{}) (ok bool) {
	_, ok = s[v]
	s[v] = true

	return !ok
}

// Contains returns true if the set contains a value.
func (s Set) Contains(v interface{}) (ok bool) {
	_, ok = s[v]

	return
}

// Remove a value from the set. Returns true if the value was present in the
// set.
func (s Set) Remove(v interface{}) (ok bool) {
	_, ok = s[v]
	if ok {
		delete(s, v)
	}

	return
}

// Clear clears the set, removing all values.
func (s *Set) Clear() {
	*s = make(map[interface{}]bool)
}

// Collect returns a slice of all the elements present in the set.
func (s Set) Collect() (vs []interface{}) {
	vs = make([]interface{}, 0)

	for k := range s {
		vs = append(vs, k)
	}

	return
}

// Union returns the set of values that are in s or in s2, without duplicates.
func (s Set) Union(s2 Set) (set Set) {
	set = make(map[interface{}]bool)

	for k := range s {
		set[k] = true
	}

	for k := range s2 {
		set[k] = true
	}

	return
}

// Difference returns the set of values that are in s but not in s2.
func (s Set) Difference(s2 Set) (set Set) {
	set = make(map[interface{}]bool)

	for k := range s {
		if ok := s2.Contains(k); !ok {
			set[k] = true
		}
	}

	return
}

// SymmetricDifference returns the set of values that are in s or in s2 but not
// in both.
func (s Set) SymmetricDifference(s2 Set) (set Set) {
	set = make(map[interface{}]bool)

	for k := range s {
		if ok := s2.Contains(k); !ok {
			set[k] = true
		}
	}

	for k := range s2 {
		if ok := s.Contains(k); !ok {
			set[k] = true
		}
	}

	return
}

// Intersection returns the set of values that are both in s and s2.
func (s Set) Intersection(s2 Set) (set Set) {
	set = make(map[interface{}]bool)

	for k := range s {
		if ok := s2.Contains(k); ok {
			set[k] = true
		}
	}

	return
}

// IsEmpty returns true if the set contains no elements.
func (s Set) IsEmpty() bool {
	return len(s) == 0
}

// MakeSet creates a new Set.
func MakeSet(values ...interface{}) (s Set) {
	s = make(map[interface{}]bool)
	for _, v := range values {
		s[v] = true
	}

	return
}

// TODO: add SymmetricDifference
