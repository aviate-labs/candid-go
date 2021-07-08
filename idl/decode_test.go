package idl_test

import (
	"github.com/allusion-be/candid-go/idl"
	"testing"
)

func TestInvalid(t *testing.T) {
	for _, test := range []struct {
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
	} {
		ts, err := idl.Decode(test.data)
		switch err := err.(type) {
		case *idl.FormatError:
			if err.Description != test.desc {
				t.Fatal(err)
			}
		default:
			t.Fatal(err)
		}

		if test.tuple.String() != ts.String() {
			t.Error(ts)
		}
	}
}
