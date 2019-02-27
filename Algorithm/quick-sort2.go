package main

import "fmt"

func quick_sort(arr[] int)  {
	lenth := len(arr)
	quick_sort_c(arr,0,lenth-1)
}

func quick_sort_c(arr[] int, first,last int)  {
	if first >= last{
		return
	}

	pivot := partition(arr,first,last)

	quick_sort_c(arr,first,pivot-1)
	quick_sort_c(arr,pivot+1,last)

	
}

func partition(arr[] int,first,last int) int {
	pivot := arr[last]
	i := first

	for j := first; j <= last -1 ;j++  {
		if arr[j] < pivot{
			arr[i],arr[j] = arr[j] ,arr[i]
			i ++
		}
	}

	arr[i],arr[last] = arr[last],arr[i]
	return i
}

func main() {
	var arr = []int{2,4,54,65,2,3,8,5,9,7}
	quick_sort(arr)
	fmt.Println(arr)
}