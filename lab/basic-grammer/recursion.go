package main

import "fmt"

func main() {
	fmt.Println(fact(7))

	fib := fact()

}

func fact(n int) int {
	if 0 == n {
		return 1
	}
	return n * fact(n-1)
}
