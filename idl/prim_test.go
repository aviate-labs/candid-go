package idl_test

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/allusion-be/candid-go/idl"
	"github.com/allusion-be/candid-go/internal/candidtest"
	"github.com/di-wu/parser"
	"github.com/di-wu/parser/ast"

	"testing"
)

func TestPrim(t *testing.T) {
	raw, err := ioutil.ReadFile("./testdata/prim.test.did")
	if err != nil {
		t.Fatal(err)
	}
	p, err := ast.New(raw)
	if err != nil {
		t.Fatal(err)
	}
	n, err := candidtest.TestData(p)
	if err != nil {
		t.Fatal(err)
	}
	for _, n := range n.Children() {
		switch n.Type {
		case candidtest.CommentTextT: // ignore
		case candidtest.TestT:
			var bs []byte
			for i, n := range n.Children() {
				switch i {
				case 0:
					switch n.Type {
					case candidtest.BlobInputT:
						bs = toBin(n.Value)
					default:
						t.Fatal(n)
					}
				case 1:
					tup, dErr := idl.Decode(bs)
					switch n.Type {
					case candidtest.TestGoodT:
						if !n.IsParent() {
							p, err := ast.New([]byte(n.Value))
							if err != nil {
								t.Fatal(err)
							}
							ntup, err := candidtest.ValuesBr(p)
							if err != nil {
								t.Fatal(err)
							}
							switch ntup {
							case nil: // ()
								fmt.Println(tup, ntup, n.Value)
								if len(tup) != 0 && tup[0].Name() != "null" && len(tup) != 1 {
									t.Fatal(n, ntup)
								}
							default:
								if dErr != nil {
									t.Fatal(err)
								}
								fmt.Println(n, ntup)
								t.Fatal()
							}
							continue
						}
					case candidtest.TestBadT:
					case candidtest.TestTestT:
					default:
						t.Fatal(n)
					}
				default:
					if n.Type != candidtest.DescriptionT {
						t.Fatal(n)
					}
				}
			}
		default:
			t.Fatal(n)
		}
	}
	if _, err := p.Expect(parser.EOD); err != nil {
		t.Fatal(err)
	}
}

func toBin(v string) []byte {
	var (
		vs = strings.Split(v, "\\")
		bs = []byte(vs[0])
	)
	for i := 1; i < len(vs); i++ {
		h, _ := hex.DecodeString(vs[i])
		bs = append(bs, h...)
	}
	return bs
}
