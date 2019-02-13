package main

import "fmt"

func main() {
	var container []int = []int{1,2,3}

	value, ok := interface{}(container).([]int)

	fmt.Println(value,ok)


}

