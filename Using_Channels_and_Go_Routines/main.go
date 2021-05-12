package main

// Issues
//     Goal: output 2 true and 3 false values from line 27
//     Output: sometimes gives 1 true and 4 false
//     Seens like some threads

import (
	"fmt"
	"os"
	"time"
)

func main() {

	///////////////////////////////////////////////////////////////////////
	///                 A more complex Test  of my own design           ///
	///////////////////////////////////////////////////////////////////////

	/*
		names := [5]string{"Valid", "Invalid", "Valid", "jkdksl", "lkj"}

		c := make(chan bool)
		for j, name := range names {

			go verify(name, c, j)

		}

		results := [len(names)]bool{true, true, true, true, true}

		f, err := os.OpenFile("Learning-Go/Using_Channels_and_Go_Routines/channel_output.txt", os.O_APPEND|os.O_CREATE|os.O_RDONLY, 0644)
		f.WriteString("----- NEW TEST SEGMENT -----\n\n")
		if err != nil {
			log.Println(err)
		}

		for i := range results {
			result := <-c
			writestring := fmt.Sprintf("%t%d%s", result, i, "\n")
			f.WriteString(writestring)
			fmt.Println(i, result)
		}
		f.WriteString("\n-----END TEST SEGMENT -----\n\n")
		f.Close()
	*/

	///////////////////////////////////////////////////////////////////////
	///                 A more Simple Test (nomad design)               ///
	///////////////////////////////////////////////////////////////////////

	people := [5]string{"Taylor", "Lucy", "Sophie", "Duke", "Fletcher"}

	s := make(chan string)
	for _, person := range people {
		go checker(person, s)
	}

	for i := 0; i < len(people); i++ {
		fmt.Println(<-s)
	}
}

func verify(name string, c chan bool, i int) {

	f, err := os.OpenFile("Learning-Go/Using_Channels_and_Go_Routines/channel_output.txt", os.O_APPEND|os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}

	output := fmt.Sprintf("%s%d%s%s%s", "V: ", i, " ", name, "\n")
	f.WriteString(output)

	//f.Close()
	if name == "Valid" {
		c <- true
	}
	c <- false
}

func checker(name string, s chan string) {
	time.Sleep(time.Second * 2)

	if name == "Taylor" {
		s <- name + " is a person."
	} else {
		s <- name + " is a pet."
	}
}
