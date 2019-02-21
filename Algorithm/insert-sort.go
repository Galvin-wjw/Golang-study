package main

import "fmt"

func InsertSort(a[] int) []int {
	lenth := len(a)

	for i := 1; i < lenth; i++ {
		base_value := a[i]
		j := i - 1
		for ; j >=0 ;j--  {
			if a[j] > base_value {
				a[j+1] = a[j]
			}else{
				break
			}
		}
		a[j + 1] = base_value

	}
	return a
}

func main() {
	a := []int{3,2,4,8,1,2,5}
	fmt.Println(InsertSort(a))
}
