package typ_test

import (
	"fmt"

	"github.com/aviate-labs/candid-go/typ"
)

func ExampleNat() {
	fmt.Println(typ.NewNat[uint](0))
	fmt.Println(typ.NewNatFromString("123456789876543210"))
	// Output:
	// 0
	// 123456789876543210
}

func ExampleInt() {
	fmt.Println(typ.NewInt(0))
	fmt.Println(typ.NewIntFromString("-123456789876543210"))
	// Output:
	// 0
	// -123456789876543210
}
