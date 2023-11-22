package main

import (
	"fmt"
	"linn221/gorm1.1/models"
	"log"
	"os"
)

func createLoan() {
	input := models.Loan{
		// ID:     12,
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
	loans, err := models.ListLoans()
	if err != nil {
		log.Fatal(err)
	}

	dump(loans)
}

func getLoan() {
	loan, err := models.FindLoanByID(3)
	if err != nil {
		log.Fatal(err)
	}
	dump(loan)
}

func updateLoan() {
	loan := models.Loan{
		ID:     3,
		Name:   "zaw zaw",
		Amount: 4000,
	}

	_, err := loan.Update()
	if err != nil {
		log.Fatal(err)
	}
	dump(loan)
}

func deleteLoan() {
	fmt.Println("deleting loan of id 3")
	loan, err := models.FindLoanByID(3)
	if err != nil {
		log.Fatal(err)
	}
	loan.Delete()
	indexLoans()
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
