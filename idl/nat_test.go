package idl_test

import (
	"math/big"
	"testing"

	"github.com/allusion-be/candid-go/idl"
)

func TestNat(t *testing.T) {
	test(t, []idl.Type{idl.Nat(*big.NewInt(0))}, "4449444c00017d00")
	test(t, []idl.Type{idl.Nat(*big.NewInt(42))}, "4449444c00017d2a")
	test(t, []idl.Type{idl.Nat(*big.NewInt(1234567890))}, "4449444c00017dd285d8cc04")
	test(t, []idl.Type{idl.Nat(*newInt("60000000000000000"))}, "4449444c00017d808098f4e9b5ca6a")
}
