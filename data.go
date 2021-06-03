package candid

import (
	"fmt"
	spec "github.com/allusion-be/candid-go/internal/grammar"
	"github.com/di-wu/parser/ast"
	"math/big"
	"strconv"
	"strings"
)

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
func (f Func) data()      {}
func (a Service) data()   {}
func (p Principal) data() {}

func convertData(n *ast.Node) Data {
	switch n.Type {
	case spec.BlobT:
		return Blob{}
	case spec.OptT:
		return Optional{
			Data: convertData(n.FirstChild),
		}
	case spec.VecT:
		return Vector{
			Data: convertData(n.FirstChild),
		}
	case spec.RecordT:
		var record Record
		for _, n := range n.Children() {
			record = append(
				record,
				convertField(n),
			)
		}
		return record
	case spec.VariantT:
		var variant Variant
		for _, n := range n.Children() {
			variant = append(
				variant,
				convertField(n),
			)
		}
		return variant
	case spec.FuncT:
		return convertFunc(n.FirstChild)
	case spec.ServiceT:
		return convertService(n.FirstChild)
	case spec.PrincipalT:
		return Principal{}
	case spec.PrimTypeT:
		return Primitive(n.Value)
	case spec.IdT:
		return DataId(n.Value)
	default:
		panic(n)
	}
}

// Blob can be used for binary data, that is, sequences of bytes.
type Blob struct{}

func (b Blob) String() string {
	return "blob"
}

// Optional is used to express that some value is optional, meaning that data might
// be present as some value of type t, or might be absent as the value null.
type Optional struct {
	Data Data
}

func (o Optional) String() string {
	return fmt.Sprintf("opt %s", o.Data.String())
}

// Vector represents vectors (sequences, lists, arrays).
// e.g. 'vec bool', 'vec nat8', 'vec vec text', etc
type Vector struct {
	Data Data
}

func (v Vector) String() string {
	return fmt.Sprintf("vec %s", v.Data.String())
}

// Record a collection of labeled values.
type Record []Field

func (r Record) String() string {
	s := "record {\n"
	for _, f := range r {
		s += fmt.Sprintf("  %s;\n", f.String())
	}
	return s + "}"
}

// Variant represents a value that is from exactly one of the given cases, or tags.
type Variant []Field

func (v Variant) String() string {
	s := "variant {\n"
	for _, f := range v {
		s += fmt.Sprintf("  %s;\n", f.String())
	}
	return s + "}"
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
		case spec.NatT:
			field.Nat = convertNat(n)
		case spec.TextT, spec.IdT:
			field.Name = &n.Value
		default:
			panic(n)
		}
	}
	switch n := n.LastChild; n.Type {
	case spec.NatT:
		field.NatData = convertNat(n)
	case spec.IdT:
		field.NameData = &n.Value
	default:
		data := convertData(n)
		field.Data = &data
	}
	return field
}

type Primitive string

func (p Primitive) String() string {
	return string(p)
}

type DataId string

func (i DataId) String() string {
	return string(i)
}

// Principal is the common scheme to identify canisters, users, and other entities.
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
