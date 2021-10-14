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
		op.Or{
			OptValue,
			Num,
			Bool,
			Null,
			Text,
			Record,
			Variant,
		},
	)
}

func OptValue(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.Or{
			op.And{
				"opt",
				Sp,
				Num,
			},
			Bool,
			Null,
			Text,
			Record,
			Variant,
		},
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

func Bool(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.And{
			BoolValue,
			op.Optional(
				op.And{
					Sp,
					':',
					Sp,
					"bool",
				},
			),
		},
	)
}

func BoolValue(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        BoolValueT,
			TypeStrings: NodeTypes,
			Value: op.Or{
	"true",
	"false",
},
		},
	)
}

func Null(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        NullT,
			TypeStrings: NodeTypes,
			Value:       "null",
		},
	)
}

func Text(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        TextT,
			TypeStrings: NodeTypes,
			Value: op.And{
				TextValue,
				op.Optional(
					op.And{
						Sp,
						':',
						Sp,
						"text",
					},
				),
			},
		},
	)
}

func TextValue(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        TextValueT,
			TypeStrings: NodeTypes,
			Value: op.And{
				'"',
				op.MinZero(
					Char,
				),
				'"',
			},
		},
	)
}

func Char(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.Or{
			Utf,
			op.And{
				ESC,
				op.Repeat(2,
					Hex,
				),
			},
			op.And{
				ESC,
				Escape,
			},
			op.And{
				"\\u{",
				HexNum,
				'}',
			},
		},
	)
}

func Utf(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.Or{
			Ascii,
			UtfEnc,
		},
	)
}

func UtfEnc(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.Or{
			op.And{
				parser.CheckRuneRange(0x00C2, 0x00DF),
				Utfcont,
			},
			op.And{
				0x00E0,
				parser.CheckRuneRange(0x00A0, 0x00BF),
				Utfcont,
			},
			op.And{
				0x00ED,
				parser.CheckRuneRange(0x0080, 0x009F),
				Utfcont,
			},
			op.And{
				parser.CheckRuneRange(0x00E1, 0x00EC),
				op.Repeat(2,
					Utfcont,
				),
			},
			op.And{
				parser.CheckRuneRange(0x00EE, 0x00EF),
				op.Repeat(2,
					Utfcont,
				),
			},
			op.And{
				0x00F0,
				parser.CheckRuneRange(0x0090, 0x00BF),
				op.Repeat(2,
					Utfcont,
				),
			},
			op.And{
				0x00F4,
				parser.CheckRuneRange(0x0080, 0x008F),
				op.Repeat(2,
					Utfcont,
				),
			},
			op.And{
				parser.CheckRuneRange(0x00F1, 0x00F3),
				op.Repeat(3,
					Utfcont,
				),
			},
		},
	)
}

func HexNum(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.And{
			Hex,
			op.MinZero(
				op.And{
					op.Optional(
						'_',
					),
					Hex,
				},
			),
		},
	)
}

func Record(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        RecordT,
			TypeStrings: NodeTypes,
			Value: op.And{
				"record",
				Sp,
				'{',
				Ws,
				op.Optional(
					RecordFields,
				),
				Ws,
				'}',
			},
		},
	)
}

func RecordFields(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.And{
			RecordField,
			Sp,
			op.MinZero(
				op.And{
					';',
					Ws,
					RecordField,
					Sp,
				},
			),
			op.Optional(
				';',
			),
		},
	)
}

func RecordField(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        RecordFieldT,
			TypeStrings: NodeTypes,
			Value: op.And{
				Id,
				Sp,
				'=',
				Sp,
				Value,
			},
		},
	)
}

func Variant(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        VariantT,
			TypeStrings: NodeTypes,
			Value: op.And{
				"variant",
				Sp,
				'{',
				Ws,
				VariantField,
				Ws,
				'}',
			},
		},
	)
}

func VariantField(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.And{
			Id,
			op.Optional(
				op.And{
					Sp,
					'=',
					Sp,
					Value,
				},
			),
			op.Optional(
				';',
			),
		},
	)
}

func Id(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        IdT,
			TypeStrings: NodeTypes,
			Value: op.And{
				op.Or{
					Letter,
					'_',
				},
				op.MinZero(
					op.Or{
						Letter,
						Digit,
						'_',
					},
				),
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

func Ws(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.MinZero(
			op.Or{
				' ',
				0x0009,
				0x000A,
				0x000D,
				op.And{
					0x000D,
					0x000A,
				},
			},
		),
	)
}

func Utfcont(p *parser.Parser) (*parser.Cursor, bool) {
	return p.Check(parser.CheckRuneRange(0x0080, 0x00BF))
}

func Ascii(p *parser.Parser) (*parser.Cursor, bool) {
	return p.Check(op.Or{
		parser.CheckRuneRange(0x0020, 0x0021),
		parser.CheckRuneRange(0x0023, 0x005B),
		parser.CheckRuneRange(0x005D, 0x007E),
	})
}

func Escape(p *parser.Parser) (*parser.Cursor, bool) {
	return p.Check(op.Or{
		'n',
		'r',
		't',
		ESC,
		0x0022,
		0x0027,
	})
}

func Letter(p *parser.Parser) (*parser.Cursor, bool) {
	return p.Check(op.Or{
		parser.CheckRuneRange('A', 'Z'),
		parser.CheckRuneRange('a', 'z'),
	})
}

func Digit(p *parser.Parser) (*parser.Cursor, bool) {
	return p.Check(parser.CheckRuneRange('0', '9'))
}

func Hex(p *parser.Parser) (*parser.Cursor, bool) {
	return p.Check(op.Or{
		Digit,
		parser.CheckRuneRange('A', 'F'),
		parser.CheckRuneRange('a', 'f'),
	})
}

// Token Definitions
const (
	// CANDID (github.com/di-wu/candid-go/internal/candid/candidvalue)

	ESC = 0x005C // \
)

// Node Types
const (
	Unknown = iota

	// CANDID (github.com/di-wu/candid-go/internal/candid/candidvalue)

	ValuesT      // 001
	NumT         // 002
	NumValueT    // 003
	NumTypeT     // 004
	BoolValueT   // 005
	NullT        // 006
	TextT        // 007
	TextValueT   // 008
	RecordT      // 009
	RecordFieldT // 010
	VariantT     // 011
	IdT          // 012
)

var NodeTypes = []string{
	"UNKNOWN",

	// CANDID (github.com/di-wu/candid-go/internal/candid/candidvalue)

	"Values",
	"Num",
	"NumValue",
	"NumType",
	"BoolValue",
	"Null",
	"Text",
	"TextValue",
	"Record",
	"RecordField",
	"Variant",
	"Id",
}
