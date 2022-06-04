package idl_test

import (
	"fmt"

	"github.com/aviate-labs/candid-go/idl"
)

func ExampleTokens() {
	fmt.Println(idl.NewRec(map[string]idl.Type{
		"e8s": idl.Nat64(),
	}))
	// Output:
	// record {e8s:nat64}
}

func ExampleLedger() {
	var service = idl.NewInterface(func(typ idl.IDL) *idl.Service {
		accountIdentitier := typ.Vec(typ.Nat8)
		accountBalanceArgs := typ.Record(map[string]idl.Type{
			"account": accountIdentitier,
		})
		tokens := idl.NewRec(map[string]idl.Type{
			"e8s": idl.Nat64(),
		})
		return typ.Service(map[string]*idl.Func{
			"account_balance": typ.Func([]idl.Type{accountBalanceArgs}, []idl.Type{tokens}, []string{"query"}),
			// etc.
		})
	})
	fmt.Println(service)
	// Output:
	// service {account_balance:(record {account:vec nat8}) -> (record {e8s:nat64}) query}
}
