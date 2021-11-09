package collections

type (
	// MapKey represent a key present in a Map.
	MapKey interface{}
	// MapValue represent a value present in a Map.
	MapValue interface{}
	// Map is just a map.
	Map map[MapKey]MapValue
)

// Insert inserts a key-value pair into the map. If the map did not have this
// key present, nil is returned. If the map did have this key present, the value
// is updated, and the old value is returned.
func (m Map) Insert(k MapKey, v MapValue) (old MapValue) {
	old = m[k]
	m[k] = v

	return
}

// ContainsKey returns true if the map contains a key.
func (m Map) ContainsKey(k MapKey) (ok bool) {
	_, ok = m[k]

	return
}

// Remove removes a key from the map, returning the value at the key if the key
// was previously in the map.
func (m Map) Remove(k MapKey) (v MapValue) {
	v = m[k]
	delete(m, k)

	return
}

// IsEmpty returns true if the map contains no elements.
func (m Map) IsEmpty() bool {
	return len(m) == 0
}

// Get returns the value corresponding to the key.
func (m Map) Get(k MapKey) (v MapValue) {
	v = m[k]

	return
}

// Keys returns all keys in arbitrary order.
func (m Map) Keys() (mk []MapKey) {
	mk = make([]MapKey, 0)
	for k := range m {
		mk = append(mk, k)
	}

	return
}

// Clear clears the map, removing all keys and values.
func (m *Map) Clear() {
	*m = MakeMap()
}

// Values returns all the values in arbitrary order.
func (m Map) Values() (mv []MapValue) {
	mv = make([]MapValue, 0)
	for _, v := range m {
		mv = append(mv, v)
	}

	return
}

// ValuesToSet returns a set from the values in the map.
func (m Map) ValuesToSet() (s Set) {
	s = MakeSet()
	for _, v := range m {
		s.Insert(v)
	}

	return
}

// KeysToSet returns a set from the keys in the map.
func (m Map) KeysToSet() (s Set) {
	s = MakeSet()
	for k := range m {
		s.Insert(k)
	}

	return
}

// MakeMap creates a new map.
func MakeMap() (m Map) {
	m = make(map[MapKey]MapValue)

	return
}
