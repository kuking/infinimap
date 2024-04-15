package V1

import (
	"github.com/zeebo/assert"
	"os"
	"testing"
)

func TestShrinkExpandExpand(t *testing.T) {
	tempFile, _ := os.CreateTemp(os.TempDir(), "infinimap")
	defer deferredCleanup(tempFile)

	imap, err := Create[uint64, string](tempFile.Name(), NewCreateParameters())
	assert.NoError(t, err)
	assert.Equal(t, 0xf4231b1e3cf, imap.BytesAvailable())
	assert.NoError(t, imap.Shrink())

	_, _, err = imap.Put(123, "123")
	assert.Error(t, err)
	assert.Equal(t, 0, imap.BytesAvailable())

	assert.NoError(t, imap.Expand(10_000_000))
	assert.Equal(t, 10_000_000, imap.BytesAvailable())
	_, _, err = imap.Put(123, "123")
	assert.NoError(t, err)
}
