package candid_test

import (
	"fmt"
	"github.com/di-wu/candid-go"
	"github.com/di-wu/parser"
	"github.com/di-wu/parser/ast"
	"testing"
)

func TestWs(t *testing.T) {
	p, err := ast.New([]byte("\n  \t\n "))
	if err != nil {
		t.Error(err)
		return
	}
	if _, err := candid.Ws(p); err != nil {
		t.Error(err)
	}
	if _, err := p.Expect(parser.EOD); err != nil {
		t.Error(err)
	}
}

func TestName(t *testing.T) {
	for _, name := range []string{
		"addUser", "userName", "userAge", "deleteUser",
	} {
		p, err := ast.New([]byte(name))
		if err != nil {
			t.Error(err)
			return
		}
		if _, err := candid.Name(p); err != nil {
			t.Error(err)
		}
	}
}

func ExampleArgType() {
	p := func(s string) *ast.Parser {
		p, _ := ast.New([]byte(s))
		return p
	}
	fmt.Println(candid.ArgType(p("name : text")))
	fmt.Println(candid.ArgType(p("age : nat8")))
	fmt.Println(candid.ArgType(p("id : nat64")))
	// output:
	// [009] [[019] name, [012] text] <nil>
	// [009] [[019] age, [012] nat8] <nil>
	// [009] [[019] id, [012] nat64] <nil>
}

func ExampleTupType() {
	for _, tuple := range []string{
		"(name : text, age : nat8)",
		"(id : nat64)",
		"()",
	} {
		p, _ := ast.New([]byte(tuple))
		fmt.Println(candid.TupType(p))
	}
	// output:
	// [008] [[009] [[019] name, [012] text], [009] [[019] age, [012] nat8]] <nil>
	// [008] [[009] [[019] id, [012] nat64]] <nil>
	// [008] () <nil>
}

func ExampleMethType() {
	for _, method := range []string{
		"addUser : (name : text, age : nat8) -> (id : nat64)",
		"userName : (id : nat64) -> (text) query",
		"userAge : (id : nat64) -> (nat8) query",
		"deleteUser : (id : nat64) -> () oneway",
	} {
		p, _ := ast.New([]byte(method))
		fmt.Println(candid.MethType(p))
	}
	// output:
	// [005] [[019] addUser, [006] [[008] [[009] [[019] name, [012] text], [009] [[019] age, [012] nat8]], [008] [[009] [[019] id, [012] nat64]]]] <nil>
	// [005] [[019] userName, [006] [[008] [[009] [[019] id, [012] nat64]], [008] [[009] [[012] text]], [007] query]] <nil>
	// [005] [[019] userAge, [006] [[008] [[009] [[019] id, [012] nat64]], [008] [[009] [[012] nat8]], [007] query]] <nil>
	// [005] [[019] deleteUser, [006] [[008] [[009] [[019] id, [012] nat64]], [008] (), [007] oneway]] <nil>
}

func ExampleActorType() {
	var example = `{
	addUser : (name : text, age : nat8) -> (id : nat64);
	userName : (id : nat64) -> (text) query;
	userAge : (id : nat64) -> (nat8) query;
	deleteUser : (id : nat64) -> () oneway;
}`
	p, _ := ast.New([]byte(example))
	actor, _ := candid.ActorType(p)
	fmt.Println(len(actor.Children()))
	// output:
	// 4
}

func ExampleFuncType() {
	for _, function := range []string{
		"(text, text, nat16) -> (text, nat64)",
		"(name : text, address : text, nat16) -> (text, id : nat64)",
		"(name : text, address : text, nr : nat16) -> (nick : text, id : nat64)",
	} {
		p, _ := ast.New([]byte(function))
		fmt.Println(candid.FuncType(p))
	}
	// output:
	// [006] [[008] [[009] [[012] text], [009] [[012] text], [009] [[012] nat16]], [008] [[009] [[012] text], [009] [[012] nat64]]] <nil>
	// [006] [[008] [[009] [[019] name, [012] text], [009] [[019] address, [012] text], [009] [[012] nat16]], [008] [[009] [[012] text], [009] [[019] id, [012] nat64]]] <nil>
	// [006] [[008] [[009] [[019] name, [012] text], [009] [[019] address, [012] text], [009] [[019] nr, [012] nat16]], [008] [[009] [[019] nick, [012] text], [009] [[019] id, [012] nat64]]] <nil>
}

func ExampleConsType() {
	for _, record := range []string{
		"record {\n  num : nat;\n}",
		"record { nat; nat }",
		"record { 0 : nat; 1 : nat }",
	} {
		p, _ := ast.New([]byte(record))
		fmt.Println(candid.ConsType(p))
	}
	// output:
	// [016] [[010] [[019] num, [012] nat]] <nil>
	// [016] [[010] [[012] nat], [010] [[012] nat]] <nil>
	// [016] [[010] [[021] 0, [012] nat], [010] [[021] 1, [012] nat]] <nil>
}

func ExampleDef() {
	for _, def := range []string{
		"type list = opt node",
		"type color = variant { red; green; blue }",
		"type tree = variant {\n  leaf : int;\n  branch : record {left : tree; val : int; right : tree};\n}",
		"type stream = opt record {head : nat; next : func () -> stream}",
	} {
		p, _ := ast.New([]byte(def))
		fmt.Println(candid.Def(p))
	}
	// output:
	// [002] [[019] list, [014] [[019] node]] <nil>
	// [002] [[019] color, [017] [[010] [[019] red], [010] [[019] green], [010] [[019] blue]]] <nil>
	// [002] [[019] tree, [017] [[010] [[019] leaf, [012] int], [010] [[019] branch, [016] [[010] [[019] left, [019] tree], [010] [[019] val, [012] int], [010] [[019] right, [019] tree]]]]] <nil>
	// [002] [[019] stream, [014] [[016] [[010] [[019] head, [012] nat], [010] [[019] next, [018] [[006] [[008] (), [009] [[019] stream]]]]]]] <nil>
}

func TestDef_service(t *testing.T) {
	var example = `type broker = service {
  findCounterService : (name : text) ->
    (service {up : () -> (); current : () -> nat});
}`
	p, err := ast.New([]byte(example))
	if err != nil {
		t.Error(err)
		return
	}
	if _, err := candid.Def(p); err != nil {
		t.Error(err)
		return
	}
	if _, err := p.Expect(parser.EOD); err != nil {
		t.Error(err)
	}
}

func TestDef_function(t *testing.T) {
	var example = `type engine = service {
  search : (query : text, callback : func (vec result) -> ());
}`
	p, err := ast.New([]byte(example))
	if err != nil {
		t.Error(err)
		return
	}
	if _, err := candid.Def(p); err != nil {
		t.Error(err)
		return
	}
	if _, err := p.Expect(parser.EOD); err != nil {
		t.Error(err)
	}
}
