package idl

import (
	"bytes"

	"github.com/allusion-be/leb128"
)

func Decode(bs []byte) ([]Type, []interface{}, error) {
	if len(bs) == 0 {
		return nil, nil, &FormatError{
			Description: "empty",
		}
	}

	r := bytes.NewReader(bs)

	{ // 'DIDL'

		magic := make([]byte, 4)
		n, err := r.Read(magic)
		if err != nil {
			return nil, nil, err
		}
		if n < 4 {
			return nil, nil, &FormatError{
				Description: "no magic bytes",
			}
		}
		if !bytes.Equal(magic, []byte{'D', 'I', 'D', 'L'}) {
			return nil, nil, &FormatError{
				Description: "wrong magic bytes",
			}
		}
	}

	var tds []Type
	{ // T
		tdtl, err := leb128.DecodeUnsigned(r)
		if err != nil {
			return nil, nil, err
		}
		for i := 0; i < int(tdtl.Int64()); i++ {
			tid, err := leb128.DecodeSigned(r)
			if err != nil {
				return nil, nil, err
			}
			switch tid.Int64() {
			case optType:
				tid, err := leb128.DecodeSigned(r)
				if err != nil {
					return nil, nil, err
				}
				v, err := getType(tid.Int64(), tds)
				if err != nil {
					return nil, nil, err
				}
				tds = append(tds, &Opt{v})
			case vecType:
				tid, err := leb128.DecodeSigned(r)
				if err != nil {
					return nil, nil, err
				}
				v, err := getType(tid.Int64(), tds)
				if err != nil {
					return nil, nil, err
				}
				tds = append(tds, &Vec{v})
			case recType:
				l, err := leb128.DecodeUnsigned(r)
				if err != nil {
					return nil, nil, err
				}
				var fields []field
				for i := 0; i < int(l.Int64()); i++ {
					h, err := leb128.DecodeUnsigned(r)
					if err != nil {
						return nil, nil, err
					}
					tid, err := leb128.DecodeSigned(r)
					if err != nil {
						return nil, nil, err
					}
					v, err := getType(tid.Int64(), tds)
					if err != nil {
						return nil, nil, err
					}
					fields = append(fields, field{
						s: h.String(),
						t: v,
					})
				}
				tds = append(tds, &Rec{fields: fields})
			}
		}
	}

	tsl, err := leb128.DecodeUnsigned(r)
	if err != nil {
		return nil, nil, err
	}

	var ts []Type
	{ // I
		for i := 0; i < int(tsl.Int64()); i++ {
			tid, err := leb128.DecodeSigned(r)
			if err != nil {
				return nil, nil, err
			}
			t, err := getType(tid.Int64(), tds)
			if err != nil {
				return nil, nil, err
			}
			ts = append(ts, t)
		}
	}

	var vs []interface{}
	{ // M
		for i := 0; i < int(tsl.Int64()); i++ {
			v, err := ts[i].Decode(r)
			if err != nil {
				return nil, nil, err
			}
			vs = append(vs, v)
		}
	}

	return ts, vs, nil
}
