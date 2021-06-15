package idl

import (
	"math/big"

	"github.com/allusion-be/leb128"
)

type Int big.Int

func (Int) Name() string {
	return "nat"
}

func (Int) Encode() []byte {
	bs, _ := leb128.EncodeSigned(intType)
	return bs
}

func (n Int) EncodeValue() []byte {
	bi := big.Int(n)
	bs, _ := leb128.EncodeSigned(&bi)
	return bs
}

func (Int) BuildTypeTable(*TypeTable) {}
