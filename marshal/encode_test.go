package marshal_test

import (
	"fmt"
	"math/big"

	"github.com/aviate-labs/candid-go"
	"github.com/aviate-labs/candid-go/idl"
	"github.com/aviate-labs/candid-go/marshal"
	"github.com/aviate-labs/candid-go/typ"
	"github.com/aviate-labs/principal-go"
)

func ExampleMarshal_bool() {
	fmt.Println(idl.Encode([]idl.Type{new(idl.BoolType)}, []any{true}))
	fmt.Println(candid.EncodeValue("(true)"))
	fmt.Println(marshal.Marshal([]any{true}))
	// Output:
	// [68 73 68 76 0 1 126 1] <nil>
	// [68 73 68 76 0 1 126 1] <nil>
	// [68 73 68 76 0 1 126 1] <nil>
}

func ExampleMarshal_nat() {
	fmt.Println(idl.Encode([]idl.Type{new(idl.NatType)}, []any{big.NewInt(5)}))
	fmt.Println(candid.EncodeValue("(5 : nat)"))
	fmt.Println(marshal.Marshal([]any{typ.NewNat[uint](5)}))
	// Output:
	// [68 73 68 76 0 1 125 5] <nil>
	// [68 73 68 76 0 1 125 5] <nil>
	// [68 73 68 76 0 1 125 5] <nil>
}

func ExampleMarshal_principal() {
	p, _ := principal.Decode("aaaaa-aa")
	fmt.Println(marshal.Marshal([]any{&p}))
	fmt.Println(marshal.Marshal([]any{p}))
	// Output:
	// [68 73 68 76 0 1 104 1 0] <nil>
	// [68 73 68 76 0 1 104 1 0] <nil>
}

func ExampleMarshal_reserved() {
	fmt.Println(marshal.Marshal([]any{new(typ.Reserved)}))
	// Output:
	// [68 73 68 76 0 1 112] <nil>
}

func ExampleMarshal_null() {
	fmt.Println(marshal.Marshal([]any{new(typ.Null)}))
	// Output:
	// [68 73 68 76 0 1 127] <nil>
}

func ExampleMarshal_empty() {
	fmt.Println(marshal.Marshal([]any{new(typ.Empty)}))
	// Output:
	// [68 73 68 76 0 1 111] <nil>
}
