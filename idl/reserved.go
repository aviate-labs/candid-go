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

func (Reserved) EncodeType(_ *TypeTable) ([]byte, error) {
	return leb128.EncodeSigned(big.NewInt(reservedType))
}

func (Reserved) EncodeValue() ([]byte, error) {
	return []byte{}, nil
}

func (Reserved) Name() string {
	return "reserved"
}

func (Reserved) String() string {
	return "reserved"
}
