package idl

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/allusion-be/leb128"
)

type Int big.Int

func (Int) Name() string {
	return "int"
}

func (Int) Encode() []byte {
	bs, _ := leb128.EncodeSigned(intType)
	return bs
}

func (n Int) EncodeValue() []byte {
	bi := big.Int(n)
	bs, _ := leb128.EncodeSigned(&bi)
	return bs
}

func (n *Int) Decode(r *bytes.Reader) error {
	bi, err := leb128.DecodeSigned(r)
	if err != nil {
		return err
	}
	*n = Int(*bi)
	return nil
}

func (Int) BuildTypeTable(*TypeTable) {}

func (i Int) String() string {
	return fmt.Sprintf("int: %v", big.Int(i))
}
