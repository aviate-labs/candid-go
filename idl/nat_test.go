package idl_test

import (
	"math/big"
	"testing"

	"github.com/allusion-be/candid-go/idl"
)

func TestNat(t *testing.T) {
	test(t, "4449444c00017d00", idl.NewNat(big.NewInt(0)))
	test(t, "4449444c00017d2a", idl.NewNat(big.NewInt(42)))
	test(t, "4449444c00017dd285d8cc04", idl.NewNat(big.NewInt(1234567890)))
	test(t, "4449444c00017d808098f4e9b5ca6a", idl.NewNat(newInt("60000000000000000")))

	t.Run("8", func(t *testing.T) {
		test(t, "4449444c00017b00", idl.NewNat8(0))
		test(t, "4449444c00017b2a", idl.NewNat8(42))
		test(t, "4449444c00017bff", idl.NewNat8(255))
	})

	t.Run("16", func(t *testing.T) {
		test(t, "4449444c00017a0000", idl.NewNat16(0))
		test(t, "4449444c00017a2a00", idl.NewNat16(42))
		test(t, "4449444c00017affff", idl.NewNat16(65535))
	})

	t.Run("32", func(t *testing.T) {
		test(t, "4449444c00017900000000", idl.NewNat32(0))
		test(t, "4449444c0001792a000000", idl.NewNat32(42))
		test(t, "4449444c000179ffffffff", idl.NewNat32(4294967295))
	})

	t.Run("64", func(t *testing.T) {
		test(t, "4449444c0001780000000000000000", idl.NewNat64(0))
		test(t, "4449444c0001782a00000000000000", idl.NewNat64(42))
		test(t, "4449444c000178d202964900000000", idl.NewNat64(1234567890))
	})

}
