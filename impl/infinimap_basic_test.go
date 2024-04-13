package impl

import (
	"fmt"
	"github.com/kuking/infinimap"
	"github.com/zeebo/assert"
	"log"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
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

func TestBasicDrill(t *testing.T) {
	tempFile, _ := os.CreateTemp(os.TempDir(), "infinimap")
	defer deferredCleanup(tempFile)

	records := uint64(1_000_000)

	imap, err := CreateInfinimap[uint64, string](tempFile.Name(), NewCreateParameters().WithCapacity(int(records*10)))
	assert.NoError(t, err)

	t0 := time.Now()
	for i := uint64(0); i < records; i++ {
		_, _, err := imap.Put(i, strconv.Itoa(int(i)))
		assert.NoError(t, err)
	}
	elapsed := time.Since(t0)
	log.Printf("Took %v to insert %.1fM records or %.2fK records/s",
		elapsed.Truncate(time.Microsecond), float64(records)/1000.0/1000.0, float64(records)/float64(elapsed.Seconds())/1000.0)

	t0 = time.Now()
	for i := uint64(0); i < records; i++ {
		val, found := imap.Get(i)
		assert.True(t, found)
		assert.Equal(t, strconv.Itoa(int(i)), val)
	}
	elapsed = time.Since(t0)
	log.Printf("Took %v to read %.1fM records or %.2fK records/s",
		elapsed.Truncate(time.Microsecond), float64(records)/1000.0/1000.0, float64(records)/float64(elapsed.Seconds())/1000.0)

	t0 = time.Now()
	c := 0
	for _ = range imap.Keys() {
		c++
	}
	assert.Equal(t, int(records), c)
	elapsed = time.Since(t0)
	log.Printf("Took %v to read the keys of %.1fM records or %.2fK keys/s",
		elapsed.Truncate(time.Microsecond), float64(records)/1000.0/1000.0, float64(records)/float64(elapsed.Seconds())/1000.0)

	t0 = time.Now()
	c = 0
	for _ = range imap.Values() {
		c++
	}
	assert.Equal(t, int(records), c)
	elapsed = time.Since(t0)
	log.Printf("Took %v to read the values of %.1fM records or %.2fK values/s",
		elapsed.Truncate(time.Microsecond), float64(records)/1000.0/1000.0, float64(records)/float64(elapsed.Seconds())/1000.0)
}

func TestBasicCollisions(t *testing.T) {
	tempFile, _ := os.CreateTemp(os.TempDir(), "infinimap")
	defer deferredCleanup(tempFile)

	imap, err := CreateInfinimap[uint64, string](tempFile.Name(), NewCreateParameters().WithCapacity(1100))
	assert.NoError(t, err)

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	var keys []uint64
	gomap := map[uint64]string{}

	for i := uint64(0); i < 1000; i++ {
		k := rnd.Uint64()
		v := fmt.Sprintf("Val %d", k)
		_, _, err := imap.Put(k, v)
		assert.NoError(t, err)
		gomap[k] = v
		keys = append(keys, k)
	}

	for i := 0; i < 1_000_000; i++ {
		if len(keys) == 0 {
			continue
		}
		keyIndex := rnd.Intn(len(keys))
		key := keys[keyIndex]

		// println(i, key)

		if val, found := imap.Get(key); found {
			// asserts has the expected value
			expectedVal := fmt.Sprintf("Val %d", key)
			assert.Equal(t, expectedVal, val)
			assert.Equal(t, expectedVal, gomap[key])

			// deletes
			assert.True(t, imap.Delete(key))
			delete(gomap, key)
			keys = append(keys[:keyIndex], keys[keyIndex+1:]...)

			// Add a new random
			newKey := rnd.Uint64()
			newValue := fmt.Sprintf("Val %d", newKey)
			_, _, err = imap.Put(newKey, newValue)
			assert.NoError(t, err)
			gomap[newKey] = newValue
			keys = append(keys, newKey)
		}
	}

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
