package idl

import "github.com/allusion-be/leb128"

type Null struct{}

func (Null) Name() string {
	return "null"
}

func (Null) Encode() []byte {
	bs, _ := leb128.EncodeSigned(nullType)
	return bs
}

func (Null) EncodeValue(v interface{}) []byte {
	return []byte{}
}

func (Null) BuildTypeTable(*TypeTable) {}

func (Null) Covariant(v interface{}) bool {
	return v == nil
}
