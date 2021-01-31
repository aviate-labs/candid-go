// Do not edit. This file is auto-generated.
// Grammar: CANDID (v0.1.0) github.com/di-wu/candid-go/spec
package candid

import (
	"github.com/di-wu/parser"
	"github.com/di-wu/parser/ast"
	"github.com/di-wu/parser/op"
)

func Prog(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: ProgT,
			Value: op.And{
				op.MinZero(
					op.And{
						Def,
						';',
					},
				),
				op.MinZero(
					op.And{
						Actor,
						';',
					},
				),
			},
		},
	)
}

func Def(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: DefT,
			Value: op.Or{
				op.And{
					"type",
					Sp,
					Id,
					Sp,
					'=',
					Sp,
					DataType,
				},
				op.And{
					"import",
					Sp,
					Text,
				},
			},
		},
	)
}

func Actor(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: ActorT,
			Value: op.And{
				"service",
				Sp,
				op.Optional(
					op.And{
						Id,
						Sp,
					},
				),
				':',
				Sp,
				op.Optional(
					op.And{
						TupType,
						Sp,
						"->",
						Ws,
					},
				),
				op.Or{
					ActorType,
					Id,
				},
			},
		},
	)
}

func ActorType(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: ActorTypeT,
			Value: op.And{
				'{',
				Ws,
				MethType,
				op.MinZero(
					op.And{
						';',
						Ws,
						MethType,
					},
				),
				op.Optional(
					';',
				),
				Ws,
				'}',
			},
		},
	)
}

func MethType(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: MethTypeT,
			Value: op.And{
				Name,
				Sp,
				':',
				Sp,
				op.Or{
					FuncType,
					Id,
				},
			},
		},
	)
}

func FuncType(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: FuncTypeT,
			Value: op.And{
				TupType,
				op.Optional(
					op.And{
						Sp,
						"->",
						Ws,
						TupType,
						op.Optional(
							op.And{
								Sp,
								FuncAnn,
							},
						),
					},
				),
			},
		},
	)
}

func FuncAnn(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: FuncAnnT,
			Value: op.Or{
				"oneway",
				"query",
			},
		},
	)
}

func TupType(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: TupTypeT,
			Value: op.Or{
				op.And{
					'(',
					op.Optional(
						op.And{
							ArgType,
							op.MinZero(
								op.And{
									',',
									Sp,
									ArgType,
								},
							),
						},
					),
					')',
				},
				ArgType,
			},
		},
	)
}

func ArgType(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: ArgTypeT,
			Value: op.And{
				op.Optional(
					op.And{
						Name,
						Sp,
						':',
						Sp,
					},
				),
				DataType,
			},
		},
	)
}

func FieldType(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: FieldTypeT,
			Value: op.Or{
				op.And{
					op.Optional(
						op.And{
							op.Or{
								Nat,
								Name,
							},
							Sp,
							':',
							Sp,
						},
					),
					DataType,
				},
				Nat,
				Name,
			},
		},
	)
}

func DataType(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: DataTypeT,
			Value: op.Or{
				ConsType,
				RefType,
				PrimType,
				Id,
			},
		},
	)
}

func PrimType(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: PrimTypeT,
			Value: op.Or{
				NumType,
				"bool",
				"text",
				"null",
				"reserved",
				"empty",
			},
		},
	)
}

func NumType(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.Or{
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
	)
}

func ConsType(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.Or{
			Blob,
			Opt,
			Vec,
			Record,
			Variant,
		},
	)
}

func Blob(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:  BlobT,
			Value: "blob",
		},
	)
}

func Opt(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: OptT,
			Value: op.And{
				"opt",
				Sp,
				DataType,
			},
		},
	)
}

func Vec(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: VecT,
			Value: op.And{
				"vec",
				Sp,
				DataType,
			},
		},
	)
}

func Record(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: RecordT,
			Value: op.And{
				"record",
				Sp,
				'{',
				op.Optional(
					Fields,
				),
				Ws,
				'}',
			},
		},
	)
}

func Variant(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: VariantT,
			Value: op.And{
				"variant",
				Sp,
				'{',
				op.Optional(
					Fields,
				),
				Ws,
				'}',
			},
		},
	)
}

func Fields(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.And{
			Ws,
			FieldType,
			op.MinZero(
				op.And{
					';',
					Ws,
					FieldType,
				},
			),
			op.Optional(
				';',
			),
		},
	)
}

func RefType(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: RefTypeT,
			Value: op.Or{
				op.And{
					"func",
					Sp,
					FuncType,
				},
				op.And{
					"service",
					Sp,
					ActorType,
				},
				"principal",
			},
		},
	)
}

func Name(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.Or{
			Id,
			Text,
		},
	)
}

func Id(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: IdT,
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

func Text(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: TextT,
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

func Num(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.And{
			Digit,
			op.MinZero(
				op.And{
					'_',
					Digit,
				},
			),
		},
	)
}

func Nat(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: NatT,
			Value: op.Or{
				Num,
				op.And{
					"0x",
					HexNum,
				},
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
					'_',
					Hex,
				},
			),
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

func Ws(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.MinZero(
			op.Or{
				Sp,
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

func Sp(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.MinOne(
			' ',
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
	TODO = '\u0000' // TODO: remove this.

	// CANDID (github.com/di-wu/candid-go/spec)
	ESC = 0x005C // \

)

// Node Types
const (
	Unknown = iota

	// CANDID (github.com/di-wu/candid-go/spec)
	ProgT      // 001
	DefT       // 002
	ActorT     // 003
	ActorTypeT // 004
	MethTypeT  // 005
	FuncTypeT  // 006
	FuncAnnT   // 007
	TupTypeT   // 008
	ArgTypeT   // 009
	FieldTypeT // 010
	DataTypeT  // 011
	PrimTypeT  // 012
	BlobT      // 013
	OptT       // 014
	VecT       // 015
	RecordT    // 016
	VariantT   // 017
	RefTypeT   // 018
	IdT        // 019
	TextT      // 020
	NatT       // 021
)

var NodeTypes = []string{
	"UNKNOWN",

	// CANDID (github.com/di-wu/candid-go/spec)
	"Prog",
	"Def",
	"Actor",
	"ActorType",
	"MethType",
	"FuncType",
	"FuncAnn",
	"TupType",
	"ArgType",
	"FieldType",
	"DataType",
	"PrimType",
	"Blob",
	"Opt",
	"Vec",
	"Record",
	"Variant",
	"RefType",
	"Id",
	"Text",
	"Nat",
}
