package idl_test

import (
	"encoding/hex"
	"testing"

	"github.com/allusion-be/candid-go/idl"
)

func test(t *testing.T, ts []idl.Type, vs []interface{}, x string) {
	bs, err := idl.Encode(ts, vs)
	if err != nil {
		t.Fatal(err)
	}
	if h := hex.EncodeToString(bs); h != x {
		t.Error(x, h)
	}
}
