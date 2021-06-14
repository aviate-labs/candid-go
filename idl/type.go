package idl

import "math/big"

type TypeID *big.Int

var (
	nullType TypeID = big.NewInt(-1)
)

type Type interface {
	Name() string
	Encode() []byte
	EncodeValue(v interface{}) []byte
	BuildTypeTable(*TypeTable)
	Covariant(v interface{}) bool
}
