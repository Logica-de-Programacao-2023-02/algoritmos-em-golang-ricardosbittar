package main

import "fmt"

func main() {
	var x int
	fmt.Println("numero:")
	fmt.Scanln(&x)

	if x%3 == 0 {
		fmt.Println("Fizz")
	}
	if x%5 == 0 {
		fmt.Println("Buzz")

	}
	if x%3 == 0 && x%5 == 0 {
		fmt.Println("FizzBuzz")
	}

}
