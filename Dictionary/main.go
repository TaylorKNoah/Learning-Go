package main

import (
	"fmt"

	"github.com/taylorknoah/Learning-Go/Dictionary/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"First": "First Word"}

	def, err := dictionary.Search("Second")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(def)
	}
}
