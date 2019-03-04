package main

import "fmt"

//二分查找 n(logn)
func binary_search(arr[] int,value int) int {

	lenth := len(arr)
	if lenth == 0 {
		return -1
	}
	var low = 0
	var high = lenth - 1
	var mid int

	for low <= high{
		mid = low + (high -low)>>1
		if (arr[mid] == value) {
			return mid
		}else if (arr[mid] > value){
			high = mid - 1
		}else {
			low = mid + 1
		}
	}
	return -1
}

//递归
func binary_search_recursion(arr[] int,low,high,value int)  int{
	lenth := len(arr)
	if low > high || lenth == 0{
		return -1
	}


	mid := low + (high - low) >> 1
	if arr[mid] == value {
		return mid
	}else if arr[mid] > value {
		return binary_search_recursion(arr,low,mid-1,value)
	}else {
		return binary_search_recursion(arr,mid+1,high,value)
	}
}


//二分查找，第一个等于n的数字
func bsearch_fisrt_No(arr[] int,value int ) int  {
	lenth := len(arr)
	low := 0
	high := lenth - 1

	for(low <= high){
		mid := low + (high - low) >> 1
		if  arr[mid] > value{
			high = mid - 1
		}else if arr[mid] < value{
			low = mid +1
		}else {
			if (mid == 0 || arr[mid -1] != value) {
				return mid
			}else {
				high = mid - 1
			}
		}

	}
	return -1
}

//最后一个等于n的数字
func bsearch_last_No(arr[] int,value int) int {
	lenth := len(arr)
	low := 0
	high := lenth - 1

	for (low <= high) {
		mid := low + (high - low) >> 1
		if arr[mid] > value {
			high = mid - 1
		}else if arr[mid] < value{
			low = mid + 1
		}else {
			if(mid == lenth - 1 || arr[mid + 1] != value){
				return mid
			}else {
				low = mid + 1
			}
		}
	}
	return -1
}

func main() {
	var arr = []int{2,2,3,4,5,6,7,8,8,8,8,8,9,10}
	value := binary_search(arr,2)
	value2 := binary_search_recursion(arr,0,8,3)
	value3 := bsearch_fisrt_No(arr,2)
	value4 := bsearch_last_No(arr,8)
	fmt.Println(value,value2,value3,value4)
}