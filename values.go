package main

import "fmt"

func main() {
	//these are comments
	// works like in c and cpp
	/* A multi line comment
	just like C++
	*/
	fmt.Println("go" + "lang") // string concatenation
	fmt.Println("1+1= ", 1+1)
	fmt.Println("7.0/3.0 = ", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println("This gets printed in a newline because of Println")
	fmt.Print("this & ")
	fmt.Print("this gets printed in same line, because of Print")
	fmt.Println()
	fmt.Println("Now a new line")
	fmt.Println(!true)

}
