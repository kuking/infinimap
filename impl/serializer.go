package impl

import (
	"encoding/binary"
	"fmt"
	"reflect"
)

type BasicTypesSerializer struct{}

func (_ BasicTypesSerializer) Write(v interface{}, dst []byte) (int, error) {
	switch v.(type) {
	case uint64:
		binary.LittleEndian.PutUint64(dst[:], v.(uint64))
		return 8, nil
	case string:
		bytes := []byte(v.(string))
		size := uint32(len(bytes))
		binary.LittleEndian.PutUint32(dst[:], size)
		return copy(dst[4:], bytes) + 4, nil
	default:
		return 0, fmt.Errorf("basic_types_serializer: uknown type %T", v)
	}
}

func (_ BasicTypesSerializer) Read(dst []byte, k reflect.Kind) (interface{}, error) {
	switch k {
	case reflect.Uint64:
		return binary.LittleEndian.Uint64(dst), nil
	case reflect.String:
		size := binary.LittleEndian.Uint32(dst)
		return string(dst[4 : size+4]), nil
	default:
		return 0, fmt.Errorf("basic_types_serializer: uknown kind %v", k)
	}
}
