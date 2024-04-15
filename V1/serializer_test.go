package V1

import (
	"github.com/zeebo/assert"
	"math"
	"reflect"
	"testing"
)

func TestSerializer(t *testing.T) {
	s := BasicTypesSerializer{}
	buf := make([]byte, 64)

	// Testing boolean
	_, _ = s.Write(true, buf)
	v, _ := s.Read(buf, reflect.TypeFor[bool]())
	assert.Equal(t, true, v)

	// Testing int8
	_, _ = s.Write(int8(-120), buf)
	v, _ = s.Read(buf, reflect.TypeFor[int8]())
	assert.Equal(t, int8(-120), v)

	// Testing uint8
	_, _ = s.Write(uint8(250), buf)
	v, _ = s.Read(buf, reflect.TypeFor[uint8]())
	assert.Equal(t, uint8(250), v)

	// Testing int16
	_, _ = s.Write(int16(-32767), buf)
	v, _ = s.Read(buf, reflect.TypeFor[int16]())
	assert.Equal(t, int16(-32767), v)

	// Testing uint16
	_, _ = s.Write(uint16(65535), buf)
	v, _ = s.Read(buf, reflect.TypeFor[uint32]())
	assert.Equal(t, uint16(65535), v)

	// Testing int32
	_, _ = s.Write(int32(-2147483647), buf)
	v, _ = s.Read(buf, reflect.TypeFor[int32]())
	assert.Equal(t, int32(-2147483647), v)

	// Testing uint32
	_, _ = s.Write(uint32(4294967295), buf)
	v, _ = s.Read(buf, reflect.TypeFor[uint32]())
	assert.Equal(t, uint32(4294967295), v)

	// Testing int64
	_, _ = s.Write(int64(-math.MaxInt64+1), buf)
	v, _ = s.Read(buf, reflect.TypeFor[int64]())
	assert.Equal(t, int64(-math.MaxInt64+1), v)

	// Testing uint64
	_, _ = s.Write(uint64(math.MaxUint64), buf)
	v, _ = s.Read(buf, reflect.TypeFor[uint64]())
	assert.Equal(t, uint64(math.MaxUint64), v)

	// Testing float32
	_, _ = s.Write(float32(123.456), buf)
	v, _ = s.Read(buf, reflect.TypeFor[float32]())
	assert.Equal(t, float32(123.456), v)

	// Testing float64
	_, _ = s.Write(float64(789.0123456789), buf)
	v, _ = s.Read(buf, reflect.TypeFor[float64]())
	assert.Equal(t, float64(789.0123456789), v)

	// Testing string
	testString := "Hello, World!"
	_, _ = s.Write(testString, buf)
	v, _ = s.Read(buf, reflect.TypeFor[string]())
	assert.Equal(t, testString, v)
}

func TestSerializerUnboundedShouldFail(t *testing.T) {
	s := BasicTypesSerializer{}
	buf := make([]byte, 64)

	_, err := s.Write(12345, buf)
	assert.Error(t, err)

	_, err = s.Write(uintptr(123), buf)
	assert.Error(t, err)
}

func TestSerializerArray(t *testing.T) {
	s := BasicTypesSerializer{}
	buf := make([]byte, 64)

	_, _ = s.Write([]int16{1, 2, 3, 4}, buf)
	v, _ := s.Read(buf, reflect.TypeFor[[]int16]())
	assert.DeepEqual(t, []int16{1, 2, 3, 4}, v)

	_, err := s.Write([]string{"hello", "world", "!"}, buf)
	assert.Error(t, err) // some values are not of fix value

	_, _ = s.Write([]float64{math.E, math.Phi, math.Pi}, buf)
	v, _ = s.Read(buf, reflect.TypeFor[[]float64]())
	assert.DeepEqual(t, []float64{math.E, math.Phi, math.Pi}, v)

	_, _ = s.Write([]byte{1, 2, 3, 9, 8, 7, 6}, buf)
	v, _ = s.Read(buf, reflect.TypeFor[[]byte]())
	assert.DeepEqual(t, []byte{1, 2, 3, 9, 8, 7, 6}, v)
}
