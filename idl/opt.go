package idl

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/aviate-labs/leb128"
)

type OptionalType[typ Type] struct {
	Type typ
}

func NewOptionalType[typ Type](t typ) *OptionalType[typ] {
	return &OptionalType[typ]{
		Type: t,
	}
}

func (o OptionalType[t]) AddTypeDefinition(tdt *TypeDefinitionTable) error {
	if err := o.Type.AddTypeDefinition(tdt); err != nil {
		return err
	}

	id, err := leb128.EncodeSigned(big.NewInt(optType))
	if err != nil {
		return err
	}
	v, err := o.Type.EncodeType(tdt)
	if err != nil {
		return err
	}
	tdt.Add(o, concat(id, v))
	return nil
}

func (o OptionalType[t]) Decode(r *bytes.Reader) (interface{}, error) {
	l, err := r.ReadByte()
	if err != nil {
		return nil, err
	}
	switch l {
	case 0x00:
		return nil, nil
	case 0x01:
		return o.Type.Decode(r)
	default:
		return nil, fmt.Errorf("invalid option value")
	}
}

func (o OptionalType[t]) EncodeType(tdt *TypeDefinitionTable) ([]byte, error) {
	idx, ok := tdt.Indexes[o.String()]
	if !ok {
		return nil, fmt.Errorf("missing type index for: %s", o)
	}
	return leb128.EncodeSigned(big.NewInt(int64(idx)))
}

func (o OptionalType[t]) EncodeValue(v interface{}) ([]byte, error) {
	if v == nil {
		return []byte{0x00}, nil
	}
	v_, err := o.Type.EncodeValue(v)
	if err != nil {
		return nil, err
	}
	return concat([]byte{0x01}, v_), nil
}

func (o OptionalType[t]) String() string {
	return fmt.Sprintf("opt %s", o.Type)
}
