package collections_test

import (
	"testing"

	coll "github.com/americanas-go/utils/collections"
	"github.com/stretchr/testify/assert"
)

// ====================
// Tests
// ====================

func TestMap_Keys(t *testing.T) {
	mp := coll.MakeMap()
	mp.Insert(1, "cat")
	mp.Insert(2, "bat")

	ks := mp.Keys()
	assert.ElementsMatch(t, []coll.MapKey{1, 2}, ks, "they should be equal")
}

func TestMap_Values(t *testing.T) {
	mp := coll.MakeMap()
	mp.Insert(1, "cat")
	mp.Insert(2, "bat")

	vs := mp.Values()
	assert.ElementsMatch(t, []coll.MapValue{"cat", "bat"}, vs, "they should be equal")
}

func TestMap_Remove(t *testing.T) {
	mp := coll.MakeMap()
	mp.Insert(10, "cat")

	v := mp.Remove(10)
	assert.Equal(t, "cat", v, "they should be equal")

	v = mp.Remove(10)
	assert.Equal(t, nil, v, "they should be equal")
}

func TestMap_Get(t *testing.T) {
	mp := coll.MakeMap()

	mp.Insert(10, "cat")
	assert.Equal(t, "cat", mp.Get(10), "they should be equal")

	mp.Remove(10)
	assert.Equal(t, nil, mp.Get(10), "they should be equal")
}

func TestMap_ContainsKey(t *testing.T) {
	mp := coll.MakeMap()
	mp.Insert("cat", 10)
	mp.Insert("bat", 11)
	mp.Insert("frog", 12)

	ok := mp.ContainsKey("cat")
	assert.Equal(t, true, ok, "they should be equal")

	ok = mp.ContainsKey("buffalo")
	assert.Equal(t, false, ok, "they should be equal")
}

func TestMap_Clear(t *testing.T) {
	mp := coll.MakeMap()
	mp.Insert(10, "cat")
	mp.Insert(11, "bat")
	mp.Insert(12, "frog")

	mp.Clear()
	assert.Equal(t, 0, len(mp), "they should be equal")
}

func TestMap_Insert(t *testing.T) {
	mp := coll.MakeMap()

	assert.Equal(t, nil, mp.Insert(10, "cat"), "they should be equal")
	assert.Equal(t, "cat", mp.Insert(10, "bat"), "they should be equal")
	assert.Equal(t, "bat", mp.Get(10), "they should be equal")
}

func TestMap_IsEmpty(t *testing.T) {
	mp := coll.MakeMap()
	assert.Equal(t, true, mp.IsEmpty(), "they should be equal")

	mp.Insert(10, "cat")
	assert.Equal(t, false, mp.IsEmpty(), "they should be equal")
}

func TestMap_ValuesToSet(t *testing.T) {
	mp := coll.MakeMap()
	mp.Insert(10, "cat")
	mp.Insert(11, "bat")
	mp.Insert(12, "frog")

	s := mp.ValuesToSet().Collect()
	assert.ElementsMatch(t, s, coll.SetSlice{"cat", "bat", "frog"})
}

func TestMap_KeysToSet(t *testing.T) {
	mp := coll.MakeMap()
	mp.Insert(10, "cat")
	mp.Insert(11, "bat")
	mp.Insert(12, "frog")

	s := mp.KeysToSet().Collect()
	assert.ElementsMatch(t, s, coll.SetSlice{10, 11, 12})
}
