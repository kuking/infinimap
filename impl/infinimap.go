package impl

import (
	"errors"
	"fmt"
	"github.com/edsrzf/mmap-go"
	"github.com/kuking/infinimap"
	"os"
)

type im[K comparable, V any] struct {
	compression infinimap.Compression
	hashing     infinimap.Hashing
	hasher      infinimap.Hasher
	seraliser   infinimap.Serializer
	buckets     uint32
	file        *os.File
	mem         mmap.MMap
}

func (m *im[K, V]) Put(k K, v V) (previous V, replace bool, err error) {

	lo, hi, err := m.resolveHash(k)
	if err != nil {
		return v, false, err
	}

	bucket := m.calBucketFromHash(lo, hi)
	startingBucket := bucket
	for {
		blo, bhi, bofs := m.readBucket(bucket)
		if blo == 0 && bhi == 0 && bofs == 0 {
			// we found an empty space, it is a vanilla insert
			ofs := m.getFreeNextByte()
			m.writeBucket(bucket, lo, hi, ofs)
			size, err := m.writeRecord(ofs, lo, hi, k, v)
			if err != nil {
				return zero[V](), false, err
			}
			m.incFreeNextByte(size)
			m.incCount()
			return zero[V](), false, nil
		} else if blo == lo && bhi == hi {
			bucketRecordKey, err := m.readRecordKey(bofs)
			if err != nil {
				return zero[V](), false, err
			}
			if bucketRecordKey == k {
				// this is an update
				previous, err := m.readRecordValue(bofs)
				if err != nil {
					return zero[V](), false, err
				}
				m.eraseRecord(bofs)
				ofs := m.getFreeNextByte()
				m.writeBucket(bucket, lo, hi, ofs)
				size, err := m.writeRecord(ofs, lo, hi, k, v)
				if err != nil {
					return zero[V](), false, err
				}
				m.incFreeNextByte(size)
				return previous, true, nil
			}
		}

		bucket = (bucket + 1) % m.buckets
		if startingBucket == bucket {
			return zero[V](), false, errors.New("no empty bucket found, consider increasing the map capacity at creation time or convert it")
		}
		bucket++
	}

}

func (m *im[K, V]) Get(k K) (V, bool) {

	lo, hi, err := m.resolveHash(k)
	if err != nil {
		return zero[V](), false
	}

	bucket := m.calBucketFromHash(lo, hi)

	blo, bhi, bofs := m.readBucket(bucket)
	if blo == lo && bhi == hi && m.isSlotForKey(bofs, lo, hi, k) {
		if v, err := m.readRecordValue(bofs); err == nil {
			return v, true
		}
	}

	return zero[V](), false

}

func (m *im[K, V]) Delete(k K) bool {
	lo, hi, err := m.resolveHash(k)
	if err != nil {
		return false
	}
	bucket := m.calBucketFromHash(lo, hi)
	blo, bhi, bofs := m.readBucket(bucket)
	if blo == lo && bhi == hi && m.isSlotForKey(bofs, lo, hi, k) {
		m.eraseBucket(bucket)
		m.eraseRecord(bofs)
		m.decCount()
	}
	panic(errors.New("not implemented"))
}

func (m *im[K, V]) Count() int {
	return int(m.readCount())
}

func (m *im[K, V]) Keys() <-chan K {
	panic(errors.New("not implemented"))
}

func (m *im[K, V]) Values() <-chan V {
	panic(errors.New("not implemented"))
}

func (m *im[K, V]) Each(f func(K, V) bool) error {
	for bucket := uint32(0); bucket < m.buckets; bucket++ {
		lo, hi, ofs := m.readBucket(bucket)
		if lo != 0 && hi != 0 && ofs != 0 {
			k, err := m.readRecordKey(ofs)
			if err != nil {
				return err
			}
			v, err := m.readRecordValue(ofs)
			if err != nil {
				return err
			}
			if !f(k, v) {
				return nil // finish if function returns false
			}
		}
	}
	return nil
}

func (m *im[K, V]) Compact() error {
	panic(errors.New("not implemented"))
}

func (m *im[K, V]) Sync() error {
	panic(errors.New("not implemented"))
}

func (m *im[K, V]) Close() error {
	panic(errors.New("not implemented"))
}

// ---------------------------------------------------------------------------------------------------------------------------------------------------------

func (m *im[K, V]) calBucketFromHash(lo, hi uint64) uint32 {
	loMod := uint32(lo % uint64(m.buckets))
	hiMod := uint32(hi % uint64(m.buckets))
	return (loMod + hiMod) % m.buckets
}

func (m *im[K, V]) resolveHash(k K) (lo uint64, hi uint64, err error) {
	if m.hashing == infinimap.XX128_HASHING {
		if lo, hi, ok := m.hasher.XX128(k); ok {
			return lo, hi, nil
		}
	} else if m.hashing == infinimap.CITY128 {
		if lo, hi, ok := m.hasher.CityHash128(k); ok {
			return lo, hi, nil
		}
	}
	return 0, 0, errors.New(fmt.Sprintf("hasher type unimplemented: %v", m.hashing))
}

//func (m *im[K, V]) findEmptyBucket(bucket uint32) (emptyBucket uint32, err error) {
//	emptyBucket = bucket
//	lo, hi, ofs := m.readBucket(emptyBucket)
//	for lo != 0 && hi != 0 && ofs != 0 {
//		emptyBucket = (emptyBucket + 1) % m.buckets
//		if emptyBucket == bucket {
//			return emptyBucket, errors.New("no empty bucket found, consider increasing the map capacity at creation time or convert it")
//		}
//		lo, hi, ofs = m.readBucket(emptyBucket)
//	}
//	return
//}
