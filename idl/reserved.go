package idl

import (
	"bytes"

	"github.com/allusion-be/leb128"
)

type Reserved struct{}

func (Reserved) BuildTypeTable(*TypeTable) {}

func (Reserved) Decode(*bytes.Reader) error {
	return nil
}

func (Reserved) EncodeType() []byte {
	bs, _ := leb128.EncodeSigned(reservedType)
	return bs
}

func (Reserved) EncodeValue() []byte {
	return []byte{}
}

func (Reserved) Name() string {
	return "reserved"
}
