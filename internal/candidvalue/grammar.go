// Do not edit. This file is auto-generated.
// Grammar: CANDID (v0.1.0) github.com/di-wu/candid-go/internal/candid/candidvalue

package candidvalue

import (
	"github.com/di-wu/parser"
	"github.com/di-wu/parser/ast"
	"github.com/di-wu/parser/op"
)

func Values(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        ValuesT,
			TypeStrings: NodeTypes,
			Value: op.Or{
				op.And{
					'(',
					Sp,
					op.Optional(
						op.And{
							Value,
							op.MinZero(
								op.And{
									Sp,
									',',
									Sp,
									Value,
								},
							),
						},
					),
					Sp,
					')',
				},
				Value,
			},
		},
	)
}

func Value(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		Num,
	)
}

func Num(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        NumT,
			TypeStrings: NodeTypes,
			Value: op.And{
				NumValue,
				op.Optional(
					op.And{
						Sp,
						':',
						Sp,
						NumType,
					},
				),
			},
		},
	)
}

func NumValue(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        NumValueT,
			TypeStrings: NodeTypes,
			Value: op.And{
				op.Optional(
					'-',
				),
				Digit,
				op.MinZero(
					op.And{
						op.Optional(
							'_',
						),
						Digit,
					},
				),
				op.Optional(
					op.And{
						'.',
						Digit,
						op.MinZero(
							op.And{
								op.Optional(
									'_',
								),
								Digit,
							},
						),
					},
				),
			},
		},
	)
}

func NumType(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        NumTypeT,
			TypeStrings: NodeTypes,
			Value: op.Or{
				"nat8",
				"nat16",
				"nat32",
				"nat64",
				"nat",
				"int8",
				"int16",
				"int32",
				"int64",
				"int",
				"float32",
				"float64",
			},
		},
	)
}

func Sp(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.MinZero(
			' ',
		),
	)
}

func Digit(p *parser.Parser) (*parser.Cursor, bool) {
	return p.Check(parser.CheckRuneRange('0', '9'))
}

// Node Types
const (
	Unknown = iota

	// CANDID (github.com/di-wu/candid-go/internal/candid/candidvalue)

	ValuesT   // 001
	NumT      // 002
	NumValueT // 003
	NumTypeT  // 004
)

var NodeTypes = []string{
	"UNKNOWN",

	// CANDID (github.com/di-wu/candid-go/internal/candid/candidvalue)

	"Values",
	"Num",
	"NumValue",
	"NumType",
}
