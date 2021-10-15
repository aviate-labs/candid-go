package candid

import (
	"github.com/aviate-labs/candid-go/did"
	"github.com/aviate-labs/candid-go/idl"
	"github.com/aviate-labs/candid-go/internal/candid"
	"github.com/aviate-labs/candid-go/internal/candidvalue"
	"github.com/di-wu/parser"
	"github.com/di-wu/parser/ast"
)

// ParseDID parses the given raw .did files and returns the Program that is defined in it.
func ParseDID(raw []byte) (did.Description, error) {
	p, err := ast.New(raw)
	if err != nil {
		return did.Description{}, err
	}
	n, err := candid.Prog(p)
	if err != nil {
		return did.Description{}, err
	}
	if _, err := p.Expect(parser.EOD); err != nil {
		return did.Description{}, err
	}
	return did.ConvertDescription(n), nil
}

func EncodeValue(value string) ([]byte, error) {
	p, err := ast.New([]byte(value))
	if err != nil {
		return nil, err
	}
	n, err := candidvalue.Values(p)
	if err != nil {
		return nil, err
	}
	if _, err := p.Expect(parser.EOD); err != nil {
		return nil, err
	}
	types, args, err := did.ConvertValues(n)
	if err != nil {
		return nil, err
	}
	return idl.Encode(types, args)
}
