package idl_test

import (
	"testing"

	"github.com/allusion-be/candid-go/idl"
)

func TestFloat(t *testing.T) {
	test(t, "4449444c000173000000bf", idl.NewFloat32(-0.5))
	test(t, "4449444c00017300000000", idl.NewFloat32(0))
	test(t, "4449444c0001730000003f", idl.NewFloat32(0.5))
	test(t, "4449444c00017300004040", idl.NewFloat32(3))

	test(t, "4449444c000172000000000000e0bf", idl.NewFloat64(-0.5))
	test(t, "4449444c0001720000000000000000", idl.NewFloat64(0))
	test(t, "4449444c000172000000000000e03f", idl.NewFloat64(0.5))
	test(t, "4449444c0001720000000000000840", idl.NewFloat64(3))
}
