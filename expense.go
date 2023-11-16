package main

import (
	"errors"
	"fmt"
)

type account struct {
	username string
	age      int
	password string
}

func newAccount(name string, password string, age int) account {
	acc := account{
		username: name,
		password: password,
		age:      age,
	}
	return acc
}

func (acc *account) checkLogin(username string, password string) (bool, error) {
	// validation
	if username == "" || password == "" {
		return false, errors.New("empty credentials")
	}
	if username == acc.username && password == acc.password {
		return true, nil
	}
	return false, nil
}

func (acc account) sayAge() {
	fmt.Printf("Dear %v, you are %v years old\n", acc.username, acc.age)
}

func main() {
	mgmg := newAccount("mgmg", "password", 20)
	var name string
	var password string
	fmt.Println("Enter username:")
	fmt.Scan(&name)
	fmt.Println("Enter password:")
	fmt.Scan(&password)

	status, err := mgmg.checkLogin(name, password)
	if err != nil {
		fmt.Println("illegal input")
	}
	if status {
		mgmg.sayAge()
	} else {
		fmt.Println("invalid crendentials")
	}
}
