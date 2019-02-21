package main

import "fmt"

func BubbleSort(a[] int) []int  {
	n := len(a)
	for i := 0; i < n - 1; i++{
		var flag bool = false  //交换标志位
		for j := 0; j < n - i - 1; j++  {
			if a[j] > a[j+1] {
				a[j],a[j+1] = a[j+1],a[j]
				flag = true
				fmt.Println(a)
			}
		}
		if (!flag) {
			break
		}
	}

	return a
}

func main() {
	var a = []int{3,5,4,6,77,3,4,5,5}
	//var a = []int{1,2,3,4,5}
	b := BubbleSort(a)
	fmt.Println(b)
}