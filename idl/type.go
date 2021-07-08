package idl

import (
	"bytes"
	"fmt"
	"math/big"
	"strings"
)

type Tuple []Type

func (ts Tuple) String() string {
	var s []string
	for _, t := range ts {
		s = append(s, t.String())
	}
	return fmt.Sprintf("(%s)", strings.Join(s, ", "))
}

type Type interface {
	BuildTypeTable(*TypeTable)
	Decode(*bytes.Reader) error
	EncodeType() []byte
	EncodeValue() []byte
	Name() string

	fmt.Stringer
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

type PrimType interface {
	prim()
}

type primType struct{}

func (primType) prim() {}

func (primType) BuildTypeTable(_ *TypeTable) {}
