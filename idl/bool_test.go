package idl_test

import (
	"testing"

	"github.com/allusion-be/candid-go/idl"
)

func TestBool(t *testing.T) {
	vt := idl.Bool(true)
	test(t, "4449444c00017e01", &vt)
	vf := idl.Bool(false)
	test(t, "4449444c00017e00", &vf)
}
