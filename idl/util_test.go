package idl_test

import (
	"fmt"
	"reflect"

	"github.com/allusion-be/candid-go/idl"
)

func test(types []idl.Type, args []interface{}) {
	e, err := idl.Encode(types, args)
	if err != nil {
		fmt.Println("enc:", err)
		return
	}
	fmt.Printf("%x\n", e)

	ts, vs, err := idl.Decode(e)
	if err != nil {
		fmt.Println("dec:", err)
		return
	}
	if !reflect.DeepEqual(ts, types) {
		fmt.Println("types:", types, ts)
	}
	if !reflect.DeepEqual(vs, args) {
		fmt.Println("args", args, vs)
	}
}
