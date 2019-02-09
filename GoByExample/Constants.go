//A numeric constant has no type until it’s given one, such as by an explicit cast.

package main

import (
	"fmt"
	"math"
)

const s string  = "constant"

func main() {
	fmt.Println(s)

	const  n  = 500000

	const d = 3e20 / n
	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
