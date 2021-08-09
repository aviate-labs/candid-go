package idl_test

import (
	"github.com/allusion-be/candid-go/idl"
	"github.com/allusion-be/principal-go"
)

func ExampleFunc() {
	p, _ := principal.Decode("w7x7r-cok77-xa")
	test_(
		[]idl.Type{
			idl.NewFunc(
				[]idl.Type{new(idl.Text)},
				[]idl.Type{new(idl.Nat)},
				nil,
			),
		},
		[]interface{}{
			idl.PrincipalMethod{
				Principal: p,
				Method:    "foo",
			},
		},
	)
	// Output:
	// 4449444c016a0171017d000100010103caffee03666f6f
}
