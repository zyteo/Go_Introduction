package main

import "fmt"

// ZY, 15 Jan 2022, 8.49pm
// Learning Go through JetBrains Academy free lessons

//declaring a variable
//var variableName type // The default value
func main() {

	var firstNum int = 1
	var secondNum = 20 // Skipping the type
	// You can also use := in statements. It is a declare and assign construction:
	thirdNum := 30
	// Please remember that you can assign only a new variable with this operator. It means that you cannot assign a value to the already existing one. In this case, it will throw an error:
	fourthNum := 30
	fourthNum = 31 // Ok
	//fourthNum := 32 // Panic! This variable is already assigned
	fourthNum, fifthNum := 32, 33 // no panic
	fmt.Println(firstNum, secondNum, thirdNum, fourthNum, fifthNum)

	//	declare multiple variables
	var python, java, kotlin bool
	//	declare block of variables
	var (
		Go         bool
		Dart       bool
		Rust       bool
		TypeScript bool
	)
	fmt.Println(python, java, kotlin, Go, Dart, Rust, TypeScript)

	//	Constants
	const pi = 3.141
	const hubble int = 77
	//	block of constants
	const (
		hello = "Hello"
		e     = 2.718
	)

	//Constants have several advantages over variables. They ensure that the base value of the code remains unchanged. It doesn't matter much for small projects, but it certainly does for larger ones with numerous components contributed by multiple authors.

	//Besides, constants give the compiler a strong cue for optimization. Since the compiler knows that the value cannot be changed, it doesn't need to load the value from memory each time, and it can optimize the code to work only for the exact value of the constant.
	const KremlinTowers = 20
	//KremlinTowers = 21 // Cannot be assigned to KremlinTowers

	//The iota keyword is a handy auto-increment that generates constants. It constructs a set of related but distinct constants starting from 0.

	//It proves very useful when you have to represent some property for which different distinct values are possible. This is also known as one-hot encoding. Basically, you can use the iota keyword if you have a property and you know it's going to be one-hot encoded. Let's say it has three values: 0, 1, and 2. Then you can write it like this:
	//const (
	//	A = iota // 0
	//	B = iota // 1
	//	C = iota // 2
	//)
	//above code can be simplified to:
	const (
		A = iota // 0
		B        // 1
		C        // 2
	)

	//For example, you've got seven days of the week, and you want to define a constant for each day. In this case, you want them all to be constants with different values. However, it doesn't matter what exact value each constant has as long as Monday is different from Tuesday that is different from Wednesday, and so on. You can put it the following way:
	const (
		Monday = iota
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)

	//One of the common uses of iota is when we work with bitmasks and want to change the permission of the files, like this:
	//const (
	//	Read = 1 << iota  // << A bit operation
	//	Write
	//	Execute
	//)

	//Define a boolean variable bool with the name pizzaTime and the value true.
	var pizzaTime bool = true
	//	Define the chest variable of the string type and a value "gold".
	var chest string = "gold"
	const (
		Winter = iota
		Spring
		Summer
		Autumn
	)
}
