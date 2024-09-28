// My fisher price first go code,  learning to print, variables and some string formatting, beware them sprints.
package main

import "fmt"

func main() {

	// Strings
	var nameOne string = "Mario"
	var nameTwo string = "Luigi"
	badName := "yoshi"
	fmt.Println(nameOne)

	fmt.Println(nameOne, nameTwo)

	fmt.Println(nameOne, nameTwo, badName)

	// Ints
	var ageOne int = 20
	var ageTwo = 40

	fmt.Println(ageOne, ageTwo)

	// Print
	fmt.Print("Hello, ")
	fmt.Print("World! \n")
	fmt.Print("new line \n")

	// Println
	fmt.Println("hey")

	fmt.Println("My age is", ageOne, "and my name is", nameOne)

	// Formatted strings (Printf) %_ https://pkg.go.dev/fmt
	fmt.Printf("My age is %v and my name is %v \n", ageOne, nameOne)

	fmt.Printf("My age is %q and my name is %q \n", ageOne, nameOne)
	fmt.Printf("Age is of type %T \n", ageOne)
	fmt.Printf("you scored %f points! \n", 255.55)
	fmt.Printf("you scored %0.1f points! \n", 255.55)

	// SprintF (save formatted strings)
	var str = fmt.Sprintf("My age is %v and my name is %v \n", ageOne, nameOne)
	fmt.Println("The saves string is:", str)
	// Does str respect var changes??
	ageOne = 49
	fmt.Printf("Age has changed to %v \n", ageOne)
	// Despite changing age one, str doesnt respect var change, silly
	fmt.Println("The saves string is:", str)

}
