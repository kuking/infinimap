package impl

import (
	"github.com/kuking/infinimap"
	"os"
)

type createParams struct {
	capacity    int
	totalFileMb int
	fileMode    os.FileMode
	concurrency bool
	compression infinimap.Compression
	hashing     infinimap.Hashing
}

func (c *createParams) WithCapacity(capacity int) infinimap.CreateParameters {
	c.capacity = capacity
	return c
}

func (c *createParams) GetCapacity() int {
	return c.capacity
}

func (c *createParams) WithFileSizeMB(fileSizeMB int) infinimap.CreateParameters {
	c.totalFileMb = fileSizeMB
	return c
}

func (c *createParams) GetFileSizeMB() int {
	return c.totalFileMb
}

func (c *createParams) WithFileMode(fileMode os.FileMode) infinimap.CreateParameters {
	c.fileMode = fileMode
	return c
}

func (c *createParams) GetFileMode() os.FileMode {
	return c.fileMode
}

func (c *createParams) WithConcurrency(concurrent bool) infinimap.CreateParameters {
	c.concurrency = concurrent
	return c
}

func (c *createParams) GetConcurrency() bool {
	return c.concurrency
}

func (c *createParams) WithCompression(compression infinimap.Compression) infinimap.CreateParameters {
	c.compression = compression
	return c
}

func (c *createParams) GetCompression() infinimap.Compression {
	return c.compression
}

func (c *createParams) WithHashing(hashing infinimap.Hashing) infinimap.CreateParameters {
	c.hashing = hashing
	return c
}

func (c *createParams) GetHashing() infinimap.Hashing {
	return c.hashing
}

func NewCreateParameters() infinimap.CreateParameters {
	return &createParams{
		capacity:    5_000_000,
		totalFileMb: 16_000,
		fileMode:    0644,
		concurrency: false,
		compression: infinimap.COMPRESSION_NONE,
		hashing:     infinimap.XX128_HASHING,
	}
}
