package main

import "fmt"


func MergeSort(a[] int)  {
	lenth := len(a)
	if lenth <= 1{
		return
	}

	mergesort(a,0,lenth-1)
}


//归并函数
func merge(a[] int,first,mid,last int){
	tmpArr := make([]int,last-first + 1)

	i := first
	j := mid + 1
	k := 0

	for ; i <= mid && j <= last; k++{
		if a[i] < a[j] {
			tmpArr[k] = a[i]
			i ++
		}else {
			tmpArr[k] = a[j]
			j++
		}
	}
	//fmt.Println(tmpArr)

	for ; i <= mid; i++{
		tmpArr[k] = a[i]
		k ++
	}

	for ; j <= last; j++  {
		tmpArr[k] = a[j]
		k ++
	}

	copy(a[first:last+1],tmpArr)

}

// 归并排序
func mergesort(a [] int,first ,last int) {
	//lenth := len(a)
	if first >= last{
		return
	}

	m := (first + last)/2
	mergesort(a,first,m)
	mergesort(a,m+1,last)
	merge(a,first,m,last)

	//s1 := make([]int,lenth/2)
}

func main() {
	var a = [] int {7,5,8,6,0,2,4,3}
	MergeSort(a)
	fmt.Println(a)
}