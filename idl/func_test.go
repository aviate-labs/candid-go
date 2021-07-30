package idl_test

import (
	"github.com/allusion-be/agent-go"
	"github.com/allusion-be/candid-go/idl"
)

func ExampleFunc() {
	p, _ := agent.DecodePrincipal("w7x7r-cok77-xa")
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
