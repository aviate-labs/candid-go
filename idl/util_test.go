package idl_test

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	"github.com/allusion-be/candid-go/idl"
)

func newInt(s string) *big.Int {
	bi, _ := new(big.Int).SetString(s, 10)
	return bi
}

func strEqual(a, b interface{}) bool {
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", b)
}

func test(t *testing.T, x string, ts ...idl.Type) {
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
	if ts := idl.Tuple(ts); !strEqual(ts, ts_) {
		t.Errorf("%v, %v", ts, ts_)
	}
}
