package marshal_test

import (
	"testing"

	"github.com/aviate-labs/candid-go/marshal"
	"github.com/aviate-labs/candid-go/typ"
)

func TestUnmarshal_nat(t *testing.T) {
	data, err := marshal.Marshal([]any{typ.NewNat[uint](5)})
	if err != nil {
		t.Fatal(err)
	}
	var num typ.Nat
	if err := marshal.Unmarshal(data, []any{&num}); err != nil {
		t.Fatal(err)
	}
	if num.BigInt().Uint64() != 5 {
		t.Errorf("unexpected num: %s", num)
	}

	{ // uint8
		data, err := marshal.Marshal([]any{uint8(5)})
		if err != nil {
			t.Fatal(err)
		}
		var num uint8
		if err := marshal.Unmarshal(data, []any{&num}); err != nil {
			t.Fatal(err)
		}
		if num != 5 {
			t.Errorf("unexpected num: %d", num)
		}
	}
}

func TestUnmarshal_string_valid(t *testing.T) {
	data, err := marshal.Marshal([]any{"John"})
	if err != nil {
		t.Fatal(err)
	}
	var name string
	if err := marshal.Unmarshal(data, []any{name}); err == nil {
		t.Fatal()
	}
	if err := marshal.Unmarshal(data, []any{&name}); err != nil {
		t.Fatal(err)
	}
	if name != "John" {
		t.Errorf("unexpected name: %q", name)
	}

	{ // Multiple strings.
		data, err := marshal.Marshal([]any{"John", "Doe"})
		if err != nil {
			t.Fatal(err)
		}
		var firstName string
		var lastName string
		if err := marshal.Unmarshal(data, []any{&firstName, &lastName}); err != nil {
			t.Fatal(err)
		}
		if firstName != "John" {
			t.Errorf("unexpected first name: %q", firstName)
		}
		if lastName != "Doe" {
			t.Errorf("unexpected last name: %q", lastName)
		}
	}
}

func TestUnmarshal_string_invalid(t *testing.T) {
	data, err := marshal.Marshal([]any{true})
	if err != nil {
		t.Fatal(err)
	}
	var name string
	if err := marshal.Unmarshal(data, []any{&name}); err == nil {
		t.Fatal()
	}
}
