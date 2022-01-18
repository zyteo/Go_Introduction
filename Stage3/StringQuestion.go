package main

import "fmt"

// ZY, 18 Jan 2022, 8.28am
// Learning Go through JetBrains Academy free lessons
//String formatting question

func main() {
	fmt.Printf("%[1]d %[1]d %[1]d", 100)

	x := "Golang"
	y := '8'
	fmt.Printf("Do you know that %s is %c years old", x, y)

	fmt.Printf(
		`
%[1]s %[1]s %[2]s
%[2]s %[2]s %[1]s
%[1]s %[2]s %[1]s
%[2]s %[1]s %[2]s
%[2]s %[2]s %[2]s
        `, "White", "Black")
}
