package idl

import (
	"bytes"
	"fmt"
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

var (
	nullType      int64 = -1
	boolType      int64 = -2
	natType       int64 = -3
	intType       int64 = -4
	natXType      int64 = -5
	intXType      int64 = -9
	floatXType    int64 = -13
	textType      int64 = -15
	reservedType  int64 = -16
	emptyType     int64 = -17
	optType       int64 = -18
	vecType       int64 = -19
	recordType    int64 = -20
	variantType   int64 = -21
	funcType      int64 = -22
	serviceType   int64 = -23
	principalType int64 = -24
)

func getType(t int64) (Type, error) {
	switch t {
	case nullType:
		return new(Null), nil
	case boolType:
		return new(Bool), nil
	case natType:
		return new(Nat), nil
	case intType:
		return new(Int), nil
	case natXType:
		return Nat8(), nil
	case natXType - 1:
		return Nat16(), nil
	case natXType - 2:
		return Nat32(), nil
	case natXType - 3:
		return Nat64(), nil
	case intXType:
		return Int8(), nil
	case intXType - 1:
		return Int16(), nil
	case intXType - 2:
		return Int32(), nil
	case intXType - 3:
		return Int64(), nil
	case floatXType:
		return Float32(), nil
	case floatXType - 1:
		return Float64(), nil
	case textType:
		return new(Text), nil
	case reservedType:
		return new(Reserved), nil
	case emptyType:
		return new(Empty), nil
	default:
		if t < -24 {
			return nil, &FormatError{
				Description: "type: out of range",
			}
		}
		return nil, &FormatError{
			Description: "type: not primitive",
		}
	}
}

type PrimType interface {
	prim()
}

type primType struct{}

func (primType) prim() {}

func (primType) BuildTypeTable(_ *TypeTable) {}
