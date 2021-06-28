package idl

import (
	"bytes"

	"github.com/allusion-be/leb128"
)

type Empty struct {
	primType
}

func (Empty) Decode(*bytes.Reader) error {
	return nil
}

func (Empty) EncodeType() []byte {
	bs, _ := leb128.EncodeSigned(emptyType)
	return bs
}

func (Empty) EncodeValue() []byte {
	return []byte{}
}

func (Empty) Name() string {
	return "empty"
}
