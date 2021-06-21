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
		switch n := (*big.Int)(t).Int64(); n {
		case -1:
			vs = append(vs, Null{})
		case -2:
			b, err := rs.ReadByte()
			if err != nil {
				return nil, err
			}
			switch b {
			case 0x00:
				vs = append(vs, Bool(false))
			case 0x01:
				vs = append(vs, Bool(true))
			default:
				return nil, fmt.Errorf("invalid bool values: %x", b)
			}
		case -3:
			n, err := leb128.DecodeUnsigned(rs)
			if err != nil {
				return nil, err
			}
			vs = append(vs, Nat(*n))
		case -4:
			n, err := leb128.DecodeSigned(rs)
			if err != nil {
				return nil, err
			}
			vs = append(vs, Int(*n))
		default:
			panic(n)
		}
	}
	return vs, nil
}
