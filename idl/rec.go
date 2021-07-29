package idl

import (
	"bytes"
	"fmt"
	"math/big"
	"strings"

	"github.com/allusion-be/leb128"
)

type Rec struct {
	fields []field
}

func NewRec(fields map[string]Type) *Rec {
	var rec Rec
	for k, v := range fields {
		rec.fields = append(rec.fields, field{
			s: k,
			t: v,
		})
	}
	return &rec
}

func (r Rec) AddTypeDefinition(tdt *TypeDefinitionTable) error {
	for _, f := range r.fields {
		if err := f.t.AddTypeDefinition(tdt); err != nil {
			return err
		}
	}

	id, err := leb128.EncodeSigned(big.NewInt(recType))
	if err != nil {
		return err
	}
	l, err := leb128.EncodeUnsigned(big.NewInt(int64(len(r.fields))))
	if err != nil {
		return err
	}
	var vs []byte
	for _, f := range r.fields {
		l, err := leb128.EncodeUnsigned(Hash(f.s))
		if err != nil {
			return nil
		}
		t, err := f.t.EncodeType(tdt)
		if err != nil {
			return nil
		}
		vs = append(vs, concat(l, t)...)

	}

	tdt.Add(r, concat(id, l, vs))
	return nil
}

func (r Rec) Decode(r_ *bytes.Reader) (interface{}, error) {
	rec := make(map[string]interface{})
	for _, f := range r.fields {
		v, err := f.t.Decode(r_)
		if err != nil {
			return nil, err
		}
		rec[f.s] = v
	}
	if len(rec) == 0 {
		return nil, nil
	}
	return rec, nil
}

func (r Rec) EncodeType(tdt *TypeDefinitionTable) ([]byte, error) {
	idx, ok := tdt.Indexes[r.String()]
	if !ok {
		return nil, fmt.Errorf("missing type index for: %s", r)
	}
	return leb128.EncodeSigned(big.NewInt(int64(idx)))
}

func (r Rec) EncodeValue(v interface{}) ([]byte, error) {
	if v == nil {
		return nil, nil
	}
	fs, ok := v.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid argument: %v", v)
	}
	var vs_ []interface{}
	for _, f := range r.fields {
		vs_ = append(vs_, fs[f.s])
	}
	var vs []byte
	for i, f := range r.fields {
		v_, err := f.t.EncodeValue(vs_[i])
		if err != nil {
			return nil, err
		}
		vs = append(vs, v_...)
	}
	return vs, nil
}

func (r Rec) String() string {
	var s []string
	for _, f := range r.fields {
		s = append(s, fmt.Sprintf("%s:%s", f.s, f.t.String()))
	}
	return fmt.Sprintf("record {%s}", strings.Join(s, "; "))
}
