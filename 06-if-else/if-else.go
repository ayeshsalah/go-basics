package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter number: ")
	fmt.Scanln(&num)
	if num%2 == 0{
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}


    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }


    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }

	// There is no ternary if in Go, so you’ll need to use a full if statement even for basic conditions.

}