package idl2

type PrimType uint8

const (
	Null      PrimType = 0x7f // sleb128(-1)
	Bool      PrimType = 0x7e // sleb128(-2)
	Nat       PrimType = 0x7d // sleb128(-3)
	Int       PrimType = 0x7c // sleb128(-4)
	Nat8      PrimType = 0x7b // sleb128(-5)
	Nat16     PrimType = 0x7a // sleb128(-6)
	Nat32     PrimType = 0x79 // sleb128(-7)
	Nat64     PrimType = 0x78 // sleb128(-8)
	Int8      PrimType = 0x77 // sleb128(-9)
	Int16     PrimType = 0x76 // sleb128(-10)
	Int32     PrimType = 0x75 // sleb128(-11)
	Int64     PrimType = 0x74 // sleb128(-12)
	Float32   PrimType = 0x73 // sleb128(-13)
	Float64   PrimType = 0x72 // sleb128(-14)
	Text      PrimType = 0x71 // sleb128(-15)
	Reserved  PrimType = 0x70 // sleb128(-16)
	Empty     PrimType = 0x6f // sleb128(-17)
	Principal PrimType = 0x68 // sleb128(-24)
)

func (c PrimType) bytes() []byte {
	return []byte{uint8(c)}
}
