package candid

import (
	"fmt"
	"github.com/di-wu/parser/ast"
)

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
