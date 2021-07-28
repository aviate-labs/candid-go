package idl_test

import (
	"github.com/allusion-be/candid-go/idl"
)

func ExampleBool() {
	test([]idl.Type{new(idl.Bool)}, []interface{}{true})
	test([]idl.Type{new(idl.Bool)}, []interface{}{false})
	test([]idl.Type{new(idl.Bool)}, []interface{}{0})
	test([]idl.Type{new(idl.Bool)}, []interface{}{"false"})
	// Output:
	// 4449444c00017e01
	// 4449444c00017e00
	// invalid argument: 0
	// invalid argument: false
}
