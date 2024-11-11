package main

import "fmt"

func main() {
	sum := 0

	for i := 0; i <= 10; i++ {
		sum += i
		fmt.Println("sum is ", sum)
	}

	for sum > 10 {
		sum -= 10
		fmt.Println("decrement", sum)
	}

}
