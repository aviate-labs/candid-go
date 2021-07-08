package test_test

import (
	"io/ioutil"
	"testing"

	test "github.com/allusion-be/candid-go/internal/candid_test"
	"github.com/di-wu/parser"
	"github.com/di-wu/parser/ast"
)

func TestData(t *testing.T) {
	rawDid, _ := ioutil.ReadFile("testdata/prim.test.did")
	p, _ := ast.New(rawDid)
	if _, err := test.TestData(p); err != nil {
		t.Fatal(err)
	}
	if _, err := p.Expect(parser.EOD); err != nil {
		t.Error(err)
	}
}
