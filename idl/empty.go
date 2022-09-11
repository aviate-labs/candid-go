package idl

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/aviate-labs/leb128"
)

type EmptyType struct {
	primType
}

func (EmptyType) Decode(*bytes.Reader) (interface{}, error) {
	return nil, fmt.Errorf("cannot decode empty type")
}

func (EmptyType) EncodeType(_ *TypeDefinitionTable) ([]byte, error) {
	return leb128.EncodeSigned(big.NewInt(emptyType))
}

func (EmptyType) EncodeValue(_ interface{}) ([]byte, error) {
	return []byte{}, nil
}

func (EmptyType) String() string {
	return "empty"
}
