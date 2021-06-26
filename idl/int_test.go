package idl_test

import (
	"math/big"
	"testing"

	"github.com/allusion-be/candid-go/idl"
)

func TestInt(t *testing.T) {
	test(t, "4449444c00017c00", idl.NewInt(big.NewInt(0)))
	test(t, "4449444c00017c2a", idl.NewInt(big.NewInt(42)))
	test(t, "4449444c00017cd285d8cc04", idl.NewInt(big.NewInt(1234567890)))
	test(t, "4449444c00017caefaa7b37b", idl.NewInt(big.NewInt(-1234567890)))
	test(t, "4449444c00017c808098f4e9b5caea00", idl.NewInt(newInt("60000000000000000")))

	t.Run("8", func(t *testing.T) {
		test(t, "4449444c00017780", idl.NewInt8(-128))
		test(t, "4449444c000177d6", idl.NewInt8(-42))
		test(t, "4449444c000177ff", idl.NewInt8(-1))
		test(t, "4449444c00017700", idl.NewInt8(0))
		test(t, "4449444c00017701", idl.NewInt8(1))
		test(t, "4449444c0001772a", idl.NewInt8(42))
		test(t, "4449444c0001777f", idl.NewInt8(127))
	})

	t.Run("32", func(t *testing.T) {
		test(t, "4449444c0001752efd69b6", idl.NewInt32(-1234567890))
		test(t, "4449444c000175d6ffffff", idl.NewInt32(-42))
		test(t, "4449444c0001752a000000", idl.NewInt32(42))
		test(t, "4449444c000175d2029649", idl.NewInt32(1234567890))
	})
}
