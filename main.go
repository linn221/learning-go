package main

import (
	"fmt"
	"linn221/gorm1.1/models"
	"log"
	"os"
)

func createLoan() {
	input := models.Loan{
		Name:   "Linn Georgie",
		Amount: 50000,
	}

	_, err := input.Store()
	if err != nil {
		log.Fatal(err)
	}
	dump(input)
}

func indexLoans() {

}

func getLoan() {

}

func updateLoan() {

}

func deleteLoan() {

}

func main() {
	fmt.Println("hello world!")
	models.ConnectDatabase()
	command := os.Args[1]
	switch command {
	case "c":
		createLoan()
	case "i":
		indexLoans()
	case "g":
		getLoan()
	case "u":
		updateLoan()
	case "d":
		deleteLoan()
	}

}
