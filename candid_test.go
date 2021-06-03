package candid

import (
	"fmt"
)

func ExampleParseMotoko() {
	p, _ := ParseMotoko("testdata/counter.mo")
	fmt.Println(p)
	// Output:
	// service : {
	//   inc : () -> nat;
	// }
}
