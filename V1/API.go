package V1

import (
	"os"
	"reflect"
)

type InfiniMap[K comparable, V any] interface {
	/*
		Fundamental API
	*/
	Put(K, V) (previous V, replace bool, err error)
	Get(K) (value V, found bool)
	Delete(K) (deleted bool)
	Count() int
	Keys() <-chan K
	Values() <-chan V
	Each(func(K, V) (cont bool)) error

	/*
		Map Statistics
	*/
	StatsInserts() uint64
	StatsDeletes() uint64
	StatsUpdates() uint64
	CountU64() uint64

	/*
		Index Maintenance
	*/
	ClogRatio() uint8
	Reindex() error

	/*
		Space maintenance
	*/
	//Compact(CompactParameters) error
	Shrink() error
	Expand() error
	BytesAllocated() uint64
	BytesInUse() uint64
	BytesReclaimable() uint64
	BytesAvailable() uint64

	/*
		IO Closing, misc
	*/
	Sync() error
	Close() error
	SetCustomSerializer(Serializer)
}

type Hasher interface {
	XX128(interface{}) (low, high uint64, ok bool)
	CityHash128(interface{}) (low, high uint64, ok bool)
}

type Serializer interface {
	Write(interface{}, []byte) (int, error)
	Read([]byte, reflect.Type) (interface{}, error)
}

type Compression uint8
type Hashing uint8

const (
	COMPRESSION_NONE Compression = 0
	LZ4_COMPRESSION  Compression = 1

	XX128_HASHING Hashing = 1
	CITY128       Hashing = 2

	FILE_VERSION_1 uint16 = 0x1f17
)

type CreateParameters interface {
	WithCapacity(int) CreateParameters
	GetCapacity() int
	WithFileSizeMB(int) CreateParameters
	GetFileSizeMB() int
	WithFileMode(mode os.FileMode) CreateParameters
	GetFileMode() os.FileMode
	WithConcurrency(bool) CreateParameters
	GetConcurrency() bool
	WithCompression(Compression) CreateParameters
	GetCompression() Compression
	WithHashing(Hashing) CreateParameters
	GetHashing() Hashing
}

type CompactParameters interface {
	WithMinimumCapacity(bool) CompactParameters
	WithMinimumFileSize(bool) CompactParameters
}
