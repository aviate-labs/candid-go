package idl

import (
	"math/big"

	"github.com/allusion-be/leb128"
)

type Nat big.Int

func (Nat) Name() string {
	return "nat"
}

func (Nat) Encode() []byte {
	bs, _ := leb128.EncodeSigned(natType)
	return bs
}

func (n Nat) EncodeValue() []byte {
	bi := big.Int(n)
	bs, _ := leb128.EncodeUnsigned(&bi)
	return bs
}

func (Nat) BuildTypeTable(*TypeTable) {}
