package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetInsert(t *testing.T) {
	set := MakeSet()

	ok := set.Insert("cat")
	assert.Equal(t, true, ok, "they should be equal")

	ok = set.Insert("cat")
	assert.Equal(t, false, ok, "they should be equal")
}

func TestSetContains(t *testing.T) {
	set := MakeSet("cat", "dog", "cow")

	ok := set.Contains("cat")
	assert.Equal(t, true, ok, "they should be equal")

	ok = set.Contains("buffalo")
	assert.Equal(t, false, ok, "they should be equal")
}

func TestSetRemove(t *testing.T) {
	set := MakeSet("cat", "dog", "cow")

	ok := set.Remove("cat")
	assert.Equal(t, true, ok, "they should be equal")

	ok = set.Remove("cat")
	assert.Equal(t, false, ok, "they should be equal")
}

func TestSetIntersection(t *testing.T) {
	set1 := MakeSet("cat", "dog", "cow")
	set2 := MakeSet("cat", "duck", "bull")
	intersection := set1.Intersection(set2)

	assert.Contains(t, intersection, "cat")
	assert.NotContains(t, intersection, "dog")
	assert.NotContains(t, intersection, "cow")
	assert.NotContains(t, intersection, "duck")
	assert.NotContains(t, intersection, "bull")
}

func TestSetSymmetricDifference(t *testing.T) {
	set1 := MakeSet(1, 2, 3)
	set2 := MakeSet(4, 2, 3, 4)

	symDiff := set1.SymmetricDifference(set2)
	assert.Contains(t, symDiff, 1)
	assert.Contains(t, symDiff, 4)
	assert.NotContains(t, symDiff, 2)
	assert.NotContains(t, symDiff, 3)
}

func TestSetDifference(t *testing.T) {
	set1 := MakeSet(1, 2, 3)
	set2 := MakeSet(4, 2, 3, 4)

	diff1 := set1.Difference(set2)
	assert.Contains(t, diff1, 1)
	assert.NotContains(t, diff1, 2)
	assert.NotContains(t, diff1, 3)
	assert.NotContains(t, diff1, 4)

	diff2 := set2.Difference(set1)
	assert.Contains(t, diff2, 4)
	assert.NotContains(t, diff2, 1)
	assert.NotContains(t, diff2, 2)
	assert.NotContains(t, diff2, 3)
}

func TestSetUnion(t *testing.T) {
	set1 := MakeSet("cat", "dog", "cow")
	set2 := MakeSet("cat", "duck", "bull")
	intersection := set1.Union(set2)

	assert.Contains(t, intersection, "cat")
	assert.Contains(t, intersection, "dog")
	assert.Contains(t, intersection, "cow")
	assert.Contains(t, intersection, "duck")
	assert.Contains(t, intersection, "bull")
}

func TestSetCollect(t *testing.T) {
	els := []interface{}{"cat", "cow", 10, true, false, 10, true, false, "cat", "cow"}
	setValues := MakeSet(els...).Collect()

	assert.Contains(t, setValues, "cat")
	assert.Contains(t, setValues, 10)
	assert.Contains(t, setValues, "cow")
	assert.Contains(t, setValues, true)
	assert.Contains(t, setValues, false)
}

func TestSetClear(t *testing.T) {
	set := MakeSet("cat", "dog", "cow")

	set.Clear()
	assert.Equal(t, 0, len(set), "they should be equal")
}

func TestSetIsEmpty(t *testing.T) {
	set := MakeSet()
	assert.Equal(t, true, set.IsEmpty(), "they should be equal")

	set.Insert("cat")
	assert.Equal(t, false, set.IsEmpty(), "they should be equal")
}
