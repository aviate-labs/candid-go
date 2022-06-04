package idl_test

import (
	"github.com/aviate-labs/candid-go/idl"
	"github.com/aviate-labs/principal-go"
)

func ExamplePrincipal() {
	p, _ := principal.Decode("aaaaa-aa")
	test([]idl.Type{idl.NewOpt(new(idl.Principal))}, []interface{}{p})
	// Output:
	// 4449444c016e680100010100
}
