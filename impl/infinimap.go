package impl

import (
	"errors"
	"fmt"
	"github.com/edsrzf/mmap-go"
	"github.com/kuking/infinimap"
	"log"
	"math"
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

	firstTombstone := uint32(math.MaxUint32)
	bucket := m.calBucketFromHash(lo, hi)
	startingBucket := bucket
	visitedBuckets := uint32(0)
	defer func() {
		m.updateClogRatio(visitedBuckets)
	}()
	for {
		blo, bhi, bofs := m.readBucket(bucket)
		if m.isNeverUsedBucket(blo, bhi, bofs) || (startingBucket == bucket && firstTombstone != math.MaxUint32) {
			// it is either a never used bucket, of we went around the whole bucket-space, nothing is clean but only tombstones
			// we found a never used bucket; so definitely not an update
			var insertBucket uint32
			if firstTombstone != math.MaxUint32 {
				// insert into the first found tombstone
				insertBucket = firstTombstone
			} else {
				insertBucket = bucket
			}
			ofs := m.getFreeNextByte()
			m.writeBucket(insertBucket, lo, hi, ofs)
			size, err := m.writeRecord(ofs, lo, hi, k, v)
			if err != nil {
				return zero[V](), false, err
			}
			m.incFreeNextByte(size)
			m.incCount()
			m.incInserts()
			return zero[V](), false, nil
		} else if m.isTombstoneBucket(blo, bhi, bofs) {
			if firstTombstone == math.MaxUint32 {
				firstTombstone = bucket
			}
		} else if blo == lo && bhi == hi {
			// it will be an update
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
				m.incUpdates()
				return previous, true, nil
			}
		}
		visitedBuckets++
		bucket = (bucket + 1) % m.buckets
		if startingBucket == bucket {
			if firstTombstone == math.MaxUint32 {
				return zero[V](), false, errors.New("no empty bucket found, consider increasing the map capacity at creation time or convert it")
			}
		}
	}
}

func (m *im[K, V]) Get(k K) (v V, found bool) {

	lo, hi, err := m.resolveHash(k)
	if err != nil {
		return zero[V](), false
	}
	bucket := m.calBucketFromHash(lo, hi)
	startingBucket := bucket
	for {
		blo, bhi, bofs := m.readBucket(bucket)
		if m.isNeverUsedBucket(blo, bhi, bofs) {
			return zero[V](), false
		} else if m.isTombstoneBucket(blo, bhi, bofs) {
			// continue
		} else if blo == lo && bhi == hi && m.isRecordForKey(bofs, lo, hi, k) {
			if v, err := m.readRecordValue(bofs); err == nil {
				return v, true
			}
		}
		bucket = (bucket + 1) % m.buckets
		if startingBucket == bucket {
			return zero[V](), false // not found
		}
	}
}

func (m *im[K, V]) Delete(k K) (deleted bool) {
	lo, hi, err := m.resolveHash(k)
	if err != nil {
		return false
	}
	bucket := m.calBucketFromHash(lo, hi)
	startingBucket := bucket
	visitedBuckets := uint32(0)
	defer func() {
		m.updateClogRatio(visitedBuckets)
	}()
	for {
		blo, bhi, bofs := m.readBucket(bucket)
		if m.isNeverUsedBucket(blo, bhi, bofs) {
			// not found
			return false
		} else if m.isTombstoneBucket(blo, bhi, bofs) {
			// continue
		} else if blo == lo && bhi == hi && m.isRecordForKey(bofs, lo, hi, k) {
			// found
			m.eraseBucket(bucket)
			m.eraseRecord(bofs)
			m.decCount()
			m.incDeletes()
			return true
		}
		visitedBuckets++
		bucket = (bucket + 1) % m.buckets
		if startingBucket == bucket {
			return false // not found after iterating over all buckets
		}
	}
}

func (m *im[K, V]) Count() int {
	return int(m.CountU64())
}

func (m *im[K, V]) Keys() <-chan K {
	ch := make(chan K, 10)
	go func() {
		defer close(ch)
		for bucket := uint32(0); bucket < m.buckets; bucket++ {
			lo, hi, ofs := m.readBucket(bucket)
			if m.isUsedBucket(lo, hi, ofs) {
				key, err := m.readRecordKey(ofs)
				if err != nil {
					log.Print(err)
				}
				ch <- key
			}
		}
	}()
	return ch
}

func (m *im[K, V]) Values() <-chan V {
	ch := make(chan V, 10)
	go func() {
		defer close(ch)
		for bucket := uint32(0); bucket < m.buckets; bucket++ {
			lo, hi, ofs := m.readBucket(bucket)
			if m.isUsedBucket(lo, hi, ofs) {
				value, err := m.readRecordValue(ofs)
				if err != nil {
					log.Print(err)
				}
				ch <- value
			}
		}
	}()
	return ch
}

func (m *im[K, V]) Each(f func(K, V) (cont bool)) error {
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
