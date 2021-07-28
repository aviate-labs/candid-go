package idl

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/allusion-be/leb128"
)

type Int struct {
	base uint8
	primType
}

func Int16() *Int {
	return &Int{
		base: 16,
	}
}

func Int32() *Int {
	return &Int{
		base: 32,
	}
}

func Int64() *Int {
	return &Int{
		base: 64,
	}
}

func Int8() *Int {
	return &Int{
		base: 8,
	}
}

func (n *Int) Decode(r *bytes.Reader) (interface{}, error) {
	if n.base == 0 {
		return leb128.DecodeSigned(r)
	}
	return readInt(r, int(n.base/8))
}

func (n Int) EncodeType(_ *TypeDefinitionTable) ([]byte, error) {
	if n.base == 0 {
		return leb128.EncodeSigned(big.NewInt(intType))
	}
	intXType := new(big.Int).Set(big.NewInt(intXType))
	intXType = intXType.Add(
		intXType,
		big.NewInt(3-int64(log2(n.base))),
	)
	return leb128.EncodeSigned(intXType)
}

func (n Int) EncodeValue(v interface{}) ([]byte, error) {
	v_, ok := v.(*big.Int)
	if !ok {
		return nil, fmt.Errorf("invalid argument: %v", v)
	}
	if n.base == 0 {
		return leb128.EncodeSigned(v_)
	}
	{
		exp := big.NewInt(int64(n.base) - 1)
		lim := big.NewInt(2)
		lim = lim.Exp(lim, exp, nil)
		min := new(big.Int).Set(lim)
		min = min.Mul(min, big.NewInt(-1))
		max := new(big.Int).Set(lim)
		max = max.Add(max, big.NewInt(-1))
		if v_.Cmp(min) < 0 || max.Cmp(v_) < 0 {
			return nil, fmt.Errorf("invalid value: %s", v_)
		}
	}
	return writeInt(v_, int(n.base/8)), nil
}

func (n Int) String() string {
	if n.base == 0 {
		return "int"
	}
	return fmt.Sprintf("int%d", n.base)
}
