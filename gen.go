// +build ignore

package main

import (
	"github.com/pegn/pegn-go"
	"io/ioutil"
	"log"
)

func main() {
	rawGrammar, _ := ioutil.ReadFile("spec/grammar.pegn")
	if err := pegn.GenerateFromFiles(".", pegn.Config{
		ModulePath:     "github.com/di-wu/candid-go",
		IgnoreReserved: true,
		TypeSuffix:     "T",
	}, rawGrammar); err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully generated the pegn sub-module.")
}
