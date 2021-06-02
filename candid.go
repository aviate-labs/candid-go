//go:generate go get github.com/pegn/pegn-go
//go:generate go run internal/gen.go
//go:generate go mod tidy

package candid

import (
	"fmt"
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

func (p Program) String() string {
	var s []string
	for _, d := range p.definitions {
		s = append(s, d.String())
	}
	for _, a := range p.actors {
		s = append(s, a.String())
	}
	return strings.Join(s, ";\n") + ";"
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
	fmt.Stringer
}

func (t Type) def()   {}
func (i Import) def() {}

type Type struct {
	Id   string
	Data Data
}

func (t Type) String() string {
	return fmt.Sprintf("type %s = %s", t.Id, t.Data.String())
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

func (i Import) String() string {
	return fmt.Sprintf("import %q", i.Text)
}

type Actor struct {
	Id *string

	Methods  []Method
	MethodId *string
}

func (a Actor) String() string {
	s := "service "
	if id := a.Id; id != nil {
		s += fmt.Sprintf("%s ", *id)
	}
	s += ": "
	if id := a.MethodId; id != nil {
		return s + *id
	}
	s += "{\n"
	for _, m := range a.Methods {
		s += fmt.Sprintf("  %s;\n", m.String())
	}
	return s + "}"
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
					f := convertFunc(n)
					actor.Methods = append(
						actor.Methods,
						Method{
							Name: name,
							Func: &f,
						},
					)
				case grammar.IdT, grammar.TextT:
					id := n.Value
					actor.Methods = append(
						actor.Methods,
						Method{
							Name: name,
							Id:   &id,
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

type Method struct {
	Name string

	Func *Func
	Id   *string
}

func (m Method) String() string {
	s := fmt.Sprintf("%s : ", m.Name)
	if id := m.Id; id != nil {
		return s + *id
	}
	return s + m.Func.String()
}

type Func struct {
	From, To   Tuple
	Annotation *FuncAnnotation
}

func (f Func) String() string {
	s := fmt.Sprintf("%s -> %s", f.From.String(), f.To.String())
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
			f.From = convertTuple(n)
		case 1:
			f.To = convertTuple(n)
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

type Field struct {
	Nat  *big.Int
	Name *string

	Data     *Data
	NatData  *big.Int
	NameData *string
}

func (f Field) String() string {
	var s string
	if n := f.Nat; n != nil {
		s += fmt.Sprintf("%s : ", n.String())
	} else if f.Name != nil {
		s += fmt.Sprintf("%s : ", *f.Name)
	}
	if f.Data != nil {
		d := *f.Data
		s += d.String()
	} else if n := f.NatData; n != nil {
		s += n.String()
	} else {
		s += *f.NameData
	}
	return s
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
	case grammar.NatT:
		field.NatData = convertNat(n)
	case grammar.IdT:
		field.NameData = &n.Value
	default:
		data := convertData(n)
		field.Data = &data
	}
	return field
}

type Data interface {
	data()
	fmt.Stringer
}

func (p Primitive) data() {}
func (i DataId) data()    {}
func (b Blob) data()      {}
func (o Optional) data()  {}
func (v Vector) data()    {}
func (r Record) data()    {}
func (v Variant) data()   {}

func convertData(n *ast.Node) Data {
	switch n.Type {
	case grammar.BlobT:
		return Blob{}
	case grammar.OptT:
		return Optional{
			Data: convertData(n.FirstChild),
		}
	case grammar.VecT:
		return Vector{
			Data: convertData(n.FirstChild),
		}
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

func (b Blob) String() string {
	return "blob"
}

type Optional struct {
	Data Data
}

func (o Optional) String() string {
	return fmt.Sprintf("opt %s", o.Data.String())
}

type Vector struct {
	Data Data
}

func (v Vector) String() string {
	return fmt.Sprintf("vec %s", v.Data.String())
}

type Record []Field

func (r Record) String() string {
	s := "record {\n"
	for _, f := range r {
		s += fmt.Sprintf("  %s;\n", f.String())
	}
	return s + "}"
}

type Variant []Field

func (v Variant) String() string {
	s := "variant {\n"
	for _, f := range v {
		s += fmt.Sprintf("  %s;\n", f.String())
	}
	return s + "}"
}

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

func (p Primitive) String() string {
	return string(p)
}

type DataId string

func (i DataId) String() string {
	return string(i)
}

type Principal struct{}

func (p Principal) String() string {
	return "principal"
}

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
