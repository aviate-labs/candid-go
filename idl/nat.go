package idl

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/aviate-labs/candid-go/typ"
	"github.com/aviate-labs/leb128"
)

func encodeNat16(v interface{}) (uint16, error) {
	if v, ok := v.(uint16); ok {
		return v, nil
	}
	v_, err := encodeNat8(v)
	return uint16(v_), err
}

func encodeNat32(v interface{}) (uint32, error) {
	if v, ok := v.(uint32); ok {
		return v, nil
	}
	v_, err := encodeNat16(v)
	return uint32(v_), err
}

func encodeNat64(v interface{}) (uint64, error) {
	if v, ok := v.(uint); ok {
		return uint64(v), nil
	}
	if v, ok := v.(uint64); ok {
		return v, nil
	}
	v_, err := encodeNat16(v)
	return uint64(v_), err
}

func encodeNat8(v interface{}) (uint8, error) {
	if v, ok := v.(uint8); ok {
		return v, nil
	}
	return 0, fmt.Errorf("invalid value: %v", v)
}

type Nat struct {
	size uint8
	primType
}

func Nat16() *Nat {
	return &Nat{
		size: 2,
	}
}

func Nat32() *Nat {
	return &Nat{
		size: 4,
	}
}

func Nat64() *Nat {
	return &Nat{
		size: 8,
	}
}

func Nat8() *Nat {
	return &Nat{
		size: 1,
	}
}

func (n Nat) Base() uint {
	return uint(n.size)
}

func (n *Nat) Decode(r *bytes.Reader) (interface{}, error) {
	switch n.size {
	case 0:
		bi, err := leb128.DecodeUnsigned(r)
		if err != nil {
			return nil, err
		}
		return typ.NewBigNat(bi), nil
	case 8:
		v := make([]byte, 8)
		n, err := r.Read(v)
		if err != nil {
			return nil, err
		}
		if n != 8 {
			return nil, fmt.Errorf("nat64: too short")
		}
		return binary.LittleEndian.Uint64(v), nil
	case 4:
		v := make([]byte, 4)
		n, err := r.Read(v)
		if err != nil {
			return nil, err
		}
		if n != 4 {
			return nil, fmt.Errorf("nat32: too short")
		}
		return binary.LittleEndian.Uint32(v), nil
	case 2:
		v := make([]byte, 2)
		n, err := r.Read(v)
		if err != nil {
			return nil, err
		}
		if n != 2 {
			return nil, fmt.Errorf("nat16: too short")
		}
		return binary.LittleEndian.Uint16(v), nil
	case 1:
		return r.ReadByte()
	default:
		return nil, fmt.Errorf("invalid int type with size %d", n.size)
	}
}

func (n Nat) EncodeType(_ *TypeDefinitionTable) ([]byte, error) {
	if n.size == 0 {
		return leb128.EncodeSigned(big.NewInt(natType))
	}
	natXType := new(big.Int).Set(big.NewInt(natXType))
	natXType = natXType.Add(
		natXType,
		big.NewInt(3-int64(log2(n.size*8))),
	)
	return leb128.EncodeSigned(natXType)
}

func (n Nat) EncodeValue(v interface{}) ([]byte, error) {
	switch n.size {
	case 0:
		v, ok := v.(*big.Int)
		if !ok {
			return nil, fmt.Errorf("invalid value: %v", v)
		}
		return leb128.EncodeUnsigned(v)
	case 8:
		v, err := encodeNat64(v)
		if err != nil {
			return nil, err
		}
		bs := make([]byte, 8)
		binary.LittleEndian.PutUint64(bs, v)
		return bs, nil
	case 4:
		v, err := encodeNat32(v)
		if err != nil {
			return nil, err
		}
		bs := make([]byte, 4)
		binary.LittleEndian.PutUint32(bs, v)
		return bs, nil
	case 2:
		v, err := encodeNat16(v)
		if err != nil {
			return nil, err
		}
		bs := make([]byte, 2)
		binary.LittleEndian.PutUint16(bs, v)
		return bs, nil
	case 1:
		v, err := encodeNat8(v)
		if err != nil {
			return nil, err
		}
		return []byte{v}, nil
	default:
		return nil, fmt.Errorf("invalid argument: %v", v)
	}
}

func (n Nat) String() string {
	if n.size == 0 {
		return "nat"
	}
	return fmt.Sprintf("nat%d", n.size*8)
}
