package main

import "fmt"
// Recursion
func Fibonacci(n int) (result int) {
	if n <= 1 {
		result = 1
	}else {
		result = Fibonacci(n-2) + Fibonacci(n-1)
	}
	return
}

func main() {
	var num int
	var array []int
	fmt.Println("please input num:")
	fmt.Scanln(&num)

	for i := 0; i < num; i ++{
		results := Fibonacci(i)
		array = append(array,results)
	}
	fmt.Println(array)
}





