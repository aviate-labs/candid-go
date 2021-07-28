package idl

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/allusion-be/leb128"
)

type Null struct {
	primType
}

func (n *Null) Decode(_ *bytes.Reader) (interface{}, error) {
	return nil, nil
}

func (Null) EncodeType() ([]byte, error) {
	return leb128.EncodeSigned(big.NewInt(nullType))
}

func (Null) EncodeValue(v interface{}) ([]byte, error) {
	if v != nil {
		return nil, fmt.Errorf("invalid argument: %v", v)
	}
	return []byte{}, nil
}

func (n Null) String() string {
	return "null"
}
