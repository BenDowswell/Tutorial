// My fisher price first go code,  learning to print, variables and some string formatting, beware them sprints.
package main

import (
	"fmt"
	"strings"
)

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

	// Arrays, slices and ranges
	// Longhand: var ages [3]int = [3]int{20, 25, 30}
	var ages = [3]int{20, 25, 30}
	names := [4]string{"Yoshi", "Mario", "Bowswer", "Peach"}

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// Slices (dynamic (traditional arrays))
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

	// Slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)

	// strings example imported package
	greeting := "hello there"
	// Should return true
	fmt.Println(strings.Contains(greeting, "hello"))
	// Should return false
	fmt.Println(strings.Contains(greeting, "hello!"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "ll"))
	// original value
	fmt.Println("Original string value :", greeting)

}
