package idl_test

import (
	"testing"

	"github.com/allusion-be/candid-go/idl"
)

func TestText(t *testing.T) {
	e := idl.Text("")
	test(t, "4449444c00017100", &e)
	s := idl.Text("Motoko")
	test(t, "4449444c000171064d6f746f6b6f", &s)
}
