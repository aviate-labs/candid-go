package idl_test

import (
	"testing"

	"github.com/allusion-be/candid-go/idl"
)

func TestNull(t *testing.T) {
	test(t, []idl.Type{idl.Null{}}, "4449444c00017f")
}
