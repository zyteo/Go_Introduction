package main

import "fmt"

// ZY, 18 Jan 2022, 7.57pm
// Learning Go through JetBrains Academy free lessons
//Control statements questions

func main() {
	var number int
	fmt.Println("Write a number")
	fmt.Scanf("%d", &number)

	// Write your code here.
	switch {
	case number > 0:
		fmt.Println("Positive!")
	case number == 0:
		fmt.Println("Zero!")
	case number < 0:
		fmt.Println("Negative!")
	}

	var emoji = "❓"
	fmt.Scanf("%s", &emoji)

	// Please do not delete the emojis after the case statement, just fix the code errors.
	// Also please do not delete or change the text within the fmt.Println functions!
	switch emoji {
	case "⭕":
		fmt.Println("You haved picked the circle. Not the easiest shape!")
	case "🔺":
		fmt.Println("You haved picked the triangle. The easiest shape!")
	case "⭐":
		fmt.Println("You haved picked the star. Easier than circle, harder than triangle.")
	case "☂️":
		fmt.Println("You haved picked the umbrella. This is the hardest shape! GOOD LUCK.")
	default:
		fmt.Println("You haved picked an invalid emoji. Please try again or be eliminated from the game.")
	}

	var score int
	fmt.Scanf("%d", &score)

	// Write the code required to validate the student's score here.
	switch {
	case score >= 71:
		fmt.Println("Passed!")
	default:
		fmt.Println("Failed!")
	}

	//Keanu wants to create a Go program that recommends one of his movies to the user, based on their age:
	//    Age <= 14 | "Toy Story 4"
	//    Age 15 - 18 | "The Matrix"
	//    Age 19 - 25 | "John Wick"
	//    Age 26 - 35 | "Constantine"
	//    Age > 35 | "Speed"

	var age int
	fmt.Scanf("%d", &age)

	// Code your switch or if...else-if statement here.
	if age > 35 {
		fmt.Println("Speed")
	} else if 25 < age && age < 35 {
		fmt.Println("Constantine")
	} else if 18 < age && age < 25 {
		fmt.Println("John Wick")
	} else if 14 < age && age < 19 {
		fmt.Println("The Matrix")
	} else {
		fmt.Println("Toy Story 4")
	}

}
