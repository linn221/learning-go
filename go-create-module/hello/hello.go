// run me to discover greetings module
// go mod edit -replace example.com/greetings=../greetings

package main

import (
	"fmt"
	"greetings"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

/*
	var name string
	fmt.Println("Enter your name:")
	fmt.Scanf()
	*/

	greeting, err := greetings.Hello("Linn Georgie")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(greeting)

	names := [] string {
		"Linn Georgie",
		"Kyaw Kyaw",
		"Zaw Zaw",
	}
	greetings, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(greetings)
}
