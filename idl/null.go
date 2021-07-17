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

func (Null) EncodeType() []byte {
	bs, _ := leb128.EncodeSigned(big.NewInt(nullType))
	return bs
}

func (Null) EncodeValue() []byte {
	return []byte{}
}

func (Null) Name() string {
	return "null"
}

func (n Null) String() string {
	return n.Name()
}
