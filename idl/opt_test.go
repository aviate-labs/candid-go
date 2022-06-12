package idl_test

import (
	"math/big"

	"github.com/aviate-labs/candid-go/idl"
)

func ExampleOpt() {
	var optNat *idl.Opt[*idl.Nat] = idl.NewOpt(new(idl.Nat))
	test([]idl.Type{optNat}, []interface{}{nil})
	test([]idl.Type{optNat}, []interface{}{big.NewInt(1)})
	// Output:
	// 4449444c016e7d010000
	// 4449444c016e7d01000101
}
