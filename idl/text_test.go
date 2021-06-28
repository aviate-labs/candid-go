package idl_test

import (
	"testing"

	"github.com/allusion-be/candid-go/idl"
)

func TestText(t *testing.T) {
	test(t, "4449444c00017100", idl.NewText(""))
	test(t, "4449444c000171064d6f746f6b6f", idl.NewText("Motoko"))
}
