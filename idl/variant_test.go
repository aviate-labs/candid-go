package idl_test

import "github.com/aviate-labs/candid-go/idl"

func ExampleVariantType() {
	result := map[string]idl.Type{
		"ok":  new(idl.TextType),
		"err": new(idl.TextType),
	}
	test_([]idl.Type{idl.NewVariantType(result)}, []interface{}{idl.FieldValue{
		Name:  "ok",
		Value: "good",
	}})
	test_([]idl.Type{idl.NewVariantType(result)}, []interface{}{idl.FieldValue{
		Name:  "err",
		Value: "uhoh",
	}})
	// Output:
	// 4449444c016b029cc20171e58eb4027101000004676f6f64
	// 4449444c016b029cc20171e58eb402710100010475686f68
}
