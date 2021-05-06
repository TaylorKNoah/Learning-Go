package main

import "fmt"

func main() {

	names := [2]string{"Valid", "Invalid"}

	c := make(chan bool)
	for _, name := range names {

		go verify(name, c)

	}

	results := [len(names)]bool{true, true}

	for i := range results {
		results[i] = <-c
	}

	fmt.Println(results)

}

func verify(name string, c chan bool) {
	if name == "Valid" {
		c <- true
	}

	c <- false
}
