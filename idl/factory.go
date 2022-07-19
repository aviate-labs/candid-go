package idl

type IDL struct {
	Null      *Null
	Bool      *Bool
	Nat       *Nat
	Int       *Int
	Nat8      *Nat
	Nat16     *Nat
	Nat32     *Nat
	Nat64     *Nat
	Int8      *Int
	Int16     *Int
	Int32     *Int
	Int64     *Int
	Float32   *Float
	Float64   *Float
	Text      *Text
	Reserved  *Reserved
	Empty     *Empty
	Opt       func(typ Type) *Opt[Type]
	Tuple     func(ts ...Type) *Tuple
	Vec       func(t Type) *Vec
	Record    func(fields map[string]Type) *Rec
	Variant   func(fields map[string]Type) *Variant
	Func      func(args []Type, ret []Type, annotations []string) *Func
	Service   func(functions map[string]*Func) *Service
	Principal *Principal
}

type IDLFactory = func(types IDL) *Service

func NewInterface(factory IDLFactory) *Service {
	return factory(IDL{
		Bool:     new(Bool),
		Null:     new(Null),
		Nat:      new(Nat),
		Int:      new(Int),
		Nat8:     Nat8(),
		Nat16:    Nat16(),
		Nat32:    Nat32(),
		Nat64:    Nat64(),
		Int8:     Int8(),
		Int16:    Int16(),
		Int32:    Int32(),
		Int64:    Int64(),
		Text:     new(Text),
		Reserved: new(Reserved),
		Empty:    new(Empty),
		Opt: func(typ Type) *Opt[Type] {
			return &Opt[Type]{Type: typ}
		},
		Tuple: func(ts ...Type) *Tuple {
			tuple := Tuple(ts)
			return &tuple
		},
		Vec: func(t Type) *Vec {
			return NewVec(t)
		},
		Record: func(fields map[string]Type) *Rec {
			return NewRec(fields)
		},
		Variant: func(fields map[string]Type) *Variant {
			return NewVariant(fields)
		},
		Func: func(argumentTypes, returnTypes []Type, annotations []string) *Func {
			return NewFunc(argumentTypes, returnTypes, annotations)
		},
		Service: func(methods map[string]*Func) *Service {
			return NewService(methods)
		},
		Principal: new(Principal),
	})
}
