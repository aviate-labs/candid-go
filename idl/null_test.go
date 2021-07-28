package idl_test

import (
	"github.com/allusion-be/candid-go/idl"
)

func ExampleNull() {
	test([]idl.Type{new(idl.Null)}, []interface{}{nil})
	// Output:
	// 4449444c00017f
}
