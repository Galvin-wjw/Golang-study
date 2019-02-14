package main

import (
	"sort"
	"fmt"
)

type bylenth [] string

func (s bylenth) Len() int  {
	return len(s)
}

func (s bylenth) Swap(i,j int)  {
	s[i],s[j] = s[j],s[i]
}

func (s bylenth) Less(i,j int) bool  {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach","banana","kiwi"}
	sort.Sort(bylenth(fruits))
	fmt.Println(fruits)
}