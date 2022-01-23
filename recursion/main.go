package main

import "fmt"

func main() {
	number := 15
	fmt.Printf("fibonacci of %v is %d", number, fibonacci(number))
}

func fibonacci(number int) int {
	if number == 0 {
		return 0
	} else if number == 1 {
		return 1
	}
	return fibonacci(number-1) + fibonacci(number-2)

}
