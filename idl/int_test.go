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
}
