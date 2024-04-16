package V1

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"math"
	"reflect"
)

type FastAndSlowSerializer struct{}

// the explicit converts at the beginning are for speed purposes, you don't want to create multiple objects just for parsing an uint64, if there is no specific
// function to write/read, it will default to binary.gob, which is good, but slower.

func (_ FastAndSlowSerializer) Write(v interface{}, dst []byte) (int, error) {
	switch v.(type) {
	case bool:
		if v.(bool) {
			dst[0] = 1
		} else {
			dst[0] = 0
		}
		return 1, nil
	case int8:
		dst[0] = byte(v.(int8))
		return 1, nil
	case uint8:
		dst[0] = v.(uint8)
		return 1, nil
	case int16:
		binary.LittleEndian.PutUint16(dst, uint16(v.(int16)))
		return 2, nil
	case uint16:
		binary.LittleEndian.PutUint16(dst, v.(uint16))
		return 2, nil
	case int32:
		binary.LittleEndian.PutUint32(dst, uint32(v.(int32)))
		return 4, nil
	case uint32:
		binary.LittleEndian.PutUint32(dst, v.(uint32))
		return 4, nil
	case int64:
		binary.LittleEndian.PutUint64(dst, uint64(v.(int64)))
		return 8, nil
	case uint64:
		binary.LittleEndian.PutUint64(dst, v.(uint64))
		return 8, nil
	case float32:
		binary.LittleEndian.PutUint32(dst, math.Float32bits(v.(float32)))
		return 4, nil
	case float64:
		binary.LittleEndian.PutUint64(dst, math.Float64bits(v.(float64)))
		return 8, nil
	case string:
		bytes := []byte(v.(string))
		size := uint32(len(bytes))
		binary.LittleEndian.PutUint32(dst[:], size)
		return copy(dst[4:], bytes) + 4, nil
	case []byte:
		size := uint32(len(v.([]byte)))
		binary.LittleEndian.PutUint32(dst[:], size)
		return copy(dst[4:], v.([]byte)) + 4, nil
	default: // slow gob
		bw := bytes.NewBuffer(dst[:0])
		encoder := gob.NewEncoder(bw)
		err := encoder.Encode(v)
		return bw.Len(), err
	}
}

func (_ FastAndSlowSerializer) Read(dst []byte, k reflect.Type) (interface{}, error) {
	switch k.Kind() {
	case reflect.Int8:
		return int8(dst[0]), nil
	case reflect.Uint8:
		return dst[0], nil
	case reflect.Bool:
		return dst[0] != 0, nil
	case reflect.Int16:
		return int16(binary.LittleEndian.Uint16(dst)), nil
	case reflect.Uint16:
		return binary.LittleEndian.Uint16(dst), nil
	case reflect.Int32:
		return int32(binary.LittleEndian.Uint32(dst)), nil
	case reflect.Uint32:
		return binary.LittleEndian.Uint32(dst), nil
	case reflect.Float32:
		return math.Float32frombits(binary.LittleEndian.Uint32(dst)), nil
	case reflect.Int64:
		return int64(binary.LittleEndian.Uint64(dst)), nil
	case reflect.Uint64:
		return binary.LittleEndian.Uint64(dst), nil
	case reflect.Float64:
		return math.Float64frombits(binary.LittleEndian.Uint64(dst)), nil
	case reflect.String:
		size := binary.LittleEndian.Uint32(dst)
		return string(dst[4 : size+4]), nil
	default: // slow gob
		v := reflect.New(k)
		r := bytes.NewReader(dst)
		decoder := gob.NewDecoder(r)
		err := decoder.Decode(v.Interface())
		return v.Interface(), err
	}
}
