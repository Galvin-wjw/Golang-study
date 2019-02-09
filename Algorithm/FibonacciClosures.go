package main

import "fmt"

func FibonacciCloures() (func() int) {
	seq1,seq2 := 0,1
	return func() int {
		seq1 , seq2 = seq2,(seq1+seq2)
		return seq1
	}
}

func main() {
	const LTM  = 200
	f := FibonacciCloures()
	var arr [LTM] int
	for i := 0; i<LTM;i++ {
		arr[i] = f()
	}
	fmt.Println(arr)
}
