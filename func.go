package candid

import (
	"fmt"
	"github.com/di-wu/parser/ast"
	"strings"
)

// Func indicates the function’s signature (argument and results types, annotations),
// and values of this type are references to functions with that signature.
type Func struct {
	ArgTypes, ResTypes Tuple
	Annotation         *FuncAnnotation
}

func (f Func) String() string {
	s := fmt.Sprintf("%s -> %s", f.ArgTypes.String(), f.ResTypes.String())
	if f.Annotation != nil {
		s += fmt.Sprintf(" %s", *f.Annotation)
	}
	return s
}

func convertFunc(n *ast.Node) Func {
	var f Func
	for i, n := range n.Children() {
		switch i {
		case 0:
			f.ArgTypes = convertTuple(n)
		case 1:
			f.ResTypes = convertTuple(n)
		case 2:
			ann := FuncAnnotation(n.Value)
			f.Annotation = &ann
		default:
			panic(n)
		}
	}
	return f
}

// FuncAnnotation represents a function annotation.
type FuncAnnotation string

const (
	// AnnOneWay indicates that this function returns no response, intended for
	// fire-and-forget scenarios.
	AnnOneWay FuncAnnotation = "oneway"
	// AnnQuery indicates that the referenced function is a query method, meaning
	// it does not alter the state of its canister, and that it can be invoked
	// using the cheaper “query call” mechanism.
	AnnQuery FuncAnnotation = "query"
)

// Tuple represents one or more arguments.
type Tuple []Argument

func (t Tuple) String() string {
	if len(t) == 1 {
		s := t[0].String()
		if strings.Contains(s, " ") {
			return "(" + s + ")"
		}
		return s
	}
	s := "("
	for i, a := range t {
		s += a.String()
		if i != len(t)-1 {
			s += ", "
		}
	}
	return s + ")"
}

func convertTuple(n *ast.Node) Tuple {
	var tuple Tuple
	for _, n := range n.Children() {
		tuple = append(tuple, convertArgument(n))
	}
	return tuple
}

type Argument struct {
	Name *string
	Data Data
}

func (a Argument) String() string {
	var s string
	if a.Name != nil {
		s += fmt.Sprintf("%s : ", *a.Name)
	}
	return s + a.Data.String()
}

func convertArgument(n *ast.Node) Argument {
	data := convertData(n.LastChild)
	if len(n.Children()) == 1 {
		return Argument{
			Data: data,
		}
	}
	return Argument{
		Name: &n.FirstChild.Value,
		Data: data,
	}
}
