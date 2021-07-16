package idl

import (
	"bytes"
	"math/big"

	"github.com/allusion-be/leb128"
)

func Decode(bs []byte) (Tuple, error) {
	if len(bs) == 0 {
		return nil, &FormatError{
			Description: "empty",
		}
	}

	rs := bytes.NewReader(bs)
	magic := make([]byte, 4)
	n, err := rs.Read(magic)
	if err != nil {
		return nil, err
	}
	if n < 4 {
		return nil, &FormatError{
			Description: "no magic bytes",
		}
	}
	if !bytes.Equal(magic, []byte{'D', 'I', 'D', 'L'}) {
		return nil, &FormatError{
			Description: "wrong magic bytes",
		}
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
			v = Nat8()
		case -6:
			v = Nat16()
		case -7:
			v = Nat32()
		case -8:
			v = Nat64()
		case -9:
			v = Int8()
		case -10:
			v = Int16()
		case -11:
			v = Int32()
		case -12:
			v = Int64()
		case -13:
			v = Float32()
		case -14:
			v = Float64()
		case -15:
			v = new(Text)
		default:
			return nil, &FormatError{
				Description: "wrong type",
			}
		}
		if err := v.Decode(rs); err != nil {
			return nil, err
		}
		vs = append(vs, v)
	}
	return vs, nil
}
