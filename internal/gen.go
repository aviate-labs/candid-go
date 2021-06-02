// +build ignore

package main

import (
	"github.com/pegn/pegn-go"
	"io/ioutil"
	"log"
)

func main() {
	rawGrammar, _ := ioutil.ReadFile("internal/grammar/grammar.pegn")
	if err := pegn.GenerateFromFiles("internal/grammar/", pegn.Config{
		ModulePath:     "github.com/di-wu/candid-go/grammar",
		IgnoreReserved: true,
		TypeSuffix:     "T",
	}, rawGrammar); err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully generated the pegn sub-module.")
}
