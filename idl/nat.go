package idl

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/allusion-be/leb128"
)

type Nat struct {
	i    *big.Int
	base uint8
}

func Nat16() *Nat {
	return &Nat{
		i:    new(big.Int),
		base: 16,
	}
}

func Nat32() *Nat {
	return &Nat{
		i:    new(big.Int),
		base: 32,
	}
}

func Nat64() *Nat {
	return &Nat{
		i:    new(big.Int),
		base: 64,
	}
}

func Nat8() *Nat {
	return &Nat{
		i:    new(big.Int),
		base: 8,
	}
}

func NewNat(i *big.Int) *Nat {
	return &Nat{i: i}
}

func NewNat16(i uint16) *Nat {
	return &Nat{
		i:    big.NewInt(int64(i)),
		base: 16,
	}
}

func NewNat32(i uint32) *Nat {
	return &Nat{
		i:    big.NewInt(int64(i)),
		base: 32,
	}
}

func NewNat64(i uint64) *Nat {
	return &Nat{
		i:    big.NewInt(int64(i)),
		base: 64,
	}
}

func NewNat8(i uint8) *Nat {
	return &Nat{
		i:    big.NewInt(int64(i)),
		base: 8,
	}
}

func (Nat) BuildTypeTable(*TypeTable) {}

func (n *Nat) Decode(r *bytes.Reader) error {
	if n.base == 0 {
		bi, err := leb128.DecodeUnsigned(r)
		if err != nil {
			return err
		}
		n.i = bi
		return nil
	}
	bs, err := readUInt(r, int(n.base/8))
	if err != nil {
		return err
	}
	n.i.Set(bs)
	return nil
}

func (n Nat) EncodeType() []byte {
	if n.base == 0 {
		bs, _ := leb128.EncodeSigned(natType)
		return bs
	}
	natXType := new(big.Int).Set(natXType)
	natXType = natXType.Add(
		natXType,
		big.NewInt(3-int64(log2(n.base))),
	)
	bs, _ := leb128.EncodeSigned(natXType)
	return bs
}

func (n Nat) EncodeValue() []byte {
	if n.base == 0 {
		bs, _ := leb128.EncodeUnsigned(n.i)
		return bs
	}
	return writeInt(n.i, int(n.base/8))
}

func (Nat) Name() string {
	return "nat"
}

func (n Nat) String() string {
	if n.base == 0 {
		return fmt.Sprintf("nat: %s", n.i)
	}
	return fmt.Sprintf("nat%d: %s", n.base, n.i)
}
