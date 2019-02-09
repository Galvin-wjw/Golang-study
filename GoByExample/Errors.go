package main

import (
	"qiniupkg.com/x/errors.v7"
	"fmt"
)

func f1(arg int) (int,error)  {
	if arg == 42{
		return -1,errors.New("can not work with 42")  //errors.New 构造一个使用给定的错误信息的基本 error 值。

	}

	return arg + 3,nil
}

type argError struct {
	arg int
	prob string
}

//It’s possible to use custom types as errors by implementing the Error() method
func (e *argError) Error() string {
	return fmt.Sprintf("%d  - %s",e.arg,e.prob)
}

func f2(arg int)(int ,error){
	if arg == 42{
		return -1, &argError{arg,"can`t work with"}
	}
	return arg + 3,nil
}

func main() {
	for _, i := range []int{7,42}{
		if r,e := f1(i); e != nil{
			fmt.Println("f1 failed:",e)
		}else {
			fmt.Println("f1 worked:",r)
		}
	}

	for _, i := range []int {7,42}{
		if r,e := f2(i); e != nil{
			fmt.Println("f2 failed:",e)
		}else {
			fmt.Println("f2 worked:",r)
		}
	}

	_, e := f2(42)

	if ae,ok := e.(*argError) ; ok{
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

}