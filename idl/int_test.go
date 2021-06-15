package idl_test

import (
	"math/big"
	"testing"

	"github.com/allusion-be/candid-go/idl"
)

func TestInt(t *testing.T) {
	test(t, []idl.Type{idl.Int(*big.NewInt(0))}, "4449444c00017c00")
	test(t, []idl.Type{idl.Int(*big.NewInt(42))}, "4449444c00017c2a")
	test(t, []idl.Type{idl.Int(*big.NewInt(1234567890))}, "4449444c00017cd285d8cc04")
	test(t, []idl.Type{idl.Int(*big.NewInt(-1234567890))}, "4449444c00017caefaa7b37b")
	test(t, []idl.Type{idl.Int(*newInt("60000000000000000"))}, "4449444c00017c808098f4e9b5caea00")
}
