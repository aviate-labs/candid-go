package idl_test

import (
	"math/big"
	"testing"

	"github.com/allusion-be/candid-go/idl"
)

func TestInt(t *testing.T) {
	v0 := idl.Int(*big.NewInt(0))
	test(t, []idl.Type{&v0}, "4449444c00017c00")
	v42 := idl.Int(*big.NewInt(42))
	test(t, []idl.Type{&v42}, "4449444c00017c2a")
	v123 := idl.Int(*big.NewInt(1234567890))
	test(t, []idl.Type{&v123}, "4449444c00017cd285d8cc04")
	v123n := idl.Int(*big.NewInt(-1234567890))
	test(t, []idl.Type{&v123n}, "4449444c00017caefaa7b37b")
	v600 := idl.Int(*newInt("60000000000000000"))
	test(t, []idl.Type{&v600}, "4449444c00017c808098f4e9b5caea00")
}
