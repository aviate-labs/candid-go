package idl_test

import (
	"testing"

	"github.com/allusion-be/candid-go/idl"
)

func TestBool(t *testing.T) {
	test(t, "4449444c00017e01", idl.NewBool(true))
	test(t, "4449444c00017e00", idl.NewBool(false))
}
