package idl

import (
	"math/big"

	"github.com/allusion-be/leb128"
)

func Encode(types []Type) ([]byte, error) {
	tt := new(TypeTable)
	var (
		ts []byte
		vs []byte
	)
	for _, t := range types {
		t.BuildTypeTable(tt)
		ts = append(ts, t.Encode()...)
		vs = append(vs, t.EncodeValue()...)
	}

	magic := []byte{'D', 'I', 'D', 'L'}
	table, err := tt.Encode()
	if err != nil {
		return nil, err
	}
	l, err := leb128.EncodeUnsigned(big.NewInt(int64(len(types))))
	if err != nil {
		return nil, err
	}
	return concat(magic, table, l, ts, vs), nil
}
