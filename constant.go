package main

import (
	"fmt"
	"math"
)

// const is used to declare constant values that cant be
// changed during the execution of program

// cs is a constant global string value
const cs string = "constant string"

// main function
func main() {
	fmt.Println(cs)
	// a numeric constant has no type until it is given one
	// such as by an explicit conversion (type casting)
	const n = 500000000
	const d = 3e20 / n
	// d without a specific type
	fmt.Println(d)
	// type casting d to a 64bit integer
	fmt.Println(int64(d))

	// here the Sin function expects a 64bit floating point number
	// so "d" is type casted into a float
	fmt.Println(math.Sin(n))

}
