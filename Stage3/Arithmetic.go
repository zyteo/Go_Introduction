package main

// ZY, 17 Jan 2022, 7.43pm
// Learning Go through JetBrains Academy free lessons
//Arithmetic theory

func main() {
	x + y // The result is the sum of x and y.
	x - y // The result is the difference of x and y.
	-x    // x will change its sign.
	+x    // Return identity of x.
	x * y // Return the product of x and y.
	x / y // Return the quotient of x and y.
	x % y // Return the remainder of x / y.

	result := 10.1 + 1.234     // result = 11.334
	result2 := -5.6789 + 19.23 // result = 13.551100000000002

	var a = 8
	var b = 13.0
	result := a + b // mismatched types int and float64,
	// IDE will warn you, and if you try to compile,
	// it will produce a compile error

	//Unary arithmetic operations

	//Unary mathematical expressions consist only of one component. In Go, we can use the + and - signs paired with a value to return the value's identity or change the sign of the value. For example:
	//x := 3.3 // positive values are indicated by the plus sign.

	//y := +x  // y = 3.3
	//
	//x = -30 //  using the '+' symbol with a negative value will also return a negative value
	//y = +x
	//
	//x = 5
	//y = -x // y = -5. Using the '-' sign changes the sign of the value -> y = -5
	//
	//x = -20
	//y = -x // y = 20

	//Multiplication, division, and modulo

	//Go uses the * and / signs to perform multiplication and division, respectively. Let's take a look at multiplication:
	//x := 42.3
	//y := 13.1
	//result := x * y // result = 554.13

	//Division results in Go will vary according to the numeric type of the values. When dividing int values, the result will always be rounded down to the nearest integer; to get the result as a decimal number, we can wrap our values within float32() or float64(). For example:
	// Dividing 'int' values
	x := 13
	y := 2
	result := x / y // result = 6

	a := -100
	b := 3

	result2 := a / b // result2 = -33

	// Casting 'int' values as 'float64'
	x = 20
	y = 3
	result3 := float64(x) / float64(y) // result3 = 6.666...7

	//Finally, the % sign is the modulo operator. It returns the remainder rather than the quotient of a division and can only be used with the int type values.

	//As an example, here is a modulo operation:
	x := 100
	y := 15
	result := x % y // result = 10

	result := 20 + 20*5 // In this case, multiplication takes place first,
	// thus result = 120

	result := (10 + 10) * 5 // In this case,
	// the operation within the parentheses takes place first,
	// thus result = 100

	//Assignment operators

	//The assignment operation is performed with the equals = sign. The way it works is by assigning the value on the right to the variable on the left. For example:

	v := 23 // the integer 23 is assigned to the variable v

	//There are also compound assignment operators that perform an arithmetic operation on a variable's value and then assign the result of the operation to the variable. In straightforward terms, they combine an arithmetic operation with an assignment. Consider the following example:

	x := 5

	// Compound addition
	x += 1 // equivalent to x = x + 1 -> x = 6

	// Compound subtraction
	x -= 10 // equivalent to x = x - 10 -> x = -4

	// Compound multiplication
	x *= 4 // equivalent to x = x * 4 -> x = -16

	// Compound division
	x /= 2 // equivalent to x = x / 2 -> x = -8

	// Compound modulo
	x %= 3 // equivalent to x = x % 3 -> x = -2 (final result.)

	//Increment and decrement operators

	//The increment and decrement operator in Golang is the statement that increments or decrement the operand to the left of it by 1.

	num := 0
	num++ // num = 1
	num++ // num = 2
	num-- // num = 1
	num-- // num = 0

	//Despite the convenience of the syntax, there is something we should point out about these two special operators.

	//They are statements, not expressions; in computer science, a statement never returns a value. Since it cannot be evaluated as a value, we can't pass it to another function as an argument, or assign it directly to another variable, or combine these operators inside another expression or statement. In short, this means that the increment/decrement operators in Go are as simple as it gets:
	num1 := 1
	num1++ // Works!
	num2 := num1++ // Doesn't work!
	num3 := (num1++)++ // Doesn't work!
	num4 := 1 + num1++ // Doesn't work!

}
