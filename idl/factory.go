package idl

import "sort"

type IDLFactory = func(idl IDL) Service

type IDL struct {
	Null     Null
	Bool     Bool
	Nat      Nat
	Int      Int
	Nat8     Nat
	Nat16    Nat
	Nat32    Nat
	Nat64    Nat
	Int8     Int
	Int16    Int
	Int32    Int
	Int64    Int
	Float32  Float
	Float64  Float
	Text     Text
	Reserved Reserved
	Empty    Empty
	Opt      Opt
	Tuple    func(ts ...Type) Tuple
	Vec      func(t Type) Vec
	Record   func(fields map[string]Type) Rec
	Variant  func(fields map[string]Type) Variant
	Func     func(args []Type, ret []Type, annotations []string) Func
	Service  func(functions map[string]Func) Service
}

func NewInterface(factory IDLFactory) Service {
	return factory(IDL{
		Bool:     Bool{},
		Null:     Null{},
		Nat:      Nat{},
		Int:      Int{},
		Nat8:     Nat{Base: 8},
		Nat16:    Nat{Base: 16},
		Nat32:    Nat{Base: 32},
		Nat64:    Nat{Base: 64},
		Int8:     Int{Base: 8},
		Int16:    Int{Base: 16},
		Int32:    Int{Base: 32},
		Int64:    Int{Base: 64},
		Text:     Text{},
		Reserved: Reserved{},
		Empty:    Empty{},
		Tuple: func(ts ...Type) Tuple {
			return ts
		},
		Vec: func(t Type) Vec {
			return Vec{t}
		},
		Record: func(fields map[string]Type) Rec {
			var rec Rec
			for k, v := range fields {
				rec.Fields = append(rec.Fields, Field{
					Name: k,
					Type: v,
				})
			}
			sort.Slice(rec.Fields, func(i, j int) bool {
				return Hash(rec.Fields[i].Name).Cmp(Hash(rec.Fields[j].Name)) < 0
			})
			return rec
		},
		Variant: func(fields map[string]Type) Variant {
			var variant Variant
			for k, v := range fields {
				variant.Fields = append(variant.Fields, Field{
					Name: k,
					Type: v,
				})
			}
			sort.Slice(variant.Fields, func(i, j int) bool {
				return Hash(variant.Fields[i].Name).Cmp(Hash(variant.Fields[j].Name)) < 0
			})
			return variant
		},
		Func: func(args, ret []Type, annotations []string) Func {
			return Func{args, ret, annotations}
		},
		Service: func(methods map[string]Func) Service {
			var service Service
			for k, v := range methods {
				service.methods = append(service.methods, Method{
					Name: k,
					Func: &v,
				})
			}
			sort.Slice(service.methods, func(i, j int) bool {
				return Hash(service.methods[i].Name).Cmp(Hash(service.methods[j].Name)) < 0
			})
			return service
		},
	})
}
