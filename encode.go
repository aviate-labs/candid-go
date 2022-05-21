package candid

import (
	"fmt"
	"math/big"
	"reflect"

	"github.com/aviate-labs/candid-go/idl"
	"github.com/aviate-labs/candid-go/idl2"
	"github.com/aviate-labs/candid-go/typ"
	"github.com/aviate-labs/leb128"
)

func Marshal(args []interface{}) ([]byte, error) {
	e := newEncodeState()
	tdt, err := types(args, e)
	if err != nil {
		return nil, err
	}
	data, err := values(args, e)
	if err != nil {
		return nil, err
	}
	return concat([]byte{'D', 'I', 'D', 'L'}, tdt, data), nil
}

func types(args []interface{}, e *encodeState) ([]byte, error) {
	for _, v := range args {
		v := reflect.ValueOf(v)
		_ = v // TODO
	}

	tdtl, err := leb128.EncodeSigned(big.NewInt(int64(len(e.tdt.Indexes))))
	if err != nil {
		return nil, err
	}
	var tdte []byte
	for _, t := range e.tdt.Types {
		tdte = append(tdte, t...)
	}
	return append(tdtl, tdte...), nil
}

func values(args []interface{}, e *encodeState) ([]byte, error) {
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
	return concat(tsl, ts, vs), nil
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
			return idl2.EncodeNull()
		}
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Ptr:
		return encode(v.Elem())
	case reflect.Bool:
		return idl2.EncodeBool(v.Bool())
	case reflect.Uint, reflect.Int:
		return nil, nil, fmt.Errorf("use big.Int instead of uint/int")
	case reflect.Struct:
		switch v.Type().String() {
		case "typ.Int":
			bi := v.Interface().(typ.Int)
			return idl2.EncodeInt(bi)
		case "typ.Nat":
			bi := v.Interface().(typ.Nat)
			return idl2.EncodeNat(bi)
		}
		return nil, nil, fmt.Errorf("invalid struct type: %s", v.Type())
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
