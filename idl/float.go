package idl

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
	"reflect"

	"github.com/aviate-labs/leb128"
)

type Float struct {
	size uint8
	primType
}

func Float32() *Float {
	return &Float{
		size: 4,
	}
}

func Float64() *Float {
	return &Float{
		size: 8,
	}
}

func (f *Float) Decode(r *bytes.Reader) (interface{}, error) {
	switch f.size {
	case 4:
		v := make([]byte, f.size)
		n, err := r.Read(v)
		if err != nil {
			return nil, err
		}
		if uint8(n) != f.size {
			return nil, fmt.Errorf("float32: too short")
		}
		return math.Float32frombits(
			binary.LittleEndian.Uint32(v),
		), nil
	case 8:
		v := make([]byte, f.size)
		n, err := r.Read(v)
		if err != nil {
			return nil, err
		}
		if uint8(n) != f.size {
			return nil, fmt.Errorf("float64: too short")
		}
		return math.Float64frombits(
			binary.LittleEndian.Uint64(v),
		), nil
	default:
		return nil, fmt.Errorf("invalid float type with size %d", f.size)
	}
}

func (f Float) EncodeType(_ *TypeDefinitionTable) ([]byte, error) {
	floatXType := new(big.Int).Set(big.NewInt(floatXType))
	if f.size == 8 {
		floatXType.Add(floatXType, big.NewInt(-1))
	}
	return leb128.EncodeSigned(floatXType)
}

func (f Float) EncodeValue(v interface{}) ([]byte, error) {
	return encode(reflect.ValueOf(v), func(k reflect.Kind, v reflect.Value) ([]byte, error) {
		switch k {
		case reflect.Float32:
			bs := make([]byte, f.size)
			binary.LittleEndian.PutUint32(bs, math.Float32bits(float32(v.Float())))
			return bs, nil
		case reflect.Float64:
			if f.size == 4 {
				return nil, fmt.Errorf("can not encode float64 into float32")
			}
			bs := make([]byte, f.size)
			binary.LittleEndian.PutUint64(bs, math.Float64bits(float64(v.Float())))
			return bs, nil
		default:
			return nil, fmt.Errorf("invalid float value: %s", v.Kind())
		}
	})
}

func (f Float) String() string {
	return fmt.Sprintf("float%d", f.size*8)
}
