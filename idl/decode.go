package idl

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/allusion-be/leb128"
)

func Decode(bs []byte) ([]Type, error) {
	rs := bytes.NewReader(bs)
	magic := make([]byte, 4)
	n, err := rs.Read(magic)
	if err != nil {
		return nil, err
	}
	if n < 4 || !bytes.Equal(magic, []byte{'D', 'I', 'D', 'L'}) {
		return nil, fmt.Errorf("invalid magic number: %x", magic)
	}
	if _, err := NewTable(rs); err != nil {
		return nil, err
	}
	l, err := leb128.DecodeUnsigned(rs)
	if err != nil {
		return nil, err
	}
	var ts []TypeID
	for i := new(big.Int).Set(l); i.Sign() > 0; i.Add(i, big.NewInt(-1)) {
		t, err := leb128.DecodeSigned(rs)
		if err != nil {
			return nil, err
		}
		ts = append(ts, t)
	}
	var vs []Type
	for _, t := range ts {
		var v Type
		switch n := (*big.Int)(t).Int64(); n {
		case -1:
			v = new(Null)
		case -2:
			v = new(Bool)
		case -3:
			v = new(Nat)
		case -4:
			v = new(Int)
		case -5:
			v = NewNat8(0)
		case -6:
			v = NewNat16(0)
		case -7:
			v = NewNat32(0)
		case -8:
			v = NewNat64(0)
		default:
			panic(n)
		}
		if err := v.Decode(rs); err != nil {
			return nil, err
		}
		vs = append(vs, v)
	}
	return vs, nil
}
