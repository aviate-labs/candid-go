package idl

import (
	"bytes"
	"fmt"

	"github.com/allusion-be/leb128"
)

type Bool bool

func (Bool) BuildTypeTable(*TypeTable) {}

func (b *Bool) Decode(r *bytes.Reader) error {
	v, err := r.ReadByte()
	if err != nil {
		return err
	}
	switch v {
	case 0x00:
		*b = Bool(false)
	case 0x01:
		*b = Bool(true)
	default:
		return fmt.Errorf("invalid bool values: %x", b)
	}
	return nil
}

func (Bool) EncodeType() []byte {
	bs, _ := leb128.EncodeSigned(boolType)
	return bs
}

func (b Bool) EncodeValue() []byte {
	if b {
		return []byte{0x01}
	}
	return []byte{0x00}
}

func (Bool) Name() string {
	return "bool"
}

func (b Bool) String() string {
	return fmt.Sprintf("bool: %t", b)
}
