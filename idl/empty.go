package idl

import (
	"bytes"
	"math/big"

	"github.com/allusion-be/leb128"
)

type Empty struct {
	primType
}

func (Empty) Decode(*bytes.Reader) (interface{}, error) {
	return nil, nil
}

func (Empty) EncodeType() ([]byte, error) {
	return leb128.EncodeSigned(big.NewInt(emptyType))
}

func (Empty) EncodeValue(_ interface{}) ([]byte, error) {
	return []byte{}, nil
}

func (Empty) String() string {
	return "empty"
}
