package candid

import (
	"fmt"
	"math/big"
	"reflect"

	"github.com/aviate-labs/candid-go/idl"
	"github.com/aviate-labs/leb128"
)

func Marshal(args []interface{}) ([]byte, error) {
	e := newEncodeState()
	for _, v := range args {
		addTypes(reflect.ValueOf(v), e)
	}

	tdtl, err := leb128.EncodeSigned(big.NewInt(int64(len(e.tdt.Indexes))))
	if err != nil {
		return nil, err
	}
	var tdte []byte
	for _, t := range e.tdt.Types {
		tdte = append(tdte, t...)
	}

	tsl, err := leb128.EncodeSigned(big.NewInt(int64(len(args))))
	if err != nil {
		return nil, err
	}
	var (
		ts []byte
		vs []byte
	)
	for _, v := range args {
		t, v, err := encode(reflect.ValueOf(v))
		if err != nil {
			return nil, err
		}
		ts = append(ts, t...)
		vs = append(vs, v...)
	}
	return concat(
		[]byte{'D', 'I', 'D', 'L'},
		tdtl, tdte, tsl, ts, vs,
	), nil
}

func addTypes(v reflect.Value, e *encodeState) error {
	return nil
}

func concat(bs ...[]byte) []byte {
	var c []byte
	for _, b := range bs {
		c = append(c, b...)
	}
	return c
}

func encode(v reflect.Value) ([]byte, []byte, error) {
	if v.Kind() == reflect.Interface {
		if v.IsNil() {
			return nil, nil, fmt.Errorf("is nil")
		}
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Bool:
		typ, err := leb128.EncodeSigned(big.NewInt(-2))
		if err != nil {
			return nil, nil, err
		}
		if b := v.Bool(); b {
			return typ, []byte{0x01}, nil
		} else {
			return typ, []byte{0x00}, nil
		}
	default:
		return nil, nil, fmt.Errorf("invalid primary value: %s", v.Kind())
	}
}

type encodeState struct {
	tdt *idl.TypeDefinitionTable
}

func newEncodeState() *encodeState {
	return &encodeState{
		tdt: &idl.TypeDefinitionTable{
			Indexes: make(map[string]int),
		},
	}
}
