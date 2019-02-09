package main

import "fmt"

func intSeq() func() int  {   //This function intSeq returns another function, which we define anonymously in the body of intSeq
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())


	// another example
	var j  = 5
	a := func() (func()){
		var k  = 10
		return func() {
			fmt.Printf("i,j : %d , %d\n",k,j)
		}
	}()

	a()
	j *= 2
	a()
}
