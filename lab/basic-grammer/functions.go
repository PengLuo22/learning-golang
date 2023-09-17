package main

import "fmt"

func main() {

	// 单个返回值
	s := sum(2, 3)
	fmt.Println("2+3 =", s)

	// 多个返回值
	a, b, c := vals()
	fmt.Println("a= ", a, "b= ", b, "c= ", c)

	_, d, _ := vals()
	fmt.Println("d= ", d)

	// 可变参
	variadicSum(1, 2)
	variadicSum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	variadicSum(nums...)

}

func sum(x, y int) int {
	return x + y
}

func vals() (a int, b int, c int) {
	a = 1
	b = 2
	c = 3
	return
}

func variadicSum(nums ...int) {
	fmt.Print(nums, " ")

	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)

}
