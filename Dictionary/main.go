package main

import (
	"fmt"

	"github.com/taylorknoah/Learning-Go/Dictionary/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"First": "First Word"}

	word := "LET'S"
	definition := "not go"

	//test add fucntion
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Print(err)
	} else {
		println("INIT:", word, definition)
	}

	word = "LET'S"
	//test Search
	def2, err2 := dictionary.Search(word)
	if err2 != nil {
		fmt.Println("\n", err2)
	} else {
		fmt.Println("SEARCH: ", word, " | FOUND: ", def2)
	}

	update := "GO!!"
	//test update
	err3 := dictionary.Update(word, update)
	if err3 != nil {
		fmt.Println('\n', err3)
	} else {
		fmt.Println("UPDATE: ", word, update)
	}

	err4 := dictionary.Update("Fanny-Packs-Are-Cool", update)
	//test update error
	if err4 != nil {
		fmt.Println(err4)
	} else {
		fmt.Println("UPDATE: ", word, update)
	}

	//test delete function
	for i := 0; i < 2; i++ {
		err5 := dictionary.Delete(word)
		if err5 != nil {
			fmt.Println(err5)
		} else {
			fmt.Println("DELETE SUCCESSFUL")
		}
	}

}
