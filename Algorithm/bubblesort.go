package main

import (
	"fmt"
	"time"
	"math/rand"
)

func BubbleSort(a[] int) []int  {
	start := time.Now()

	n := len(a)
	for i := 0; i < n - 1; i++{
		var flag bool = false  //交换标志位
		for j := 0; j < n - i - 1; j++  {
			if a[j] > a[j+1] {
				a[j],a[j+1] = a[j+1],a[j]
				flag = true
				//fmt.Println(a)
			}
		}
		if (!flag) {
			break
		}
	}

	defer func() {
		cost := time.Since(start)
		fmt.Println("cost=", cost)
	}()
	return a
}

func main() {
	var a[] int
	var i = 0
	for i < 1000{
		temp := rand.Intn(10000)
		a = append(a, temp)
		i++
	}
	fmt.Println(a)
	//a := []int{3,2,4,8,1,2,5,7}
	fmt.Println(BubbleSort(a))
}