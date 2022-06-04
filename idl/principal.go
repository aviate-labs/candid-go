package idl

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/aviate-labs/leb128"
	"github.com/aviate-labs/principal-go"
)

type Principal struct {
	primType
}

func (Principal) Decode(r *bytes.Reader) (interface{}, error) {
	b, err := r.ReadByte()
	if err != nil {
		return nil, err
	}
	if b != 0x01 {
		return nil, fmt.Errorf("cannot decode principal")
	}
	l, err := leb128.DecodeUnsigned(r)
	if err != nil {
		return nil, err
	}
	if l.Uint64() == 0 {
		return &principal.Principal{Raw: []byte{}}, nil
	}
	v := make([]byte, l.Uint64())
	if _, err := r.Read(v); err != nil {
		return nil, err
	}
	return &principal.Principal{Raw: v}, nil
}

func (Principal) EncodeType(_ *TypeDefinitionTable) ([]byte, error) {
	return leb128.EncodeSigned(big.NewInt(principalType))
}

func (Principal) EncodeValue(v interface{}) ([]byte, error) {
	v_, ok := v.(*principal.Principal)
	if !ok {
		return nil, fmt.Errorf("invalid argument: %v", v)
	}
	l, err := leb128.EncodeUnsigned(big.NewInt(int64(len(v_.Raw))))
	if err != nil {
		return nil, err
	}
	return concat([]byte{0x01}, l, v_.Raw), nil
}

func (Principal) String() string {
	return "principal"
}
