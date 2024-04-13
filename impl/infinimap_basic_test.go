package impl

import (
	"github.com/kuking/infinimap"
	"github.com/zeebo/assert"
	"os"
	"testing"
)

func TestHappyPath(t *testing.T) {
	tempFile, _ := os.CreateTemp(os.TempDir(), "infinimap")
	defer deferredCleanup(tempFile)

	imap, err := CreateInfinimap[uint64, string](tempFile.Name(), NewCreateParameters())
	assert.NoError(t, err)

	previous, replaced, err := imap.Put(1, "Uno")
	assert.Equal(t, "", previous)
	assert.False(t, replaced)
	assert.NoError(t, err)

	value, found := imap.Get(1)
	assert.True(t, found)
	assert.Equal(t, "Uno", value)

	value, found = imap.Get(2)
	assert.False(t, found)

	assert.Equal(t, 1, imap.Count())

	assert.Equal(t, getKeys(imap), []uint64{1})
	assert.Equal(t, getValues(imap), []string{"Uno"})

	didEach := false
	err = imap.Each(func(k uint64, v string) bool {
		assert.Equal(t, 1, k)
		assert.Equal(t, "Uno", v)
		didEach = true
		return true
	})
	assert.NoError(t, err)
	assert.True(t, didEach)

	previous, replaced, err = imap.Put(1, "Uno 2.0")
	assert.Equal(t, "Uno", previous)
	assert.True(t, replaced)
	assert.NoError(t, err)

	assert.Equal(t, getKeys(imap), []uint64{1})
	assert.Equal(t, getValues(imap), []string{"Uno 2.0"})

	assert.True(t, imap.Delete(1))

	value, found = imap.Get(1)
	assert.False(t, found)
	assert.Equal(t, "", value)

	assert.Equal(t, 0, imap.Count())

	didEach = false
	err = imap.Each(func(k uint64, v string) bool {
		didEach = true
		return true
	})

	assert.Nil(t, getKeys(imap))
	assert.Nil(t, getValues(imap))
}

// -----------------------------------------------------------------------------------------------------------------------------------------------------------

func getKeys[K comparable, V any](imap infinimap.InfiniMap[K, V]) []K {
	var res []K
	for key := range imap.Keys() {
		res = append(res, key)
	}
	return res
}

func getValues[K comparable, V any](imap infinimap.InfiniMap[K, V]) []V {
	var res []V
	for value := range imap.Values() {
		res = append(res, value)
	}
	return res
}

func deferredCleanup(file *os.File) {
	if file != nil {
		_ = os.Remove(file.Name())
	}
}
