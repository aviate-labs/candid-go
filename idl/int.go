package idl

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/allusion-be/leb128"
)

type Int struct {
	i    *big.Int
	base uint8
}

func Int8() *Int {
	return &Int{
		i:    new(big.Int),
		base: 8,
	}
}

func Int16() *Int {
	return &Int{
		i:    new(big.Int),
		base: 16,
	}
}

func Int32() *Int {
	return &Int{
		i:    new(big.Int),
		base: 32,
	}
}

func Int64() *Int {
	return &Int{
		i:    new(big.Int),
		base: 64,
	}
}

func NewInt(i *big.Int) *Int {
	return &Int{i: i}
}

func NewInt8(i uint8) *Int {
	return &Int{
		i:    big.NewInt(int64(i)),
		base: 8,
	}
}

func NewInt16(i uint16) *Int {
	return &Int{
		i:    big.NewInt(int64(i)),
		base: 16,
	}
}

func NewInt32(i uint32) *Int {
	return &Int{
		i:    big.NewInt(int64(i)),
		base: 32,
	}
}

func NewInt64(i uint64) *Int {
	return &Int{
		i:    big.NewInt(int64(i)),
		base: 64,
	}
}

func (Int) Name() string {
	return "int"
}

func (Int) Encode() []byte {
	bs, _ := leb128.EncodeSigned(intType)
	return bs
}

func (n Int) EncodeValue() []byte {
	bs, _ := leb128.EncodeSigned(n.i)
	return bs
}

func (n *Int) Decode(r *bytes.Reader) error {
	bi, err := leb128.DecodeSigned(r)
	if err != nil {
		return err
	}
	n.i = bi
	return nil
}

func (Int) BuildTypeTable(*TypeTable) {}

func (n Int) String() string {
	return fmt.Sprintf("int: %s", n.i)
}
