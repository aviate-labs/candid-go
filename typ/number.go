package typ

import "math/big"

type Int struct {
	i *big.Int
}

func NewInt[number Integer](i number) Int {
	return Int{i: big.NewInt(int64(i))}
}

func NewIntFromString(n string) Int {
	bi, ok := new(big.Int).SetString(n, 10)
	if !ok {
		panic("number: invalid string: " + n)
	}
	return Int{bi}
}

func (i Int) BigInt() *big.Int {
	return i.i
}

func (i Int) String() string {
	return i.i.String()
}

type Natural interface {
	uint | uint64 | uint32 | uint16 | uint8
}

type Nat struct {
	n *big.Int
}

func NewNat[number Natural](n number) Nat {
	return Nat{new(big.Int).SetUint64(uint64(n))}
}

func NewNatFromString(n string) Nat {
	bi, ok := new(big.Int).SetString(n, 10)
	if !ok {
		panic("number: invalid string: " + n)
	}
	if bi.Sign() < 0 {
		panic("number: negative nat")
	}
	return Nat{bi}
}

func (n Nat) BigInt() *big.Int {
	return n.n
}

func (n Nat) String() string {
	return n.n.String()
}

type Integer interface {
	int | int64 | int32 | int16 | int8
}
