package V1

import (
	"encoding/binary"
	"github.com/go-faster/city"
	"github.com/zeebo/xxh3"
)

type BasicTypesHasher struct{}

func (n BasicTypesHasher) bytesOfThis(v interface{}) []byte {
	var b []byte
	switch t := v.(type) {
	case uint64:
		b = make([]byte, 8)
		binary.LittleEndian.PutUint64(b[0:], t)
	case int64:
		b = make([]byte, 8)
		binary.LittleEndian.PutUint64(b[0:], uint64(t))
	case uint32:
		b = make([]byte, 4)
		binary.LittleEndian.PutUint32(b[0:], t)
	case int32:
		b = make([]byte, 4)
		binary.LittleEndian.PutUint32(b[0:], uint32(t))
	case uint16:
		b = make([]byte, 2)
		binary.LittleEndian.PutUint16(b[0:], t)
	case int16:
		b = make([]byte, 2)
		binary.LittleEndian.PutUint16(b[0:], uint16(t))
	case uint8:
		b = make([]byte, 1)
		b[0] = t
	case bool:
		b = make([]byte, 1)
		if v.(bool) {
			b[0] = 1
		}
	case string:
		b = []byte(t)
	}
	return b
}

func (n BasicTypesHasher) XX128(v interface{}) (low, high uint64, ok bool) {
	b := n.bytesOfThis(v)
	if b != nil {
		h := xxh3.Hash128(b)
		return h.Lo, h.Hi, true
	}
	return 0, 0, false
}

func (n BasicTypesHasher) CityHash128(v interface{}) (low, high uint64, ok bool) {
	b := n.bytesOfThis(v)
	if b != nil {
		h := city.CH128(b)
		return h.Low, h.High, true
	}
	return 0, 0, false
}
