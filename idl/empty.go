package idl

import (
	"bytes"
	"math/big"

	"github.com/allusion-be/leb128"
)

type Empty struct {
	primType
}

func (Empty) Decode(*bytes.Reader) error {
	return nil
}

func (Empty) EncodeType(_ *TypeTable) ([]byte, error) {
	return leb128.EncodeSigned(big.NewInt(emptyType))
}

func (Empty) EncodeValue() ([]byte, error) {
	return []byte{}, nil
}

func (Empty) Name() string {
	return "empty"
}

func (Empty) String() string {
	return "empty"
}
