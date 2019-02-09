package main

import "fmt"

func main() {
	var bonus float32
	var AfterTax float32

	fmt.Println("Please input your bonus:")
	fmt.Scanln(&bonus)

	switch  {
	case bonus < 3000*12:
		AfterTax = (bonus * 0.97)
		fmt.Println(AfterTax)
		break
	case bonus >= 3000*12 && bonus < 12000*12:
		AfterTax = (bonus * 0.9 + 210)
		fmt.Println(AfterTax)
		break
	case bonus >= 12000*12 && bonus < 25000*12:
		AfterTax = (bonus * 0.8 + 1410)
		fmt.Println(AfterTax)
		break
	default:
		fmt.Println("Your bonus is over cal-capacity")


	}
}
