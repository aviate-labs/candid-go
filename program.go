package candid

import (
	spec "github.com/allusion-be/candid-go/internal/grammar"
	"github.com/di-wu/parser/ast"
	"strings"
)

// Program represents the program defined withing a .did file.
type Program struct {
	Definitions []Definition
	Services    []Service
}

func (p Program) String() string {
	var s []string
	for _, d := range p.Definitions {
		s = append(s, d.String())
	}
	for _, a := range p.Services {
		s = append(s, a.String())
	}
	return strings.Join(s, ";\n")
}

func convertProgram(n *ast.Node) Program {
	var program Program
	for _, n := range n.Children() {
		switch n.Type {
		case spec.TypeT:
			program.Definitions = append(
				program.Definitions,
				convertType(n),
			)
		case spec.ImportT:
			program.Definitions = append(
				program.Definitions,
				Import{
					Text: "",
				},
			)
		case spec.ActorT:
			program.Services = append(
				program.Services,
				convertService(n),
			)
		default:
			panic(n)
		}
	}
	return program
}
