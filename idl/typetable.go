package idl

import (
	"bytes"
	"math/big"

	"github.com/allusion-be/leb128"
)

type TypeTable struct {
	types   [][]byte
	indexes map[string]int
}

func NewTable(r *bytes.Reader) (TypeTable, error) {
	n, err := leb128.DecodeUnsigned(r)
	if err != nil {
		return TypeTable{}, err
	}
	for i := 0; i < int(n.Int64()); i++ {
		_ = i // TODO
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
