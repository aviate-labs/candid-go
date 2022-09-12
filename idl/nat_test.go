package idl_test

import (
	"math/big"

	"github.com/aviate-labs/candid-go/idl"
)

func ExampleNat16Type() {
	test([]idl.Type{idl.Nat16Type()}, []interface{}{uint16(0)})
	test([]idl.Type{idl.Nat16Type()}, []interface{}{uint16(42)})
	test([]idl.Type{idl.Nat16Type()}, []interface{}{uint16(65535)})
	test([]idl.Type{idl.Nat16Type()}, []interface{}{uint32(65536)})
	// Output:
	// 4449444c00017a0000
	// 4449444c00017a2a00
	// 4449444c00017affff
	// enc: invalid value: 65536
}

func ExampleNat32Type() {
	test([]idl.Type{idl.Nat32Type()}, []interface{}{uint32(0)})
	test([]idl.Type{idl.Nat32Type()}, []interface{}{uint32(42)})
	test([]idl.Type{idl.Nat32Type()}, []interface{}{uint32(4294967295)})
	test([]idl.Type{idl.Nat32Type()}, []interface{}{uint64(4294967296)})
	// Output:
	// 4449444c00017900000000
	// 4449444c0001792a000000
	// 4449444c000179ffffffff
	// enc: invalid value: 4294967296
}

func ExampleNat64Type() {
	test([]idl.Type{idl.Nat64Type()}, []interface{}{uint64(0)})
	test([]idl.Type{idl.Nat64Type()}, []interface{}{uint64(42)})
	test([]idl.Type{idl.Nat64Type()}, []interface{}{uint64(1234567890)})
	// Output:
	// 4449444c0001780000000000000000
	// 4449444c0001782a00000000000000
	// 4449444c000178d202964900000000
}

func ExampleNat8Type() {
	test([]idl.Type{idl.Nat8Type()}, []interface{}{uint8(0)})
	test([]idl.Type{idl.Nat8Type()}, []interface{}{uint8(42)})
	test([]idl.Type{idl.Nat8Type()}, []interface{}{uint8(255)})
	test([]idl.Type{idl.Nat8Type()}, []interface{}{uint16(256)})
	// Output:
	// 4449444c00017b00
	// 4449444c00017b2a
	// 4449444c00017bff
	// enc: invalid value: 256
}

func ExampleNatType() {
	test([]idl.Type{new(idl.NatType)}, []interface{}{big.NewInt(-1)})
	test([]idl.Type{new(idl.NatType)}, []interface{}{big.NewInt(0)})
	test([]idl.Type{new(idl.NatType)}, []interface{}{big.NewInt(42)})
	test([]idl.Type{new(idl.NatType)}, []interface{}{big.NewInt(1234567890)})
	test([]idl.Type{new(idl.NatType)}, []interface{}{func() *big.Int {
		bi, _ := new(big.Int).SetString("60000000000000000", 10)
		return bi
	}()})
	// Output:
	// enc: can not leb128 encode negative values
	// 4449444c00017d00
	// 4449444c00017d2a
	// 4449444c00017dd285d8cc04
	// 4449444c00017d808098f4e9b5ca6a
}
