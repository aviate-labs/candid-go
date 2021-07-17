package idl

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/allusion-be/leb128"
)

type Float struct {
	f    *big.Float
	base uint8
	primType
}

func Float32() *Float {
	return &Float{
		f:    new(big.Float),
		base: 32,
	}
}

func Float64() *Float {
	return &Float{
		f:    new(big.Float),
		base: 64,
	}
}

func NewFloat32(f float32) *Float {
	return &Float{
		f:    big.NewFloat(float64(f)),
		base: 32,
	}
}

func NewFloat64(f float64) *Float {
	return &Float{
		f:    big.NewFloat(f),
		base: 64,
	}
}

func (f *Float) Decode(r *bytes.Reader) error {
	f64, err := readFloat(r, int(f.base/8))
	if err != nil {
		return err
	}
	*f.f = *f64
	return nil
}

func (f Float) EncodeType() []byte {
	floatXType := new(big.Int).Set(big.NewInt(floatXType))
	if f.base == 64 {
		floatXType.Add(floatXType, big.NewInt(-1))
	}
	bs, _ := leb128.EncodeSigned(floatXType)
	return bs
}

func (f Float) EncodeValue() []byte {
	return writeFloat(f.f, int(f.base/8))
}

func (Float) Name() string {
	return "float"
}

func (f Float) String() string {
	return fmt.Sprintf("float%d: %v", f.base, f.f)
}
