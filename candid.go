package candid

import (
	"github.com/aviate-labs/candid-go/internal/candid"
	"github.com/di-wu/parser"
	"github.com/di-wu/parser/ast"
)

// ParseDID parses the given raw .did files and returns the Program that is defined in it.
func ParseDID(raw []byte) (Description, error) {
	p, err := ast.New(raw)
	if err != nil {
		return Description{}, err
	}
	n, err := candid.Prog(p)
	if err != nil {
		return Description{}, err
	}
	if _, err := p.Expect(parser.EOD); err != nil {
		return Description{}, err
	}
	return convertDescription(n), nil
}
