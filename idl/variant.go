package idl

import (
	"bytes"
	"fmt"
	"math/big"
	"sort"
	"strings"

	"github.com/allusion-be/leb128"
)

type FieldValue struct {
	Name  string
	Value interface{}
}

type Variant struct {
	fields []Field
}

func NewVariant(fields map[string]Type) *Variant {
	var variant Variant
	for k, v := range fields {
		variant.fields = append(variant.fields, Field{
			Name: k,
			Type: v,
		})
	}
	sort.Slice(variant.fields, func(i, j int) bool {
		return Hash(variant.fields[i].Name).Cmp(Hash(variant.fields[j].Name)) < 0
	})
	return &variant
}

func (v Variant) AddTypeDefinition(tdt *TypeDefinitionTable) error {
	for _, f := range v.fields {
		if err := f.Type.AddTypeDefinition(tdt); err != nil {
			return err
		}
	}

	id, err := leb128.EncodeSigned(big.NewInt(varType))
	if err != nil {
		return err
	}
	l, err := leb128.EncodeUnsigned(big.NewInt(int64(len(v.fields))))
	if err != nil {
		return err
	}
	var vs []byte
	for _, f := range v.fields {
		id, err := leb128.EncodeUnsigned(Hash(f.Name))
		if err != nil {
			return nil
		}
		t, err := f.Type.EncodeType(tdt)
		if err != nil {
			return nil
		}
		vs = append(vs, concat(id, t)...)
	}

	tdt.Add(v, concat(id, l, vs))
	return nil
}

func (v Variant) Decode(r *bytes.Reader) (interface{}, error) {
	id, err := leb128.DecodeUnsigned(r)
	if err != nil {
		return nil, err
	}
	if id.Cmp(big.NewInt(int64(len(v.fields)))) >= 0 {
		return nil, fmt.Errorf("invalid variant index: %v", id)
	}
	v_, err := v.fields[int(id.Int64())].Type.Decode(r)
	if err != nil {
		return nil, err
	}
	return FieldValue{
		Name:  id.String(),
		Value: v_,
	}, nil
}

func (v Variant) EncodeType(tdt *TypeDefinitionTable) ([]byte, error) {
	idx, ok := tdt.Indexes[v.String()]
	if !ok {
		return nil, fmt.Errorf("missing type index for: %s", v)
	}
	return leb128.EncodeSigned(big.NewInt(int64(idx)))
}

func (v Variant) EncodeValue(value interface{}) ([]byte, error) {
	fs, ok := value.(FieldValue)
	if !ok {
		return nil, fmt.Errorf("invalid argument: %v", v)
	}
	for i, f := range v.fields {
		if f.Name == fs.Name {
			id, err := leb128.EncodeUnsigned(big.NewInt(int64(i)))
			if err != nil {
				return nil, err
			}
			v_, err := f.Type.EncodeValue(fs.Value)
			if err != nil {
				return nil, err
			}
			return concat(id, v_), nil
		}
	}
	return nil, fmt.Errorf("unknown variant: %v", value)
}

func (v Variant) String() string {
	var s []string
	for _, f := range v.fields {
		s = append(s, fmt.Sprintf("%s:%s", f.Name, f.Type.String()))
	}
	return fmt.Sprintf("variant {%s}", strings.Join(s, "; "))
}
