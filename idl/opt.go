package idl

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/allusion-be/leb128"
)

type Opt struct {
	typ Type
}

func NewOpt(t Type) *Opt {
	return &Opt{
		typ: t,
	}
}

func (o Opt) AddTypeDefinition(tdt *TypeDefinitionTable) error {
	if err := o.typ.AddTypeDefinition(tdt); err != nil {
		return err
	}

	id, err := leb128.EncodeSigned(big.NewInt(optType))
	if err != nil {
		return err
	}
	v, err := o.typ.EncodeType(tdt)
	if err != nil {
		return err
	}
	tdt.Add(o, concat(id, v))
	return nil
}

func (o Opt) Decode(r *bytes.Reader) (interface{}, error) {
	l, err := r.ReadByte()
	if err != nil {
		return nil, err
	}
	switch l {
	case 0x00:
		return nil, nil
	case 0x01:
		return o.typ.Decode(r)
	default:
		return nil, fmt.Errorf("invalid option value")
	}
}

func (o Opt) EncodeType(tdt *TypeDefinitionTable) ([]byte, error) {
	idx, ok := tdt.Indexes[o.String()]
	if !ok {
		return nil, fmt.Errorf("missing type index for: %s", o)
	}
	return leb128.EncodeSigned(big.NewInt(int64(idx)))
}

func (o Opt) EncodeValue(v interface{}) ([]byte, error) {
	if v == nil {
		return []byte{0x00}, nil
	}
	v_, err := o.typ.EncodeValue(v)
	if err != nil {
		return nil, err
	}
	return concat([]byte{0x01}, v_), nil
}

func (o Opt) String() string {
	return fmt.Sprintf("opt %s", o.typ)
}
