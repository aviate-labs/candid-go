package idl

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/allusion-be/leb128"
)

var pow2_32 = big.NewInt(4294967296) // 2^32

type TypeTable struct {
	types   [][]byte
	indexes map[string]int
}

func NewTable(r *bytes.Reader) (TypeTable, error) {
	n, err := leb128.DecodeUnsigned(r)
	if err != nil {
		return TypeTable{}, err
	}
	var tt [][]interface{}
	for i := 0; i < int(n.Int64()); i++ {
		typ, err := leb128.DecodeSigned(r)
		if err != nil {
			return TypeTable{}, err
		}
		switch typ.Int64() {
		case optType, vecType:
			v, err := leb128.DecodeSigned(r)
			if err != nil {
				return TypeTable{}, err
			}
			tt = append(tt, []interface{}{typ, v})
		case recordType, variantType:
			len, err := leb128.DecodeUnsigned(r)
			if err != nil {
				return TypeTable{}, err
			}
			var (
				prev   *big.Int
				fields [][]interface{}
			)
			for i := 0; i < int(len.Int64()); i++ {
				h, err := leb128.DecodeUnsigned(r)
				if err != nil {
					return TypeTable{}, err
				}
				if h.Cmp(pow2_32) <= 0 {
					return TypeTable{}, fmt.Errorf("field id out of range")
				}
				if h.Cmp(prev) <= 0 {
					return TypeTable{}, fmt.Errorf("field collision or not sorted")
				}
				prev = h
				v, err := leb128.DecodeSigned(r)
				if err != nil {
					return TypeTable{}, err
				}
				fields = append(fields, []interface{}{h, v})
			}
			tt = append(tt, []interface{}{typ, fields})
		case funcType:
			for i := 0; i < 2; i++ {
				len, err := leb128.DecodeUnsigned(r)
				if err != nil {
					return TypeTable{}, err
				}
				for i := 0; i < int(len.Int64()); i++ {
					if _, err = leb128.DecodeSigned(r); err != nil {
						return TypeTable{}, err
					}
				}
			}
			len, err := leb128.DecodeUnsigned(r)
			if err != nil {
				return TypeTable{}, err
			}
			ann := make([]byte, len.Int64())
			if _, err := r.Read(ann); err != nil {
				return TypeTable{}, nil
			}
			tt = append(tt, []interface{}{typ, nil})
		case serviceType:
			len, err := leb128.DecodeUnsigned(r)
			if err != nil {
				return TypeTable{}, err
			}
			for i := 0; i < int(len.Int64()); i++ {
				len, err := leb128.DecodeUnsigned(r)
				if err != nil {
					return TypeTable{}, err
				}
				name := make([]byte, len.Int64())
				if _, err := r.Read(name); err != nil {
					return TypeTable{}, nil
				}
				if _, err = leb128.DecodeSigned(r); err != nil {
					return TypeTable{}, err
				}
			}
			tt = append(tt, []interface{}{typ, nil})
		}
	}
	return TypeTable{}, nil
}

func (table *TypeTable) Add(t Type, bs []byte) {
	i := len(table.types)
	table.indexes[t.Name()] = i
	table.types = append(table.types, bs)
}

func (table TypeTable) Encode() ([]byte, error) {
	bs, err := leb128.EncodeUnsigned(big.NewInt(int64(len(table.types))))
	if err != nil {
		return nil, err
	}
	for _, t := range table.types {
		bs = append(bs, t...)
	}
	return bs, nil
}

func (table TypeTable) Has(t Type) bool {
	_, ok := table.indexes[t.Name()]
	return ok
}

func (table TypeTable) Index(name string) int {
	if i, ok := table.indexes[name]; ok {
		return i
	}
	return -1
}
