package main

import "fmt"

func main() {

	var a = "initial"
	fmt.Println(a)

	var a1 string = "initial"
	fmt.Println(a1)

	// := syntax is shorthand for declaring and initializing a variable
	a2 := "initial"
    fmt.Println(a2)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
    fmt.Println(d)

	// Variables declared without a corresponding init ialization are zero-valued.
    var e int
    fmt.Println(e)


}
