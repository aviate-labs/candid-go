package idl_test

import (
	"testing"

	"github.com/allusion-be/candid-go/idl"
)

func TestNull(t *testing.T) {
	test(t, "4449444c00017f", new(idl.Null))
}
