package idl

import (
	"bytes"
	"math/big"
)

type TypeID *big.Int

var (
	nullType TypeID = big.NewInt(-1)
	boolType TypeID = big.NewInt(-2)
	natType  TypeID = big.NewInt(-3)
	intType  TypeID = big.NewInt(-4)
)

type Type interface {
	Name() string
	Encode() []byte
	EncodeValue() []byte
	Decode(*bytes.Reader) error
	BuildTypeTable(*TypeTable)
}
