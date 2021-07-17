package idl

import (
	"bytes"
	"io"
	"math/big"

	"github.com/allusion-be/leb128"
)

func Decode(bs []byte) (Tuple, error) {
	if len(bs) == 0 {
		return nil, &FormatError{
			Description: "empty",
		}
	}

	// Validate the magic number (DIDL).
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
	var ts []int64
	for i := new(big.Int).Set(l); i.Sign() > 0; i.Add(i, big.NewInt(-1)) {
		t, err := leb128.DecodeSigned(rs)
		if err != nil {
			return nil, err
		}
		ts = append(ts, t.Int64())
	}
	var vs []Type
	for _, t := range ts {
		v, err := getType(t)
		if err != nil {
			return nil, err
		}
		if err := v.Decode(rs); err != nil {
			if err == io.EOF {
				return nil, &FormatError{
					Description: "end of data",
				}
			}
			return nil, err
		}
		vs = append(vs, v)
	}
	if rs.Len() != 0 {
		return nil, &FormatError{
			Description: "too long",
		}
	}
	return vs, nil
}
