package idl

import (
	"bytes"

	"github.com/allusion-be/leb128"
)

type Null struct {
	primType
}

func (n *Null) Decode(_ *bytes.Reader) error {
	return nil
}

func (Null) EncodeType() []byte {
	bs, _ := leb128.EncodeSigned(nullType)
	return bs
}

func (Null) EncodeValue() []byte {
	return []byte{}
}

func (Null) Name() string {
	return "null"
}
