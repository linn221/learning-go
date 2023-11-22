package main

import "fmt"

func dump(i interface{}) {
	fmt.Printf("%+v\n", i)
}
