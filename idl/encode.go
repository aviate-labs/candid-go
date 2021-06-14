package idl

import (
	"fmt"
	"math/big"

	"github.com/allusion-be/leb128"
)

func Encode(types []Type, args []interface{}) ([]byte, error) {
	if len(args) != len(types) {
		return nil, fmt.Errorf("invalid number of arguments")
	}
	tt := new(TypeTable)

	var (
		ts []byte
		vs []byte
	)
	for i, t := range types {
		t.BuildTypeTable(tt)
		ts = append(ts, t.Encode()...)
		a := args[i]
		if !t.Covariant(a) {
			return nil, fmt.Errorf("invalid %s argument: %v", t.Name(), a)
		}
		vs = append(vs, t.EncodeValue(a)...)
	}

	magic := []byte{'D', 'I', 'D', 'L'}
	table, err := tt.Encode()
	if err != nil {
		return nil, err
	}
	l, err := leb128.EncodeUnsigned(big.NewInt(int64(len(args))))
	if err != nil {
		return nil, err
	}
	return concat(magic, table, l, ts, vs), nil
}
