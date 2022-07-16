package marshal

import (
	"fmt"
	"math/big"
	"reflect"

	"github.com/aviate-labs/candid-go/idl"
	"github.com/aviate-labs/candid-go/typ"
	"github.com/aviate-labs/leb128"
	"github.com/aviate-labs/principal-go"
)

func Marshal(args []any) ([]byte, error) {
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

func encode(v reflect.Value) ([]byte, []byte, error) {
	if v.Kind() == reflect.Interface {
		if v.IsNil() {
			return EncodeNull()
		}
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Ptr:
		return encode(v.Elem())
	case reflect.Bool:
		return EncodeBool(v.Bool())
	case reflect.Uint:
		return nil, nil, fmt.Errorf("use typ.Nat instead of uint")
	case reflect.Int:
		return nil, nil, fmt.Errorf("use typ.Int instead of int")
	case reflect.Uint8:
		return EncodeNat8(uint8(v.Uint()))
	case reflect.Uint16:
		return EncodeNat16(uint16(v.Uint()))
	case reflect.Uint32:
		return EncodeNat32(uint32(v.Uint()))
	case reflect.Uint64:
		return EncodeNat64(uint64(v.Uint()))
	case reflect.Int8:
		return EncodeInt8(int8(v.Uint()))
	case reflect.Int16:
		return EncodeInt16(int16(v.Uint()))
	case reflect.Int32:
		return EncodeInt32(int32(v.Uint()))
	case reflect.Int64:
		return EncodeInt64(int64(v.Uint()))
	case reflect.Float32:
		return EncodeFloat32(float32(v.Float()))
	case reflect.Float64:
		return EncodeFloat64(v.Float())
	case reflect.String:
		return EncodeText(v.String())
	case reflect.Struct:
		switch v.Type().String() {
		case "typ.Empty":
			return EncodeEmpty()
		case "typ.Int":
			bi := v.Interface().(typ.Int)
			return EncodeInt(bi)
		case "typ.Nat":
			bi := v.Interface().(typ.Nat)
			return EncodeNat(bi)
		case "typ.Null":
			return EncodeNull()
		case "typ.Reserved":
			return EncodeReserved()
		case "principal.Principal":
			p := v.Interface().(principal.Principal)
			return EncodePrincipal(p)
		}
		return nil, nil, fmt.Errorf("invalid struct type: %s", v.Type())
	default:
		return nil, nil, fmt.Errorf("invalid primary value: %s", v.Kind())
	}
}

func types(args []any, e *encodeState) ([]byte, error) {
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

func values(args []any, e *encodeState) ([]byte, error) {
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
