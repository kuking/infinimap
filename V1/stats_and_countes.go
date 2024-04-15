package V1

import (
	"encoding/binary"
)

func (m *im[K, V]) getFreeNextByte() uint64 {
	return binary.LittleEndian.Uint64(m.mem[OFS_NEXT_FREE_BYTE:])
}

func (m *im[K, V]) incFreeNextByte(size uint64) uint64 {
	nextFreeByte := m.getFreeNextByte() + size
	binary.LittleEndian.PutUint64(m.mem[OFS_NEXT_FREE_BYTE:], nextFreeByte)
	return nextFreeByte
}

func (m *im[K, V]) writeFreeNextByte(next uint64) {
	binary.LittleEndian.PutUint64(m.mem[OFS_NEXT_FREE_BYTE:], next)
}

func (m *im[K, V]) CountU64() uint64 {
	return binary.LittleEndian.Uint64(m.mem[OFS_COUNT:])
}

func (m *im[K, V]) incCount() uint64 {
	newCount := m.CountU64() + 1
	binary.LittleEndian.PutUint64(m.mem[OFS_COUNT:], newCount)
	return newCount
}

func (m *im[K, V]) decCount() uint64 {
	newCount := m.CountU64() - 1
	binary.LittleEndian.PutUint64(m.mem[OFS_COUNT:], newCount)
	return newCount
}

func (m *im[K, V]) StatsInserts() uint64 {
	return binary.LittleEndian.Uint64(m.mem[OFS_INSERTS:])
}

func (m *im[K, V]) incInserts() uint64 {
	newCount := m.StatsInserts() + 1
	binary.LittleEndian.PutUint64(m.mem[OFS_INSERTS:], newCount)
	return newCount
}

func (m *im[K, V]) StatsDeletes() uint64 {
	return binary.LittleEndian.Uint64(m.mem[OFS_DELETES:])
}

func (m *im[K, V]) incDeletes() uint64 {
	newCount := m.StatsDeletes() + 1
	binary.LittleEndian.PutUint64(m.mem[OFS_DELETES:], newCount)
	return newCount
}

func (m *im[K, V]) StatsUpdates() uint64 {
	return binary.LittleEndian.Uint64(m.mem[OFS_UPDATES:])
}

func (m *im[K, V]) incUpdates() uint64 {
	newCount := m.StatsUpdates() + 1
	binary.LittleEndian.PutUint64(m.mem[OFS_UPDATES:], newCount)
	return newCount
}

func (m *im[K, V]) ClogRatio() uint8 {
	return m.mem[OFS_CLOG_RATIO]
}

func (m *im[K, V]) resetClogRatio() {
	m.mem[OFS_CLOG_RATIO] = 0
}

func (m *im[K, V]) updateClogRatio(visitedBuckets uint32) {
	ratio := uint8(int64(visitedBuckets*0xff) / int64(m.buckets))
	if m.ClogRatio() < ratio {
		m.mem[OFS_CLOG_RATIO] = ratio
	}
}

func (m *im[K, V]) BytesAllocated() uint64 {
	return m.getFreeNextByte()
}

func (m *im[K, V]) BytesAvailable() uint64 {
	if fs, err := m.file.Stat(); err == nil {
		return uint64(fs.Size()) - m.BytesAllocated()
	}
	return 0
}

func (m *im[K, V]) BytesInUse() uint64 {
	return m.BytesAllocated() - m.BytesReclaimable()
}

func (m *im[K, V]) BytesReclaimable() uint64 {
	deleted := uint64(0)
	ofs := uint64(OFS_FIRST_RECORD) + uint64(BUCKET_SIZE)*uint64(m.buckets)
	lastOfs := m.getFreeNextByte()
	for {
		if ofs >= lastOfs {
			return deleted
		}
		newOfs := m.nextRecordOfs(ofs)
		if !m.isPopulatedRecord(ofs) {
			deleted += newOfs - ofs
		}
		ofs = newOfs
	}
}
