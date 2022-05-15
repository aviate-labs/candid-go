package candid_test

import (
	"fmt"

	"github.com/aviate-labs/candid-go"
	"github.com/aviate-labs/candid-go/idl"
)

func ExampleMarshal() {
	fmt.Println(idl.Encode([]idl.Type{new(idl.Bool)}, []interface{}{true}))
	fmt.Println(candid.EncodeValue("(true)"))
	fmt.Println(candid.Marshal([]interface{}{true}))
	// Output:
	// [68 73 68 76 0 1 126 1] <nil>
	// [68 73 68 76 0 1 126 1] <nil>
	// [68 73 68 76 0 1 126 1] <nil>
}
