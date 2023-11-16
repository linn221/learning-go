package main

import (
	"fmt"
	"os"
	"strings"

	"rsc.io/quote"
)

func hello() {
	fmt.Println("hello hello")
}

func primes() {
	var primes = [100]int{2}
	// fmt.Println(primes)
	var index int = 0
	var isPrime bool

	for n := 2; index < len(primes); n++ {
		isPrime = true
		for _, p := range primes {

			if p == 0 {
				break
			}

			if n%p == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes[index] = n
			index++
		}
	}
	fmt.Println(primes)
}

func useNewPackage() {
	fmt.Println(quote.Go())
}

func sayName(name string) {
	fmt.Printf("Hello, %v\n", name)
}

func cycleNames(names []string, f func(string)) {
	for _, name := range names {
		f(name)
	}
}

func getInitials(n string) (string, string) {
	n = strings.ToUpper(n)
	slice := strings.Split(n, " ")
	return slice[0][:1], slice[1][:1]
}

type item struct {
	name  string
	price float64
	stock int
}

func newItem(name string, price float64) item {
	i := item{
		name:  name,
		price: price,
		stock: 0,
	}

	return i
}

func (i *item) addStock(amount int) {
	i.stock += amount
}

func (i *item) updateItem(name string, price float64, stock int) {
	i.name = name
	i.price = price
	i.stock = stock
}

func main2() {
	apple := newItem("apple", 1000)
	fmt.Println(apple)
	apple.addStock(3)
	fmt.Println(apple)
	apple.updateItem("mango", 2000, 2)
	fmt.Println(apple)
	os.Exit(1)
	i1, i2 := getInitials("tifa lockhart")
	fmt.Println(i1, i2)
	names := []string{"linn", "linnie", "lin"}
	cycleNames(names, sayName)

	var nums [5]int
	for i := 1; i < 5; i++ {
		nums[i] = i
	}
	fmt.Println(nums)

	primes()
	useNewPackage()

}
