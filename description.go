package candid

import (
	spec "github.com/allusion-be/candid-go/internal/grammar"
	"github.com/di-wu/parser/ast"
	"strings"
)

// Description represents the interface description of a program. An interface description consists of a sequence of
// imports and type definitions, possibly followed by a service declaration.
type Description struct {
	// Definitions is the sequence of import and type definitions.
	Definitions []Definition
	// Services is a list of service declarations.
	Services []Service
}

func (p Description) String() string {
	var s []string
	for _, d := range p.Definitions {
		s = append(s, d.String())
	}
	for _, a := range p.Services {
		s = append(s, a.String())
	}
	return strings.Join(s, ";\n")
}

func convertDescription(n *ast.Node) Description {
	var desc Description
	for _, n := range n.Children() {
		switch n.Type {
		case spec.TypeT:
			desc.Definitions = append(
				desc.Definitions,
				convertType(n),
			)
		case spec.ImportT:
			desc.Definitions = append(
				desc.Definitions,
				Import{
					Text: "",
				},
			)
		case spec.ActorT:
			desc.Services = append(
				desc.Services,
				convertService(n),
			)
		default:
			panic(n)
		}
	}
	return desc
}