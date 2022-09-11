package idl_test

import (
	"math/big"

	"github.com/aviate-labs/candid-go/idl"
)

func ExampleOpt() {
	var optNat *idl.OptionalType[*idl.NatType] = idl.NewOptionalType(new(idl.NatType))
	test([]idl.Type{optNat}, []interface{}{nil})
	test([]idl.Type{optNat}, []interface{}{big.NewInt(1)})
	// Output:
	// 4449444c016e7d010000
	// 4449444c016e7d01000101
}
