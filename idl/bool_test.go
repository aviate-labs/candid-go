package idl_test

import (
	"testing"

	"github.com/allusion-be/candid-go/idl"
)

func TestBool(t *testing.T) {
	test(t, []idl.Type{idl.Bool(true)}, "4449444c00017e01")
	test(t, []idl.Type{idl.Bool(false)}, "4449444c00017e00")
}
