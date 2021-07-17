package idl_test

import (
	"fmt"
	"testing"

	"github.com/allusion-be/candid-go/idl"
)

func ExampleDecode() {
	fmt.Println(idl.Decode(append([]byte("DIDL"), 0x00, 0x00)))
	// Output:
	// () <nil>
}

func TestInvalid(t *testing.T) {
	for i, test := range []struct {
		tuple idl.Tuple
		data  []byte
		desc  string
	}{
		{
			desc: "empty",
		},
		{
			data: []byte{0x00, 0x00},
			desc: "no magic bytes",
		},
		{
			data: []byte{'D', 'A', 'D', 'L'},
			desc: "wrong magic bytes",
		},
		{
			data: []byte{'D', 'A', 'D', 'L', 0x00, 0x00},
			desc: "wrong magic bytes",
		},
		{
			data: []byte{'D', 'I', 'D', 'L', 0x00, 0x00, 0x00},
			desc: "too long",
		},
		{
			data: []byte{'D', 'I', 'D', 'L', 0x00, 0x01, 0x6e},
			desc: "type: not primitive",
		},
		{
			data: []byte{'D', 'I', 'D', 'L', 0x00, 0x01, 0x5e},
			desc: "type: out of range",
		},
		{
			data: []byte{'D', 'I', 'D', 'L', 0x00, 0x01, 0x7e},
			desc: "end of data",
		},
	} {
		ts, err := idl.Decode(test.data)
		switch err := err.(type) {
		case *idl.FormatError:
			if err.Description != test.desc {
				t.Fatalf("(%d) expected: %s, got %s", i, test.desc, err)
			}
		default:
			t.Fatalf("(%d) %s", i, err)
		}

		if test.tuple.String() != ts.String() {
			t.Error(ts)
		}
	}
}
