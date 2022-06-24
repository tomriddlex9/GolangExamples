package main

import "fmt"

// main function starts here
func main() {
	// shorthand var declaration and initialisation
	i := 1
	// for loop with single condition to check whether to continue
	// the loop

	// loop structure in go is :
	// for "initialisation ";"conditional/termination checking"; "modification expression"

	// here in this loop we are using the format
	// for "conditional/termination checking"

	// this is similar to a while loop in C++
	for i <= 3 {
		fmt.Println(i)
		// modification expression , to prevent infinite loop
		i += 1

	}
	// this follows the standard for loop structure
	for j := 4; j <= 9; j++ {
		fmt.Print(j, " ")
	}
	// here this statement is used to print a new line , similar to "cout<<endl"
	fmt.Println()

	// we can also use for loop without any of the three expression to convert the for loop
	// to a while or infinite loop
	for {
		fmt.Println("loop")
		//breaks after only one iteration , but we can implement conditions for break
		break
	}
	// normal for loop
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			//we can use "continue" to skip certain steps in loop
			//here we are using continue to not print even numbers
			continue
		}
		// only odd numbers will be printed
		fmt.Print(n, " ")
	}
	fmt.Println()
}
