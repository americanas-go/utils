package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	set := MakeSet()

	ok := set.Insert("cat")
	assert.Equal(t, true, ok, "they should be equal")

	ok = set.Insert("cat")
	assert.Equal(t, false, ok, "they should be equal")
}

func TestContains(t *testing.T) {
	set := MakeSet("cat", "dog", "cow")

	ok := set.Contains("cat")
	assert.Equal(t, true, ok, "they should be equal")

	ok = set.Contains("buffalo")
	assert.Equal(t, false, ok, "they should be equal")
}

func TestRemove(t *testing.T) {
	set := MakeSet("cat", "dog", "cow")

	ok := set.Remove("cat")
	assert.Equal(t, true, ok, "they should be equal")

	ok = set.Remove("cat")
	assert.Equal(t, false, ok, "they should be equal")
}
