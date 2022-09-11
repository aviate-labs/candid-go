package marshal

import (
	"fmt"

	"github.com/aviate-labs/candid-go/idl"
	"github.com/aviate-labs/candid-go/typ"
	"github.com/aviate-labs/principal-go"
)

type Optional[T any] struct {
	Value *T
}

func (o *Optional[T]) ofType(v any) error {
	if v, ok := v.(T); ok {
		o.Value = &v
		return nil
	}
	return fmt.Errorf("invalid value match: %s", v)
}

func optionalOfType[T any](dv any, value any) (Optional[T], error) {
	v, ok := value.(*Optional[T])
	if !ok {
		return Optional[T]{}, fmt.Errorf("invalid type match: %s", value)
	}
	return *v, v.ofType(dv)
}

func optionalOf(t idl.Type, dv any, value any) (any, error) {
	switch t.(type) {
	case *idl.BoolType:
		return optionalOfType[bool](dv, value)
	case *idl.NatType:
		return optionalOfType[typ.Nat](dv, value)
	case *idl.IntType:
		return optionalOfType[typ.Int](dv, value)
	case *idl.FloatType:
		return optionalOfType[float64](dv, value)
	case *idl.TextType:
		return optionalOfType[string](dv, value)
	case *idl.Principal:
		return optionalOfType[principal.Principal](dv, value)
	}
	return Optional[any]{
		Value: &value,
	}, nil
}
