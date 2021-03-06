package main

import (
	"fmt"
	"math"
)

const str string = "constant"

func main() {
	fmt.Println(str)

	const n float64 = 500000000
	fmt.Println(n)
	fmt.Println(int64(n))

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}