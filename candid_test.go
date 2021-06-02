package candid

import (
	"fmt"
	"testing"
)

var complexTypes = []byte(`type address = record {
  street : text;
  city : text;
  zip_code : nat;
  country : text;
};
service address_book : {
  set_address : (name : text, addr : address) -> ();
  get_address : (name : text) -> (opt address) query;
};`)

func TestParseDID(t *testing.T) {
	p, err := ParseDID(complexTypes)
	if err != nil {
		t.Fatal(err)
	}
	if len(p.definitions) != 1 {
		t.Fatal(p.definitions)
	}
	address := p.definitions[0].(Type)
	if address.Id != "address" {
		t.Error(address.Id)
	}
	rec := address.Data.(Record)
	if len(rec) != 4 {
		t.Error(rec)
	}
	if len(p.actors) != 1 {
		t.Fatal(p.actors)
	}
	book := p.actors[0]
	if book.Id == nil || *book.Id != "address_book" {
		t.Error(book.Id)
	}
	if len(book.Methods) != 2 {
		t.Error(book.Methods)
	}
	if p.String() != string(complexTypes) {
		fmt.Println(p)
		t.Error(p)
	}
}
