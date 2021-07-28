package idl_test

import "github.com/allusion-be/candid-go/idl"

func ExampleText() {
	test([]idl.Type{new(idl.Text)}, []interface{}{""})
	test([]idl.Type{new(idl.Text)}, []interface{}{"Motoko"})
	//  Output:
	// 4449444c00017100
	// 4449444c000171064d6f746f6b6f
}
