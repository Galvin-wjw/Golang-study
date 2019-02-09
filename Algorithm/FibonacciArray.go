package main

import "fmt"

func FibonacciArray(n int)([] int)  {
	arr := make([]int,n)
	arr[0],arr[1] = 1,1

	for i := 2; i < n ;i ++  {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr
}

func main() {
	var n int
	fmt.Println("please input n")
	fmt.Scanln(&n)
	fmt.Println(FibonacciArray(n))
}