package collections

type (
	// SetValue ...
	SetValue interface{}
	// Set ...
	Set map[SetValue]bool
)

// Insert adds a value to the set. If the set did not have this value present,
// true is returned. If the set did have this value present, false is returned.
func (s Set) Insert(v SetValue) (ok bool) {
	_, ok = s[v]
	s[v] = true
	return !ok
}

// Contains returns true if the set contains a value.
func (s Set) Contains(v SetValue) (ok bool) {
	_, ok = s[v]
	return
}

// Remove a value from the set. Returns true if the value was present in the
// set.
func (s Set) Remove(v SetValue) (ok bool) {
	_, ok = s[v]
	delete(s, v)
	return ok
}

// MakeSet creates a new Set.
func MakeSet(values ...SetValue) (s Set) {
	s = make(map[SetValue]bool)
	for _, v := range values {
		s[v] = true
	}

	return
}

// TODO: add Intersection, Difference, SymmetricDifference
