package models

import "fmt"

func Play() {
	var emptyStudent Student
	var container Student
	DB.Model(emptyStudent).Find(&container, 3)
	fmt.Println(emptyStudent)
	fmt.Println(container)
}
