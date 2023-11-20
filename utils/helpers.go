package utils

import "fmt"

func Dump(i interface{}) {
	fmt.Printf("%#v\n", i)
}

func Sdump(i interface{}) string {
	return fmt.Sprintf("%#v\n", i)
}
