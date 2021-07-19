package idl

import (
	"bytes"
	"math/big"

	"github.com/allusion-be/leb128"
)

type Null struct {
	primType
}

func (n *Null) Decode(_ *bytes.Reader) error {
	return nil
}

func (Null) EncodeType(_ *TypeTable) ([]byte, error) {
	return leb128.EncodeSigned(big.NewInt(nullType))
}

func (Null) EncodeValue() ([]byte, error) {
	return []byte{}, nil
}

func (Null) Name() string {
	return "null"
}

func (n Null) String() string {
	return n.Name()
}
