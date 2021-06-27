package idl

import (
	"bytes"
	"math/big"
)

type Type interface {
	BuildTypeTable(*TypeTable)
	Decode(*bytes.Reader) error
	EncodeType() []byte
	EncodeValue() []byte
	Name() string
}

type TypeID *big.Int

var (
	nullType     TypeID = big.NewInt(-1)
	boolType     TypeID = big.NewInt(-2)
	natType      TypeID = big.NewInt(-3)
	intType      TypeID = big.NewInt(-4)
	natXType     TypeID = big.NewInt(-5)
	intXType     TypeID = big.NewInt(-9)
	floatXType   TypeID = big.NewInt(-13)
	textType     TypeID = big.NewInt(-15)
	reservedType TypeID = big.NewInt(-16)
	emptyType    TypeID = big.NewInt(-17)
)
