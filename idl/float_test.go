package idl_test

import (
	"github.com/aviate-labs/candid-go/idl"
)

func ExampleFloat32() {
	test([]idl.Type{idl.Float32()}, []interface{}{float32(-0.5)})
	test([]idl.Type{idl.Float32()}, []interface{}{float32(0)})
	test([]idl.Type{idl.Float32()}, []interface{}{float32(0.5)})
	test([]idl.Type{idl.Float32()}, []interface{}{float32(3)})
	// Output:
	// 4449444c000173000000bf
	// 4449444c00017300000000
	// 4449444c0001730000003f
	// 4449444c00017300004040
}

func ExampleFloat64() {
	test([]idl.Type{idl.Float64()}, []interface{}{-0.5})
	test([]idl.Type{idl.Float64()}, []interface{}{float32(0)})
	test([]idl.Type{idl.Float64()}, []interface{}{0.5})
	test([]idl.Type{idl.Float64()}, []interface{}{float64(3)})
	// Output:
	// 4449444c000172000000000000e0bf
	// 4449444c0001720000000000000000
	// 4449444c000172000000000000e03f
	// 4449444c0001720000000000000840
}
