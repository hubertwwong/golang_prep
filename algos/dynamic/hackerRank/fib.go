package main

import "fmt"

func fib(val int) (int) {
	if val == 0 {
		return 0
	} else if val == 1 {
		return 1
	}

	return fib(val-1) + fib(val-2)
}

func main() {
	fmt.Println(fib(0))
	fmt.Println(fib(1))
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(4))
	fmt.Println(fib(5))
}