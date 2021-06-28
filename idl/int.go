package idl

import (
	"bytes"
	"fmt"
	"io"
	"math/big"

	"github.com/allusion-be/leb128"
)

type Int struct {
	i    *big.Int
	base uint8
	primType
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

func Int8() *Int {
	return &Int{
		i:    new(big.Int),
		base: 8,
	}
}

func NewInt(i *big.Int) *Int {
	return &Int{i: i}
}

func NewInt16(i int16) *Int {
	return &Int{
		i:    big.NewInt(int64(i)),
		base: 16,
	}
}

func NewInt32(i int32) *Int {
	return &Int{
		i:    big.NewInt(int64(i)),
		base: 32,
	}
}

func NewInt64(i *big.Int) *Int {
	return &Int{
		i:    i,
		base: 64,
	}
}

func NewInt8(i int8) *Int {
	return &Int{
		i:    big.NewInt(int64(i)),
		base: 8,
	}
}

func (n *Int) Decode(r *bytes.Reader) error {
	if n.base == 0 {
		bi, err := leb128.DecodeSigned(r)
		if err != nil {
			return err
		}
		n.i = bi
		return nil
	}
	raw, _ := io.ReadAll(r)
	*r = *bytes.NewReader(raw)
	bi, err := readInt(r, int(n.base/8))
	if err != nil {
		return err
	}
	n.i.Set(bi)
	return nil
}

func (n Int) EncodeType() []byte {
	if n.base == 0 {
		bs, _ := leb128.EncodeSigned(intType)
		return bs
	}
	intXType := new(big.Int).Set(intXType)
	intXType = intXType.Add(
		intXType,
		big.NewInt(3-int64(log2(n.base))),
	)
	bs, _ := leb128.EncodeSigned(intXType)
	return bs
}

func (n Int) EncodeValue() []byte {
	if n.base == 0 {
		bs, _ := leb128.EncodeSigned(n.i)
		return bs
	}

	return writeInt(n.i, int(n.base/8))
}

func (Int) Name() string {
	return "int"
}

func (n Int) String() string {
	return fmt.Sprintf("int: %s", n.i)
}
