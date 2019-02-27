package main

import (
	"fmt"
)

//桶排序，
func assign_bucket(a[] string) [] string{

	lenth := len(a)

	var cap_letter []string
	var lower_letter []string
	var number []string

	var final []string
	//var ca_letter = make([]string,lenth)

	for i:=0 ; i < lenth; i++{
		if a[i] <= "z" && a[i] >= "a"{
			lower_letter = append(lower_letter,a[i])
		}
		if a[i] <= "Z" && a[i] >= "A"{
			cap_letter = append(cap_letter,a[i])
		}
		if a[i] <= "9" && a[i] >= "0" {
			number = append(number,a[i])
		}

	}
	BubbleSort_bucket(cap_letter)
	BubbleSort_bucket(lower_letter)
	BubbleSort_bucket(number)

	final = append(final,lower_letter...)
	final = append(final, number...)
	final = append(final,cap_letter...)

	return final
}

func BubbleSort_bucket(a[] string) []string  {


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
	return a
}



func main() {
	var arr = []string{"3","4","5","2","a","g","H","K","2","A","b"}
	fmt.Println(assign_bucket(arr))

}
