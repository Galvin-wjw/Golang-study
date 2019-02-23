package main

import (
	"fmt"
	"math/rand"
	"time"
)

func InsertSort(a[] int) []int {
	start := time.Now()

	lenth := len(a)

	for i := 1; i < lenth; i++ {
		base_value := a[i]
		j := i - 1
		for ; j >=0 ;j--  {
			if a[j] > base_value {
				//这里交换或者移位都可以实现算法，空间时间开销不一样
				a[j+1] = a[j]
				//a[j+1],a[j] = a[j],a[j+1]
			}else{
				break
			}
		}
		a[j + 1] = base_value
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
	fmt.Println(InsertSort(a))
}
