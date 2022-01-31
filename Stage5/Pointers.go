// ZY, 31 Jan 2022, 8.55pm
// Learning Go through JetBrains Academy free lessons
//Pointers

//In this topic, we will discuss pointers — an important subject that can be somewhat difficult for beginners, though. We will try to make it as clear as possible.
//What is a pointer?
//
//As you may know, computer memory is made of blocks. Each block has an address and a state. The state represents a value of a variable, whereas the address identifies a block. Thus, every block has its own unique address, represented by a number. In modern computers, we use a 64-bit integer to represent addresses. As you can see, an address itself is nothing more than just an integer value. A variable holding a value of a memory address is called a pointer.
//
//In high-level languages like Go, we have special control flow structures like if and for statements, as well as functions for controlling execution. However, we still use pointers, in a way that we can grant permission to a certain function, for it to change an existing variable. Another case when pointers come in handy is when we have a big variable and want to avoid copying it, in order to reduce memory consumption.
//Pointer type and a new function
//
//For safety purposes, we work with pointers leading to a certain type, rather than with raw ones.
//
//var addressOfStringVariable *string
//
//In the example above, we use an asterisk before the name of the type, to declare a variable of the type pointer to a string, or in other words, we create a pointer to a string.
//
//To clarify the difference:
//
//    A variable of the type string is holding a string value
//    A variable of the type *string is holding an address of a variable of the type string
//
//Thus, when you declare a pointer, it doesn't share the same default value with the underlying type, but instead has its own default value – nil.
//
//var actualStringVariable string       // Is equal to ""
//var addressOfStringVariable *string   // Is equal to nil
//
//Now, when we just declare a pointer, it doesn't point anywhere. To create a pointer of some type and actually allocate memory for its default value, we should use a special built-in function new. Let's have a look at how we can use this function:
//
//var p *string   // Declaring p as a pointer to a string
//
//p = new(string) // Allocating the memory for the default string value
//                // and assigning its address to the p pointer
//
//This new function returns a pointer to a type, specified as the function argument. We can also merge these two lines into one:
//
//var p = new(string)
//
//Pointer dereference and value assignment
//
//Now, we have a pointer that points to the default value of a given type. Let's check if it's true. The process of getting the actual value from the pointer is called pointer dereference.
//
//var p = new(string)
//fmt.Print(p)  // Will print some memory address, for eg. 0xc000040240
//              // This address can be different in your case
//
//fmt.Print(*p) // Will print the actual value — an empty string
//
//We use an asterisk again, but this time with the name of the variable, instead. Through this, we dereference the pointer and obtain the actual value. But be careful, dereferencing a nil pointer will lead to a runtime panic.
//
//var p *string
//fmt.Print(*p) // invalid memory address or nil pointer dereference
//
//The nil pointer dereference is one of the most common mistakes in Go, so don't forget to check that a pointer is not nil before accessing its value.
//
//It would be boring if we could only have a pointer pointing at the default value. Assigning a value to the variable behind the pointer is not difficult. Follow me in this example:
//
//var p = new(string)
//*p = "my string"
//
//Remember, the p variable is nothing more than just a fancy integer representing the memory address. We cannot assign a string to it directly. By putting an asterisk before the variable name, we tell Go that we want to work not with the p variable, but with the variable hidden behind it. In this case, you cannot put an asterisk anywhere you want, only before a valid pointer.
//Getting the pointer
//
//One last thing you need to know about pointers is how to get the address of a variable. Look, it is quite simple:
//
//var s = "some string variable" // var s of the type string
//var p = &s                     // var p is a pointer to a string
//
//fmt.Println(s) // Printing the value
//fmt.Println(p) // Printing the address where the value is stored
//
//Using the & symbol before the variable name, we tell Go that we work with its address, not its value.
//
//The Go language doesn't allow to move data from one place in memory to another directly, so we can't assign a new address to the variable by calling &varName = address.
//
//These below are proper examples of getting and assigning pointers.
//
//var p *string
//var s = "my string"
//
//p = &s
//
//var p = new(string)
//var s = "my string"
//
//*p = s
//
//It is important to have the new function in the second example. Otherwise, we would have a nil pointer dereference error at *p = s.
//Advanced examples
//
//You will not see questions about the following examples in the practice section. Both of them are about having a pointer that points to another pointer that points to some "solid" type. Sounds complicated, and you'll rarely see such uses in practical application, but you can consider these code snippets and play with them for a bit. It'll definitely help you with the understanding of pointers.
//
//var p **string
//p = new(*string)
//*p = new(string)
//fmt.Print(**p)
//
//**p = "is this even possible?"
//fmt.Print(**p)
//
//This is another way:
//
//var s = "yes, it is possible"
//var p1 = &s
//var p2 = &p1
//
//fmt.Println(*p2 == p1) // Will print true
//fmt.Println(**p2 == s) // Will print true
//
//Conclusion
//
//Congrats for making it here! It's been a big topic where we've learned a lot about pointers in Go, what they are and how to use them. In particular, we've covered the following:
//
//    A pointer is a variable holding the memory address;
//    The type of a pointer is the type it points at;
//    For proper pointer creation, we use the new function;
//    For reference and dereference, you should use the & and * symbols;
//    Finally, for value assignment use *pointer = value.