package idl_test

import (
	"math/big"

	"github.com/aviate-labs/candid-go/idl"
)

func ExampleVectorType() {
	test([]idl.Type{idl.NewVectorType(new(idl.IntType))}, []interface{}{
		[]interface{}{big.NewInt(0), big.NewInt(1), big.NewInt(2), big.NewInt(3)},
	})
	// Output:
	// 4449444c016d7c01000400010203
}
