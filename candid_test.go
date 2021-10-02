package candid

import (
	"fmt"
	"io/ioutil"
)

func ExampleParseDID() {
	raw, _ := ioutil.ReadFile("testdata/counter.did")
	p, _ := ParseDID(raw)
	fmt.Println(p)
	// Output:
	// service : {
	//   inc : () -> nat;
	// }
}
