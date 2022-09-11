package idl_test

import (
	"math/big"

	"github.com/aviate-labs/candid-go/idl"
)

func ExampleRecordType() {
	test([]idl.Type{idl.NewRecordType(nil)}, []interface{}{nil})
	test_([]idl.Type{idl.NewRecordType(map[string]idl.Type{
		"foo": new(idl.TextType),
		"bar": new(idl.IntType),
	})}, []interface{}{
		map[string]interface{}{
			"foo": "💩",
			"bar": big.NewInt(42),
			"baz": big.NewInt(0),
		},
	})
	// Output:
	// 4449444c016c000100
	// 4449444c016c02d3e3aa027c868eb7027101002a04f09f92a9
}
