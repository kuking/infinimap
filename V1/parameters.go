package V1

import (
	"os"
)

type createParams struct {
	capacity    int
	totalFileMb int
	fileMode    os.FileMode
	concurrency bool
	compression Compression
	hashing     Hashing
}

func (c *createParams) WithCapacity(capacity int) CreateParameters {
	c.capacity = capacity
	return c
}

func (c *createParams) GetCapacity() int {
	return c.capacity
}

func (c *createParams) WithFileSizeMB(fileSizeMB int) CreateParameters {
	c.totalFileMb = fileSizeMB
	return c
}

func (c *createParams) GetFileSizeMB() int {
	return c.totalFileMb
}

func (c *createParams) WithFileMode(fileMode os.FileMode) CreateParameters {
	c.fileMode = fileMode
	return c
}

func (c *createParams) GetFileMode() os.FileMode {
	return c.fileMode
}

func (c *createParams) WithConcurrency(concurrent bool) CreateParameters {
	c.concurrency = concurrent
	return c
}

func (c *createParams) GetConcurrency() bool {
	return c.concurrency
}

func (c *createParams) WithCompression(compression Compression) CreateParameters {
	c.compression = compression
	return c
}

func (c *createParams) GetCompression() Compression {
	return c.compression
}

func (c *createParams) WithHashing(hashing Hashing) CreateParameters {
	c.hashing = hashing
	return c
}

func (c *createParams) GetHashing() Hashing {
	return c.hashing
}

func NewCreateParameters() CreateParameters {
	return &createParams{
		capacity:    5_000_000,
		totalFileMb: 16_000_000,
		fileMode:    0644,
		concurrency: false,
		compression: COMPRESSION_NONE,
		hashing:     XX128_HASHING,
	}
}

type compactParams struct {
	minimumCapacity bool
	minimumFileSize bool
}

func (c *compactParams) WithMinimumCapacity(b bool) CompactParameters {
	c.minimumCapacity = b
	return c
}
func (c *compactParams) WithMinimumFileSize(b bool) CompactParameters {
	c.minimumFileSize = b
	return c
}

func NewCompactParameters() CompactParameters {
	return &compactParams{
		minimumCapacity: false,
		minimumFileSize: false,
	}
}
