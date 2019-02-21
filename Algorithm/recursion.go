package main

import "fmt"

//n个阶梯，每次上 1 阶或者2阶，一共有多少种方法
// Space Complexity：O(n)
// Time Complexity:
func stairs(n int)  int{
	if n == 1{
		return 1
	}
	if n == 2 {
		return 2
	}else {
		return stairs(n-1)+stairs(n-2)
	}
}

func main() {
	fmt.Println(stairs(7))
}