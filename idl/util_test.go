package idl_test

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/allusion-be/candid-go/idl"
)

func test(t *testing.T, ts []idl.Type, x string) {
	bs, err := idl.Encode(ts)
	if err != nil {
		t.Fatal(err)
	}
	if h := hex.EncodeToString(bs); h != x {
		t.Error(x, h)
	}
}

func newInt(s string) *big.Int {
	bi, _ := new(big.Int).SetString(s, 10)
	return bi
}
