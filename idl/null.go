package idl

import (
	"bytes"

	"github.com/allusion-be/leb128"
)

type Null struct{}

func (Null) Name() string {
	return "null"
}

func (Null) EncodeType() []byte {
	bs, _ := leb128.EncodeSigned(nullType)
	return bs
}

func (Null) EncodeValue() []byte {
	return []byte{}
}

func (n *Null) Decode(r *bytes.Reader) error {
	return nil
}

func (Null) BuildTypeTable(*TypeTable) {}
