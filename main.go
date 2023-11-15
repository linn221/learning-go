package main

import (
	"fmt"
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

func main() {
	var nums [5]int
	for i := 1; i < 5; i++ {
		nums[i] = i
	}
	fmt.Println(nums)

	primes()
	useNewPackage()
}
