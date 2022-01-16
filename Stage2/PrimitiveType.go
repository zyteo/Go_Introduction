package main

// ZY, 16 Jan 2022, 3.53pm
// Learning Go through JetBrains Academy free lessons
//Primitive types

func main() {

	//Integers

	//In Go, there are the following types of integers: uint8,uint16, uint32, uint64, int8, int16,int32, and int64. 8, 16, 32, and 64 tell us how many bits each type uses. uint means "unsigned integer", whilst int means "signed integer". An unsigned integer can only take positive values (or zero).
	var number1 int
	number1 = 100

	var number2 int = 100

	//	The type of the variable is determined at compile-time based on the value. All these variables get the int data type when compiled, despite the fact that we don't explicitly specify this:
	var number3 = 100
	number4 := 100

	//Floats

	//Floating-point numbers are numbers that contain a real part (real numbers), e.g. 1.234, 123.4, 0.00001234). The way they are represented in the computer is quite difficult and is not particularly necessary for their use. For now, we just have to remember the following:

	//Floating-point numbers are inaccurate. There are cases when a number can't be represented at all. For example, the result of calculating 1.01 - 0.99 will be 0.020000000000000018 – a number very close to the expected, but not the same.

	//As integers, floating-point numbers have a certain size: 32 bits or 64 bits. Using a larger size increases the accuracy – how many digits we can use for the calculation.

	//	In addition to numbers, there are several other values, such as NaN, for things like 0/0, as well as positive and negative infinity (+∞ and −∞).

	//There are two real types in Go: float32 and float64. They are often called single- and double-precision real numbers, respectively. When working with real numbers, it is enough to use float32, however, if you want to work with more accurate numbers, you can also use float64.

	//Booleans

	//The Boolean type – bool (named after George Boole) – is a special integer type used to represent truth and false. A variable of this type will occupy only one byte. Three logical operators are used with this type:
	//&& 	AND
	//|| 	OR
	//! 	NOT
	var b bool = true

	//Example of operators:
	var b1 = true
	b2 := !b1      // false
	b3 := b2 || b1 // false OR true => true
	b4 := b2 && b3 // false AND true => false
	b5 := b1 && b3 // true AND true => true

	//Rune and byte

	//There are also some interesting numeric data type names in Go, let's have a look at them:

	//rune is an alias for int32
	//byte is an alias for uint8

	//Golang doesn't have a separate data type for characters. It uses byte and rune to represent character values. The byte data type can represent ASCII characters, and the rune data type can store Unicode characters in UTF-8 format.

	//Complex numbers

	//Unlike some other languages, Go supports complex numbers: complex128and complex64. By now you should have guessed it – complex64 contains a 32-bit real and a 32-bit imaginary part, and complex128 has 64-bit real and imaginary numbers. Go also provides a built-in function named complex for creating complex numbers. You'll need to use the complex function if you're creating a complex number with variables. Complex numbers allow all the basic arithmetic operations: + , – , * , and /.

	//We will not dwell on this, but if you are still interested in it, you can refer to the documentation of the cmplx package.
	//	https://pkg.go.dev/math/cmplx
}

func solve() uint8 {
	// declare an uint8 type variable 'a' with a value of 10
	// edit this line
	var a uint8
	a = 10

	return a
}
