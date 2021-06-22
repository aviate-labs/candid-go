package idl_test

import (
	"math/big"
	"testing"

	"github.com/allusion-be/candid-go/idl"
)

func TestNat(t *testing.T) {
	test(t, []idl.Type{idl.NewNat(big.NewInt(0))}, "4449444c00017d00")
	test(t, []idl.Type{idl.NewNat(big.NewInt(42))}, "4449444c00017d2a")
	test(t, []idl.Type{idl.NewNat(big.NewInt(1234567890))}, "4449444c00017dd285d8cc04")
	test(t, []idl.Type{idl.NewNat(newInt("60000000000000000"))}, "4449444c00017d808098f4e9b5ca6a")
}

func TestNat8(t *testing.T) {
	test(t, []idl.Type{idl.NewNat8(0)}, "4449444c00017b00")
	test(t, []idl.Type{idl.NewNat8(42)}, "4449444c00017b2a")
	test(t, []idl.Type{idl.NewNat8(255)}, "4449444c00017bff")

	test(t, []idl.Type{idl.NewNat16(0)}, "4449444c00017a0000")
	test(t, []idl.Type{idl.NewNat16(42)}, "4449444c00017a2a00")
	test(t, []idl.Type{idl.NewNat16(65535)}, "4449444c00017affff")

	test(t, []idl.Type{idl.NewNat32(0)}, "4449444c00017900000000")
	test(t, []idl.Type{idl.NewNat32(42)}, "4449444c0001792a000000")
	test(t, []idl.Type{idl.NewNat32(4294967295)}, "4449444c000179ffffffff")

	test(t, []idl.Type{idl.NewNat64(0)}, "4449444c0001780000000000000000")
	test(t, []idl.Type{idl.NewNat64(42)}, "4449444c0001782a00000000000000")
	test(t, []idl.Type{idl.NewNat64(1234567890)}, "4449444c000178d202964900000000")
}
