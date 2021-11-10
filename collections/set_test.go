package collections_test

import (
	"fmt"
	"testing"

	coll "github.com/americanas-go/utils/collections"
	"github.com/stretchr/testify/assert"
)

// ====================
// Tests
// ====================

func TestSet_Insert(t *testing.T) {
	set := coll.MakeSet()

	ok := set.Insert("cat")
	assert.Equal(t, true, ok, "they should be equal")

	ok = set.Insert("cat")
	assert.Equal(t, false, ok, "they should be equal")
}

func TestSet_Contains(t *testing.T) {
	set := coll.MakeSet("cat", "dog", "cow")

	ok := set.Contains("cat")
	assert.Equal(t, true, ok, "they should be equal")

	ok = set.Contains("buffalo")
	assert.Equal(t, false, ok, "they should be equal")
}

func TestSet_Remove(t *testing.T) {
	set := coll.MakeSet("cat", "dog", "cow")

	ok := set.Remove("cat")
	assert.Equal(t, true, ok, "they should be equal")

	ok = set.Remove("cat")
	assert.Equal(t, false, ok, "they should be equal")
}

func TestSet_Intersection(t *testing.T) {
	set1 := coll.MakeSet("cat", "dog", "cow")
	set2 := coll.MakeSet("cat", "duck", "bull")
	intersection := set1.Intersection(set2)

	assert.Contains(t, intersection, "cat")
	assert.NotContains(t, intersection, "dog")
	assert.NotContains(t, intersection, "cow")
	assert.NotContains(t, intersection, "duck")
	assert.NotContains(t, intersection, "bull")
}

func TestSet_SymmetricDifference(t *testing.T) {
	set1 := coll.MakeSet(1, 2, 3)
	set2 := coll.MakeSet(4, 2, 3, 4)

	symDiff := set1.SymmetricDifference(set2)
	assert.Contains(t, symDiff, 1)
	assert.Contains(t, symDiff, 4)
	assert.NotContains(t, symDiff, 2)
	assert.NotContains(t, symDiff, 3)
}

func TestSet_Difference(t *testing.T) {
	set1 := coll.MakeSet(1, 2, 3)
	set2 := coll.MakeSet(4, 2, 3, 4)

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

func TestSet_Union(t *testing.T) {
	set1 := coll.MakeSet("cat", "dog", "cow")
	set2 := coll.MakeSet("cat", "duck", "bull")
	union := set1.Union(set2)

	assert.Contains(t, union, "cat")
	assert.Contains(t, union, "dog")
	assert.Contains(t, union, "cow")
	assert.Contains(t, union, "duck")
	assert.Contains(t, union, "bull")
}

func TestSet_Collect(t *testing.T) {
	els := []coll.SetItem{"cat", "cow", 10, true, false, 10, true, false, "cat", "cow"}
	setValues := coll.MakeSet(els...).Collect()

	assert.Contains(t, setValues, "cat")
	assert.Contains(t, setValues, 10)
	assert.Contains(t, setValues, "cow")
	assert.Contains(t, setValues, true)
	assert.Contains(t, setValues, false)
}

func TestSet_Clear(t *testing.T) {
	set := coll.MakeSet("cat", "dog", "cow")

	set.Clear()
	assert.Equal(t, 0, len(set), "they should be equal")
}

func TestSet_IsEmpty(t *testing.T) {
	set := coll.MakeSet()
	assert.Equal(t, true, set.IsEmpty(), "they should be equal")

	set.Insert("cat")
	assert.Equal(t, false, set.IsEmpty(), "they should be equal")
}

// ====================
// Examples
// ====================

func ExampleSet_Insert() {
	set := coll.MakeSet()

	ok := set.Insert("cat")
	fmt.Println(ok)

	ok = set.Insert("cat")
	fmt.Println(ok)

	// Output:
	// true
	// false
}

func ExampleSet_Contains() {
	set := coll.MakeSet("cat", "dog", "cow")

	ok := set.Contains("cat")
	fmt.Println(ok)

	ok = set.Contains("buffalo")
	fmt.Println(ok)

	// Output:
	// true
	// false
}

func ExampleSet_Remove() {
	set := coll.MakeSet("cat", "dog", "cow")

	ok := set.Remove("cat")
	fmt.Println(ok)

	ok = set.Remove("cat")
	fmt.Println(ok)

	// Output:
	// true
	// false
}

func ExampleSet_Intersection() {
	set1 := coll.MakeSet("cat", "dog", "cow")
	set2 := coll.MakeSet("cat", "duck", "bull")
	intersection := set1.Intersection(set2)

	fmt.Println(intersection.Contains("cat"))
	fmt.Println(intersection.Contains("dog"))
	fmt.Println(intersection.Contains("cow"))
	fmt.Println(intersection.Contains("duck"))
	fmt.Println(intersection.Contains("bull"))
	// Output:
	// true
	// false
	// false
	// false
	// false
}

func ExampleSet_SymmetricDifference() {
	set1 := coll.MakeSet(1, 2, 3)
	set2 := coll.MakeSet(4, 2, 3, 4)
	symDiff := set1.SymmetricDifference(set2)

	fmt.Println(symDiff.Contains(1))
	fmt.Println(symDiff.Contains(4))
	fmt.Println(symDiff.Contains(2))
	fmt.Println(symDiff.Contains(3))
	// Output:
	// true
	// true
	// false
	// false
}

func ExampleSet_Difference() {
	set1 := coll.MakeSet(1, 2, 3)
	set2 := coll.MakeSet(4, 2, 3, 4)

	diff1 := set1.Difference(set2)
	fmt.Println(diff1.Contains(1))
	fmt.Println(diff1.Contains(2))
	fmt.Println(diff1.Contains(3))
	fmt.Println(diff1.Contains(4))

	diff2 := set2.Difference(set1)
	fmt.Println(diff2.Contains(4))
	fmt.Println(diff2.Contains(1))
	fmt.Println(diff2.Contains(2))
	fmt.Println(diff2.Contains(3))

	// Output:
	// true
	// false
	// false
	// false
	// true
	// false
	// false
	// false
}

func ExampleSet_Union() {
	set1 := coll.MakeSet("cat", "dog", "cow")
	set2 := coll.MakeSet("cat", "duck", "bull")
	union := set1.Union(set2)

	fmt.Println(union.Contains("cat"))
	fmt.Println(union.Contains("dog"))
	fmt.Println(union.Contains("cow"))
	fmt.Println(union.Contains("duck"))
	fmt.Println(union.Contains("bull"))
	// Output:
	// true
	// true
	// true
	// true
	// true
}

func ExampleSet_Clear() {
	set := coll.MakeSet("cat", "dog", "cow")

	set.Clear()
	fmt.Println(len(set))
	// Output:
	// 0
}

func ExampleSet_IsEmpty() {
	set := coll.MakeSet()
	fmt.Println(set.IsEmpty())

	set.Insert("cat")
	fmt.Println(set.IsEmpty())

	// Output:
	// true
	// false
}
