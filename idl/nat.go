package idl

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/allusion-be/leb128"
)

type Nat struct {
	base uint8
	primType
}

func Nat16() *Nat {
	return &Nat{
		base: 16,
	}
}

func Nat32() *Nat {
	return &Nat{
		base: 32,
	}
}

func Nat64() *Nat {
	return &Nat{
		base: 64,
	}
}

func Nat8() *Nat {
	return &Nat{
		base: 8,
	}
}

func (n *Nat) Decode(r *bytes.Reader) (interface{}, error) {
	if n.base == 0 {
		return leb128.DecodeUnsigned(r)
	}
	return readUInt(r, int(n.base/8))
}

func (n Nat) EncodeType(_ *TypeDefinitionTable) ([]byte, error) {
	if n.base == 0 {
		return leb128.EncodeSigned(big.NewInt(natType))
	}
	natXType := new(big.Int).Set(big.NewInt(natXType))
	natXType = natXType.Add(
		natXType,
		big.NewInt(3-int64(log2(n.base))),
	)
	return leb128.EncodeSigned(natXType)
}

func (n Nat) EncodeValue(v interface{}) ([]byte, error) {
	v_, ok := v.(*big.Int)
	if !ok {
		return nil, fmt.Errorf("invalid argument: %v", v)
	}
	if n.base == 0 {
		return leb128.EncodeUnsigned(v_)
	}
	{
		lim := big.NewInt(2)
		lim = lim.Exp(lim, big.NewInt(int64(n.base)), nil)
		if lim.Cmp(v_) <= 0 {
			return nil, fmt.Errorf("invalid value: %s", v_)
		}
	}
	return writeInt(v_, int(n.base/8)), nil
}

func (n Nat) String() string {
	if n.base == 0 {
		return "nat"
	}
	return fmt.Sprintf("nat%d", n.base)
}
