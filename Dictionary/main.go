package main

import (
	"fmt"

	"github.com/TaylorKNoah/Learning-Go/Dictionary/mydict"
)

func main() {

	fmt.Println("Hello and welcome to the dictionary 9000.")
	fmt.Println("Here you will be able to create and modify a dictionary. Wow.")

	input := 1
	ok := true
	valid := false
	dictionary := mydict.Dictionary{}

	for ok == true {

		for valid == false {
			if input < 1 || input > 4 {
				fmt.Println("Enter an int from 1-4")
			}

			fmt.Println("1) Print Dictionary Conents")
			fmt.Println("2) Add Word")
			fmt.Println("3) Delete Word")
			fmt.Println("4) Search Word")
			fmt.Println("5) Update Word")
			fmt.Println("0) Exit")
			fmt.Print("\nEnter Choice (0-5): ")

			_, err := fmt.Scanf("%d\n", &input)
			for err != nil {
				fmt.Println(err)
				fmt.Println("Enter an integer please")
				_, err = fmt.Scanf("%d\n", &input)
			}

			if input < 0 {
				valid = false
			} else if input > 5 {
				valid = false
			} else {
				valid = true
			}
		}

		//send input to filter
		if input == 0 {
			ok = false
		} else {
			menu_filter(input, dictionary)
			ok = again()
			if ok == true {
				input = 1
				valid = false
			}
		}
	}

}

func again() bool {

	var choice string
	fmt.Println("Again? Y/N")
	fmt.Scanf("%s\n", &choice)

	if ok := true; ok == false {
		fmt.Println(choice)
	}

	if choice != "Y" && choice != "N" {
		fmt.Println("Incorrect Input")
		return again()
	}

	if choice == "Y" {
		return true
	}

	return false

}

//filters input choice to correct function call
func menu_filter(input int, d mydict.Dictionary) {
	switch input {
	case 1:
		d.Display()
	case 2:
		add_to_dict(d)
	case 3:
		delete_from_dict(d)
	case 4:
		search_in_dict(d)

	case 5:
		update_dict(d)
	}

}

func add_to_dict(d mydict.Dictionary) {

	//get word and def
	var (
		word, define string
	)

	fmt.Println("Enter the Word:")
	fmt.Scanln(&word)
	fmt.Println("Enter the definition:")
	fmt.Scanln(&define)

	//test add fucntion
	err := d.Add(word, define)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Add Successful")
	}
}

func delete_from_dict(d mydict.Dictionary) {
	var word string

	fmt.Println("Enter the Word to Delete:")
	fmt.Scanln(&word)

	//test delete function
	err5 := d.Delete(word)
	if err5 != nil {
		fmt.Println(err5)
	} else {
		fmt.Println("Delete Successful")
	}
}

func search_in_dict(d mydict.Dictionary) {
	var word string

	fmt.Println("Enter the Word to Search:")
	fmt.Scanln(&word)

	//test Search
	def2, err2 := d.Search(word)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("SEARCH: ", word, " | FOUND: ", def2)
	}
}

func update_dict(d mydict.Dictionary) {
	var (
		word, update string
	)

	fmt.Println("Enter the Word:")
	fmt.Scanln(&word)
	fmt.Println("Enter the new definition:")
	fmt.Scanln(&update)

	//test update
	err3 := d.Update(word, update)
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println("Update Successful")
	}
}

/*
func SCRAP_CODE_NOT_REAL_FUNCTION() {

	word := "LET'S"
	definition := "not go"

	//test add fucntion
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println("INIT:", word, definition)
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
*/
