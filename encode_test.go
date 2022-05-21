package candid_test

import (
	"fmt"
	"math/big"

	"github.com/aviate-labs/candid-go"
	"github.com/aviate-labs/candid-go/idl"
	"github.com/aviate-labs/candid-go/typ"
)

func ExampleMarshalBool() {
	fmt.Println(idl.Encode([]idl.Type{new(idl.Bool)}, []interface{}{true}))
	fmt.Println(candid.EncodeValue("(true)"))
	fmt.Println(candid.Marshal([]interface{}{true}))
	// Output:
	// [68 73 68 76 0 1 126 1] <nil>
	// [68 73 68 76 0 1 126 1] <nil>
	// [68 73 68 76 0 1 126 1] <nil>
}

func ExampleMarshalNat() {
	fmt.Println(idl.Encode([]idl.Type{new(idl.Nat)}, []interface{}{big.NewInt(5)}))
	fmt.Println(candid.EncodeValue("(5 : nat)"))
	fmt.Println(candid.Marshal([]interface{}{typ.NewNat[uint](5)}))
	// Output:
	// [68 73 68 76 0 1 125 5] <nil>
	// [68 73 68 76 0 1 125 5] <nil>
	// [68 73 68 76 0 1 125 5] <nil>
}