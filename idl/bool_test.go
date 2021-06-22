package idl_test

import (
	"testing"

	"github.com/allusion-be/candid-go/idl"
)

func TestBool(t *testing.T) {
	vt := idl.Bool(true)
	test(t, []idl.Type{&vt}, "4449444c00017e01")
	vf := idl.Bool(false)
	test(t, []idl.Type{&vf}, "4449444c00017e00")
}
