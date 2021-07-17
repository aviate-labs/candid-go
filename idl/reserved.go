package idl

import (
	"bytes"
	"math/big"

	"github.com/allusion-be/leb128"
)

type Reserved struct {
	primType
}

func (Reserved) Decode(*bytes.Reader) error {
	return nil
}

func (Reserved) EncodeType() []byte {
	bs, _ := leb128.EncodeSigned(big.NewInt(reservedType))
	return bs
}

func (Reserved) EncodeValue() []byte {
	return []byte{}
}

func (Reserved) Name() string {
	return "reserved"
}

func (Reserved) String() string {
	return "reserved"
}
