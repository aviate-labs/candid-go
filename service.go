package candid

import (
	"fmt"
	"github.com/di-wu/parser/ast"
	spec "github.com/internet-computer/candid-go/internal/grammar"
)

// Service can be used to declare the complete interface of a service.
type Service struct {
	Id *string

	Methods  []Method
	MethodId *string
}

func (a Service) String() string {
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

func convertService(n *ast.Node) Service {
	var actor Service
	for _, n := range n.Children() {
		switch n.Type {
		case spec.IdT:
			id := n.Value
			if actor.Id == nil {
				actor.Id = &id
				continue
			}
			actor.MethodId = &id
			break
		case spec.TupTypeT:
			// TODO: what does this even do?
		case spec.ActorTypeT:
			for _, n := range n.Children() {
				name := n.FirstChild.Value
				switch n := n.LastChild; n.Type {
				case spec.FuncTypeT:
					f := convertFunc(n)
					actor.Methods = append(
						actor.Methods,
						Method{
							Name: name,
							Func: &f,
						},
					)
				case spec.IdT, spec.TextT:
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
