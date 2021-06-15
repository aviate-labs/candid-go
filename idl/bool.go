package idl

import "github.com/allusion-be/leb128"

type Bool bool

func (Bool) Name() string {
	return "bool"
}

func (Bool) Encode() []byte {
	bs, _ := leb128.EncodeSigned(boolType)
	return bs
}

func (b Bool) EncodeValue() []byte {
	if b {
		return []byte{0x01}
	}
	return []byte{0x00}
}

func (Bool) BuildTypeTable(*TypeTable) {}
