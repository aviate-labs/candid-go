//go:generate go get github.com/pegn/pegn-go
//go:generate go run internal/gen.go
//go:generate go mod tidy

package candid

import (
	"github.com/di-wu/parser"
	"github.com/di-wu/parser/ast"
	grammar "github.com/internet-computer/candid-go/internal/grammar"
	"math/big"
	"strconv"
	"strings"
)

func ParseDID(raw []byte) (Program, error) {
	p, err := ast.New(raw)
	if err != nil {
		return Program{}, err
	}
	n, err := grammar.Prog(p)
	if err != nil {
		return Program{}, err
	}
	if _, err := p.Expect(parser.EOD); err != nil {
		return Program{}, err
	}
	return convertProgram(n), nil
}

type Program struct {
	definitions []Definition
	actors      []Actor
}

func convertProgram(n *ast.Node) Program {
	var program Program
	for _, n := range n.Children() {
		switch n.Type {
		case grammar.TypeT:
			program.definitions = append(
				program.definitions,
				convertType(n),
			)
		case grammar.ImportT:
			program.definitions = append(
				program.definitions,
				Import{
					Text: "",
				},
			)
		case grammar.ActorT:
			program.actors = append(
				program.actors,
				convertActor(n),
			)
		default:
			panic(n)
		}
	}
	return program
}

type Definition interface {
	def()
}

func (t Type) def()   {}
func (i Import) def() {}

type Type struct {
	Id   string
	Data Data
}

func convertType(n *ast.Node) Type {
	var (
		id   = n.FirstChild
		data = n.LastChild
	)
	return Type{
		Id:   id.Value,
		Data: convertData(data),
	}
}

type Import struct {
	Text string
}

type Actor struct {
	Id *string

	Methods  []Method
	MethodId *string
}

func convertActor(n *ast.Node) Actor {
	var actor Actor
	for _, n := range n.Children() {
		switch n.Type {
		case grammar.IdT:
			id := n.Value
			if actor.Id == nil {
				actor.Id = &id
				continue
			}
			actor.MethodId = &id
			break
		case grammar.TupTypeT:
			// TODO
		case grammar.ActorTypeT:
			for _, n := range n.Children() {
				name := n.FirstChild.Value
				switch n := n.LastChild; n.Type {
				case grammar.FuncTypeT:
					actor.Methods = append(
						actor.Methods,
						MethodFunc{
							Name: name,
							Func: convertFunc(n),
						},
					)
				case grammar.IdT, grammar.TextT:
					actor.Methods = append(
						actor.Methods,
						MethodId{
							Name: name,
							Id:   n.Value,
						},
					)
				default:
					panic(n)
				}
			}
		default:
			panic(n)
		}
	}
	return actor
}

type Method interface {
	method()
}

func (m MethodId) method()   {}
func (m MethodFunc) method() {}

type MethodId struct {
	Name string
	Id   string
}

type MethodFunc struct {
	Name string
	Func Func
}

type Func struct {
	From, To   Tuple
	Annotation *FuncAnnotation
}

func convertFunc(n *ast.Node) Func {
	var f Func
	for i, n := range n.Children() {
		switch i {
		case 0:
			f.To = convertTuple(n)
		case 1:
			f.From = convertTuple(n)
		case 2:
			ann := FuncAnnotation(n.Value)
			f.Annotation = &ann
		default:
			panic(n)
		}
	}
	return f
}

type FuncAnnotation string

const (
	AnnOneWay FuncAnnotation = "oneway"
	AnnQuery  FuncAnnotation = "query"
)

type Tuple []Argument

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

type Field struct {
	Nat  *big.Int
	Name *string

	Data     *Data
	NatData  *big.Int
	NameData *string
}

func convertField(n *ast.Node) Field {
	var field Field
	if len(n.Children()) != 1 {
		switch n := n.FirstChild; n.Type {
		case grammar.NatT:
			field.Nat = convertNat(n)
		case grammar.TextT, grammar.IdT:
			field.Name = &n.Value
		default:
			panic(n)
		}
	}
	switch n := n.LastChild; n.Type {
	case grammar.DataTypeT:
		data := convertData(n)
		field.Data = &data
	case grammar.NatT:
		field.NatData = convertNat(n)
	case grammar.IdT:
		field.NameData = &n.Value
	}
	return field
}

type Data interface {
	data()
}

func (p Primitive) data() {}
func (i DataId) data()    {}
func (b Blob) data()      {}
func (r Record) data()    {}
func (v Variant) data()   {}

func convertData(n *ast.Node) Data {
	switch n.Type {
	case grammar.BlobT:
		return Blob{}
	case grammar.OptT:
		return Optional(convertData(n.FirstChild))
	case grammar.VecT:
		return Vector(convertData(n.FirstChild))
	case grammar.RecordT:
		var record Record
		for _, n := range n.Children() {
			record = append(
				record,
				convertField(n),
			)
		}
		return record
	case grammar.VariantT:
		var variant Variant
		for _, n := range n.Children() {
			variant = append(
				variant,
				convertField(n),
			)
		}
		return variant
	case grammar.FuncT:
		return convertFunc(n.FirstChild)
	case grammar.ServiceT:
		return convertActor(n.FirstChild)
	case grammar.PrincipalT:
		return Principal{}
	case grammar.PrimTypeT:
		return Primitive(n.Value)
	case grammar.IdT:
		return DataId(n.Value)
	default:
		panic(n)
	}
}

type Blob struct{}

type Optional Data

type Vector Data

type Record []Field

type Variant []Field

type Reference interface {
	data()
	ref()
}

func (f Func) data()      {}
func (f Func) ref()       {}
func (a Actor) data()     {}
func (a Actor) ref()      {}
func (p Principal) data() {}
func (p Principal) ref()  {}

type Primitive string

type DataId string

type Principal struct{}

func convertNat(n *ast.Node) *big.Int {
	switch n := strings.ReplaceAll(n.Value, "_", ""); {
	case strings.HasPrefix(n, "0x"):
		n = strings.TrimPrefix(n, "0x")
		i, _ := strconv.ParseInt(n, 16, 64)
		return big.NewInt(i)
	default:
		i, _ := strconv.ParseInt(n, 10, 64)
		return big.NewInt(i)
	}
}
