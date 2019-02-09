package main

import (
	"fmt"
	"time"
)

func main() {
	i:=2
	fmt.Println("write",i,"as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday,time.Sunday:
		fmt.Println("It`s the weekend")
	default:
		fmt.Println("It`s a weekday")
	}

	t := time.Now()
	switch  {
	case t.Hour() < 12:
		fmt.Println("it is before noon")
	default:
		fmt.Println("it is after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am an int")
		default:
			fmt.Println("Donot know type %T\n",t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}