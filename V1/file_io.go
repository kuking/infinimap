package V1

import (
	"encoding/binary"
	"errors"
	"github.com/edsrzf/mmap-go"
	"os"
	"reflect"
)

const (
	OFS_FILE_VERSION     int = 0
	OFS_HASHING_ALGO     int = OFS_FILE_VERSION + 2
	OFS_COMPRESSION_ALGO int = OFS_HASHING_ALGO + 1
	OFS_BUCKETS          int = OFS_COMPRESSION_ALGO + 1
	OFS_NEXT_FREE_BYTE       = OFS_BUCKETS + 4

	OFS_INSERTS      = OFS_NEXT_FREE_BYTE + 8
	OFS_UPDATES      = OFS_INSERTS + 8
	OFS_DELETES      = OFS_UPDATES + 8
	OFS_COUNT        = OFS_DELETES + 8
	OFS_CLOG_RATIO   = OFS_COUNT + 8
	OFS_FIRST_BUCKET = OFS_CLOG_RATIO + 1 // what if 1K instead of just after last field?

	BUCKET_SIZE = (128 + 64) / 8

	OFS_FIRST_RECORD = OFS_FIRST_BUCKET // + m.buckets() * BUCKET_SIZE

	RECORD_LO_HASH    = 0
	RECORD_HI_HASH    = RECORD_LO_HASH + 8
	RECORD_KEY_SIZE   = RECORD_HI_HASH + 8
	RECORD_VALUE_SIZE = RECORD_KEY_SIZE + 4
	RECORD_KEY        = RECORD_VALUE_SIZE + 4
	RECORD_VALUE      = RECORD_KEY // needs the keysized to be added

	THOMBSTONE_LO  = 1
	THOMBSTONE_HI  = 2
	THOMBSTONE_OFS = 3
)

func Create[K comparable, V any](path string, cfg CreateParameters) (InfiniMap[K, V], error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, cfg.GetFileMode())
	if err != nil {
		return nil, err
	}

	if err = file.Truncate(int64(cfg.GetFileSizeMB()) * 1024 * 1024); err != nil {
		return nil, err
	}

	mem, err := mmap.Map(file, mmap.RDWR, 0)
	if err != nil {
		return nil, err
	}

	im := &im[K, V]{
		compression: cfg.GetCompression(),
		hashing:     cfg.GetHashing(),
		hasher:      BasicTypesHasher{},
		seraliser:   BasicTypesSerializer{},
		buckets:     uint32(cfg.GetCapacity()) * 2, // twice the amount of buckets as expected capacity
		file:        file,
		mem:         mem,
	}

	binary.LittleEndian.PutUint16(mem[OFS_FILE_VERSION:], FILE_VERSION_1)
	mem[OFS_HASHING_ALGO] = byte(im.hashing)
	mem[OFS_COMPRESSION_ALGO] = byte(im.compression)
	binary.LittleEndian.PutUint32(mem[OFS_BUCKETS:], im.buckets)
	binary.LittleEndian.PutUint64(mem[OFS_NEXT_FREE_BYTE:], uint64(OFS_FIRST_RECORD)+uint64(BUCKET_SIZE)*uint64(im.buckets))
	binary.LittleEndian.PutUint64(mem[OFS_INSERTS:], 0)
	binary.LittleEndian.PutUint64(mem[OFS_DELETES:], 0)
	binary.LittleEndian.PutUint64(mem[OFS_COUNT:], 0)

	assertValidIm(im)

	return im, nil
}

func Open[K comparable, V any](path string) (InfiniMap[K, V], error) {

	file, err := os.OpenFile(path, os.O_RDWR, 0)
	if err != nil {
		return nil, err
	}

	mem, err := mmap.Map(file, mmap.RDWR, 0)
	if err != nil {
		return nil, err
	}

	im := &im[K, V]{
		compression: Compression(mem[OFS_COMPRESSION_ALGO]),
		hashing:     Hashing(mem[OFS_HASHING_ALGO]),
		hasher:      BasicTypesHasher{},
		seraliser:   BasicTypesSerializer{},
		buckets:     binary.LittleEndian.Uint32(mem[OFS_BUCKETS:]),
		file:        file,
		mem:         mem,
	}

	assertValidIm(im)

	return im, nil
}

func OpenOrCreate[K comparable, V any](path string, cfg CreateParameters) (InfiniMap[K, V], error) {
	if fi, err := os.Stat(path); err != nil && !fi.IsDir() {
		return Open[K, V](path)
	} else {
		return Create[K, V](path, cfg)
	}
}

func (m *im[K, V]) Shrink() error {
	panic(errors.New("not implemented"))
}

func (m *im[K, V]) Expand() error {
	panic(errors.New("not implemented"))
}

func (m *im[K, V]) Sync() error {
	if err := m.mem.Flush(); err != nil {
		return err
	}
	if err := m.file.Sync(); err != nil {
		return err
	}
	return nil
}

func (m *im[K, V]) Close() error {
	if err := m.Sync(); err != nil {
		return err
	}
	if err := m.mem.Unmap(); err != nil {
		return err
	}
	if err := m.file.Close(); err != nil {
		return err
	}
	return nil
}

func assertValidIm[K comparable, V any](i *im[K, V]) {
	if i.compression != COMPRESSION_NONE {
		panic(errors.New("compression not yet implemented"))
	}
}

func (m *im[K, V]) readBucket(bucket uint32) (lo, hi uint64, offset uint64) {
	bucketOfs := uintptr(OFS_FIRST_BUCKET) + uintptr(BUCKET_SIZE*bucket)
	lo = binary.LittleEndian.Uint64(m.mem[bucketOfs:])
	hi = binary.LittleEndian.Uint64(m.mem[bucketOfs+8:])
	offset = binary.LittleEndian.Uint64(m.mem[bucketOfs+16:])
	return
}

func (m *im[K, V]) writeBucket(bucket uint32, lo, hi, offset uint64) {
	bucketOfs := uintptr(OFS_FIRST_BUCKET) + uintptr(BUCKET_SIZE*bucket)
	binary.LittleEndian.PutUint64(m.mem[bucketOfs:], lo)
	binary.LittleEndian.PutUint64(m.mem[bucketOfs+8:], hi)
	binary.LittleEndian.PutUint64(m.mem[bucketOfs+16:], offset)
}

func (m *im[K, V]) eraseBucket(bucket uint32) {
	m.writeBucket(bucket, THOMBSTONE_LO, THOMBSTONE_HI, THOMBSTONE_OFS)
}

func (m *im[K, V]) resetBucket(bucket uint32) {
	m.writeBucket(bucket, 0, 0, 0)
}

func (m *im[K, V]) isNeverUsedBucket(lo, hi, ofs uint64) bool {
	return lo == 0 && hi == 0 && ofs == 0
}

func (m *im[K, V]) isTombstoneBucket(lo, hi, ofs uint64) bool {
	return lo == THOMBSTONE_LO && hi == THOMBSTONE_HI && ofs == THOMBSTONE_OFS
}

func (m *im[K, V]) isUsedBucket(lo, hi, ofs uint64) bool {
	return !m.isNeverUsedBucket(lo, hi, ofs) && !m.isTombstoneBucket(lo, hi, ofs)
}

func (m *im[K, V]) writeRecord(ofs uint64, lo uint64, hi uint64, k K, v V) (size uint64, err error) {
	keyLength, err := m.seraliser.Write(k, m.mem[ofs+RECORD_KEY:])
	if err != nil {
		return 0, err
	}
	ValueLength, err := m.seraliser.Write(v, m.mem[ofs+RECORD_VALUE+uint64(keyLength):])
	if err != nil {
		return 0, err
	}
	binary.LittleEndian.PutUint64(m.mem[ofs+RECORD_LO_HASH:], lo)
	binary.LittleEndian.PutUint64(m.mem[ofs+RECORD_HI_HASH:], hi)
	binary.LittleEndian.PutUint32(m.mem[ofs+RECORD_KEY_SIZE:], uint32(keyLength))
	binary.LittleEndian.PutUint32(m.mem[ofs+RECORD_VALUE_SIZE:], uint32(ValueLength))
	return RECORD_VALUE + uint64(keyLength) + uint64(ValueLength), nil
}

func (m *im[K, V]) isRecordForKey(ofs uint64, lo uint64, hi uint64, k K) bool {
	if binary.LittleEndian.Uint64(m.mem[ofs+RECORD_LO_HASH:]) != lo || binary.LittleEndian.Uint64(m.mem[ofs+RECORD_HI_HASH:]) != hi {
		return false
	}
	kv, err := m.seraliser.Read(m.mem[ofs+RECORD_KEY:], reflect.TypeFor[K]())
	if err == nil {
		return kv.(K) == k
	}
	return false
}

func (m *im[K, V]) readRecordKey(ofs uint64) (K, error) {
	kv, err := m.seraliser.Read(m.mem[ofs+RECORD_KEY:], reflect.TypeFor[K]())
	if err == nil {
		return kv.(K), nil
	}
	return zero[K](), err
}

func (m *im[K, V]) readRecordValue(ofs uint64) (V, error) {
	keySize := binary.LittleEndian.Uint32(m.mem[ofs+16:])
	vv, err := m.seraliser.Read(m.mem[ofs+RECORD_VALUE+uint64(keySize):], reflect.TypeFor[V]())
	if err == nil {
		return vv.(V), nil
	}
	return zero[V](), err
}

func (m *im[K, V]) eraseRecord(ofs uint64) {
	// we leave keySize and valueSize there so we can calculate next record
	keySize := binary.LittleEndian.Uint32(m.mem[ofs+RECORD_KEY_SIZE:])
	valueSize := binary.LittleEndian.Uint32(m.mem[ofs+RECORD_VALUE_SIZE:])
	for o := ofs + RECORD_KEY; o < ofs+RECORD_VALUE+uint64(keySize)+uint64(valueSize); o++ {
		m.mem[o] = 0 //XXX this needs to be efficient
	}
	binary.LittleEndian.PutUint64(m.mem[ofs+RECORD_LO_HASH:], THOMBSTONE_LO)
	binary.LittleEndian.PutUint64(m.mem[ofs+RECORD_HI_HASH:], THOMBSTONE_HI)
}

func (m *im[K, V]) isPopulatedRecord(ofs uint64) bool {
	lo := binary.LittleEndian.Uint64(m.mem[ofs+RECORD_LO_HASH:])
	hi := binary.LittleEndian.Uint64(m.mem[ofs+RECORD_HI_HASH:])
	keySize := binary.LittleEndian.Uint32(m.mem[ofs+RECORD_KEY_SIZE:])
	valueSize := binary.LittleEndian.Uint32(m.mem[ofs+RECORD_VALUE_SIZE:])
	return keySize > 0 && valueSize > 0 && lo != THOMBSTONE_LO && hi != THOMBSTONE_HI
}

func (m *im[K, V]) nextRecordOfs(ofs uint64) uint64 {
	keySize := binary.LittleEndian.Uint32(m.mem[ofs+RECORD_KEY_SIZE:])
	valueSize := binary.LittleEndian.Uint32(m.mem[ofs+RECORD_VALUE_SIZE:])
	if keySize == 0 || valueSize == 0 {
		return ofs
	}
	return ofs + RECORD_VALUE + uint64(keySize) + uint64(valueSize)
}