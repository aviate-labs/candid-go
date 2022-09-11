package idl

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/aviate-labs/candid-go/typ"
	"github.com/aviate-labs/leb128"
)

func encodeInt16(v interface{}) (int16, error) {
	if v, ok := v.(int16); ok {
		return v, nil
	}
	v_, err := encodeInt8(v)
	return int16(v_), err
}

func encodeInt32(v interface{}) (int32, error) {
	if v, ok := v.(int32); ok {
		return v, nil
	}
	v_, err := encodeInt16(v)
	return int32(v_), err
}

func encodeInt64(v interface{}) (int64, error) {
	if v, ok := v.(int); ok {
		return int64(v), nil
	}
	if v, ok := v.(int64); ok {
		return v, nil
	}
	v_, err := encodeInt16(v)
	return int64(v_), err
}

func encodeInt8(v interface{}) (int8, error) {
	if v, ok := v.(int8); ok {
		return v, nil
	}
	return 0, fmt.Errorf("invalid value: %v", v)
}

type IntType struct {
	size uint8
	primType
}

func Int16Type() *IntType {
	return &IntType{
		size: 2,
	}
}

func Int32Type() *IntType {
	return &IntType{
		size: 4,
	}
}

func Int64Type() *IntType {
	return &IntType{
		size: 8,
	}
}

func Int8Type() *IntType {
	return &IntType{
		size: 1,
	}
}

func (n IntType) Base() uint {
	return uint(n.size)
}

func (n *IntType) Decode(r *bytes.Reader) (interface{}, error) {
	switch n.size {
	case 0:
		bi, err := leb128.DecodeSigned(r)
		if err != nil {
			return nil, err
		}
		return typ.NewBigInt(bi), nil
	case 8:
		v := make([]byte, 8)
		n, err := r.Read(v)
		if err != nil {
			return nil, err
		}
		if n != 8 {
			return nil, fmt.Errorf("int64: too short")
		}
		bi, err := readInt(new(big.Int).SetUint64(binary.LittleEndian.Uint64(v)), 8)
		if err != nil {
			return nil, err
		}
		return bi.Int64(), nil
	case 4:
		v := make([]byte, 4)
		n, err := r.Read(v)
		if err != nil {
			return nil, err
		}
		if n != 4 {
			return nil, fmt.Errorf("int32: too short")
		}
		bi, err := readInt(new(big.Int).SetUint64(uint64(binary.LittleEndian.Uint32(v))), 8)
		if err != nil {
			return nil, err
		}
		return int32(bi.Int64()), nil
	case 2:
		v := make([]byte, 2)
		n, err := r.Read(v)
		if err != nil {
			return nil, err
		}
		if n != 2 {
			return nil, fmt.Errorf("int16: too short")
		}
		bi, err := readInt(new(big.Int).SetUint64(uint64(binary.LittleEndian.Uint16(v))), 8)
		if err != nil {
			return nil, err
		}
		return int16(bi.Int64()), nil
	case 1:
		v, err := r.ReadByte()
		if err != nil {
			return nil, err
		}
		bi, err := readInt(new(big.Int).SetUint64(uint64(v)), 8)
		if err != nil {
			return nil, err
		}
		return int8(bi.Int64()), nil
	default:
		return nil, fmt.Errorf("invalid int type with size %d", n.size)
	}
}

func (n IntType) EncodeType(_ *TypeDefinitionTable) ([]byte, error) {
	if n.size == 0 {
		return leb128.EncodeSigned(big.NewInt(intType))
	}
	intXType := new(big.Int).Set(big.NewInt(intXType))
	intXType = intXType.Add(
		intXType,
		big.NewInt(3-int64(log2(n.size*8))),
	)
	return leb128.EncodeSigned(intXType)
}

func (n IntType) EncodeValue(v interface{}) ([]byte, error) {
	switch n.size {
	case 0:
		v, ok := v.(*big.Int)
		if !ok {
			return nil, fmt.Errorf("invalid value: %v", v)
		}
		return leb128.EncodeSigned(v)
	case 8:
		v, err := encodeInt64(v)
		if err != nil {
			return nil, err
		}
		return writeInt(big.NewInt(v), 8), nil
	case 4:
		v, err := encodeInt32(v)
		if err != nil {
			return nil, err
		}
		return writeInt(big.NewInt(int64(v)), 4), nil
	case 2:
		v, err := encodeInt16(v)
		if err != nil {
			return nil, err
		}
		return writeInt(big.NewInt(int64(v)), 2), nil
	case 1:
		v, err := encodeInt8(v)
		if err != nil {
			return nil, err
		}
		return writeInt(big.NewInt(int64(v)), 1), nil
	default:
		return nil, fmt.Errorf("invalid argument: %v", v)
	}
}

func (n IntType) String() string {
	if n.size == 0 {
		return "int"
	}
	return fmt.Sprintf("int%d", n.size*8)
}
