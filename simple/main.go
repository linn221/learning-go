package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func connectDB() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	DB, err = gorm.Open(mysql.Open(DBURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("db connected")

}

type Loan struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Name   string `json:"name" gorm:"size:256;not null"`
	Amount uint64 `json:"amount"`
	Paid   bool   `json:"paid"`
}

func main() {
	fmt.Println("hello world")
	connectDB()
	// migration
	// DB.AutoMigrate(&Loan{})

	// create
	linn := Loan{
		Name:   "linn",
		Amount: 10000,
		Paid:   true,
	}

	DB.Create(&linn)
	fmt.Println(linn)

	// index
	var loans []Loan
	DB.Find(&loans)
	fmt.Println(loans)

	// get loan
	var result *Loan
	DB.Find(&result, 3)
	fmt.Println("loan of id 3:")
	fmt.Println(*result)

	// DB.Create(&loan)
	// fmt.Println(loan)
}
