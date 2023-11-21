package main

import (
	"fmt"
	"log"

	"github.com/linn221/go1.0/models"
)

func main() {
	models.ConnectDB()
	// migrate
	// models.Migrate()

	// create
	me := models.Student{
		Name: "linn",
		Age:  29,
		Note: "hello world",
	}
	me.Store()

	// read
	log.Println("Fetching all students")
	allStudents, err := models.IndexStudent()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(allStudents)

	log.Println("Getting student by ID 5")
	student, err := models.GetStudentByID("5")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(student)

	student.Name = "shit"
	log.Println("Deleting student of ID 5")
	student.Delete()
}
