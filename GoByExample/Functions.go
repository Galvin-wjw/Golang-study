package main

import "fmt"

func plus(a int, b int) float32  {
	return float32(a+b)
}

func plusPlus(a,b,c int) int  {
	return a+b+c
}

func main() {
	res := plus(1,2)
	fmt.Println("1+2=",res)

	res2 := plusPlus(1,2,3)
	fmt.Println("1+2+3=",res2)
}