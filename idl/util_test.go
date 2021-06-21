package idl_test

import (
	"encoding/hex"
	"math/big"
	"reflect"
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
	ts_, err := idl.Decode(bs)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(ts, ts_) {
		t.Error(ts_)
	}
}

func newInt(s string) *big.Int {
	bi, _ := new(big.Int).SetString(s, 10)
	return bi
}
