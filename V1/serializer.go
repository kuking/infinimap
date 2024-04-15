package V1

import (
	"bytes"
	"encoding/binary"
	"math"
	"reflect"
)

type BasicTypesSerializer struct{}

// the explicit converts at the beginning are for speed purposes, you don't want to create multiple objects just for parsing an uint64, if there is no specific
// function to write/read, it will use the  default which is binary.write/read, slower due to readers/writers/reflection, etc...
// ... but might be the best of worst worlds if using complex types

func (_ BasicTypesSerializer) Write(v interface{}, dst []byte) (int, error) {
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
	default: // slow
		bw := bytes.NewBuffer(dst[:0])
		// writes the slice length, if it is a slice
		val := reflect.ValueOf(v)
		if val.Kind() == reflect.Slice {
			if err := binary.Write(bw, binary.LittleEndian, int32(val.Len())); err != nil {
				return 0, err
			}
		}
		err := binary.Write(bw, binary.LittleEndian, v)
		return bw.Len(), err
	}
}

func (_ BasicTypesSerializer) Read(dst []byte, k reflect.Type) (interface{}, error) {
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
	default: // slow
		r := bytes.NewReader(dst)
		var dest reflect.Value
		var length int32
		if k.Kind() == reflect.Slice {
			// First read the length if it's a slice
			if err := binary.Read(r, binary.LittleEndian, &length); err != nil {
				return nil, err
			}
			dest = reflect.MakeSlice(k, int(length), int(length))
		} else {
			dest = reflect.New(k).Elem()
		}
		err := binary.Read(r, binary.LittleEndian, dest.Interface())
		return dest.Interface(), err
	}
}
