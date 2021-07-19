package idl

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/allusion-be/leb128"
)

type Bool struct {
	v bool
	primType
}

func NewBool(b bool) *Bool {
	return &Bool{
		v: b,
	}
}

func (b *Bool) Decode(r *bytes.Reader) error {
	v, err := r.ReadByte()
	if err != nil {
		return err
	}
	switch v {
	case 0x00:
		*b = Bool{v: false}
	case 0x01:
		*b = Bool{v: true}
	default:
		return fmt.Errorf("invalid bool values: %x", b)
	}
	return nil
}

func (Bool) EncodeType(_ *TypeTable) ([]byte, error) {
	return leb128.EncodeSigned(big.NewInt(boolType))
}

func (b Bool) EncodeValue() ([]byte, error) {
	if b.v {
		return []byte{0x01}, nil
	}
	return []byte{0x00}, nil
}

func (Bool) Name() string {
	return "bool"
}

func (b Bool) String() string {
	return fmt.Sprintf("bool: %t", b)
}
