package candid_test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/aviate-labs/candid-go"
)

func ExampleParseDID() {
	raw, _ := ioutil.ReadFile("testdata/counter.did")
	p, _ := candid.ParseDID(raw)
	fmt.Println(p)
	// Output:
	// service : {
	//   inc : () -> nat;
	// }
}

func ExampleEncodeValue() {
	e, _ := candid.EncodeValue("0")
	fmt.Printf("%x\n", e)
	// Output:
	// 4449444c00017c00
}

func TestEncodeValue(t *testing.T) {
	for _, test := range []struct {
		values  string
		encoded string
	}{
		{"opt 0", "4449444c016e7c01000100"},

		{"0", "4449444c00017c00"},
		{"(0)", "4449444c00017c00"},
		{"(0 : nat)", "4449444c00017d00"},
		{"(0 : nat8)", "4449444c00017b00"},
		{"(0 : nat16)", "4449444c00017a0000"},
		{"(0 : nat32)", "4449444c00017900000000"},
		{"(0 : nat64)", "4449444c0001780000000000000000"},
		{"(0 : int)", "4449444c00017c00"},
		{"(0 : int8)", "4449444c00017700"},
		{"(0 : int16)", "4449444c0001760000"},
		{"(0 : int32)", "4449444c00017500000000"},
		{"(0 : int64)", "4449444c0001740000000000000000"},

		{"0.0", "4449444c0001720000000000000000"},
		{"(0 : float32)", "4449444c00017300000000"},
		{"(0.0 : float32)", "4449444c00017300000000"},
		{"(0 : float64)", "4449444c0001720000000000000000"},
		{"(0.0 : float64)", "4449444c0001720000000000000000"},

		{"true", "4449444c00017e01"},
		{"(false : bool)", "4449444c00017e00"},

		{"(null)", "4449444c00017f"},

		{"\"\"", "4449444c00017100"},
		{"\"quint\"", "4449444c000171057175696e74"},
	} {
		e, err := candid.EncodeValue(test.values)
		if err != nil {
			t.Fatal(err)
		}
		if e := fmt.Sprintf("%x", e); e != test.encoded {
			t.Error(test, e)
		}
	}
}
