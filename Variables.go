package main

// ZY, 15 Jan 2022, 8.49pm
// Learning Go through JetBrains Academy free lessons

//declaring a variable
//var variableName type // The default value

var firstNum int = 1
var secondNum = 20 // Skipping the type
// You can also use := in statements. It is a declare and assign construction:
thirdNum := 30
// Please remember that you can assign only a new variable with this operator. It means that you cannot assign a value to the already existing one. In this case, it will throw an error:
fourthNum := 30
fourthNum = 31 // Ok
//fourthNum := 32 // Panic! This variable is already assigned
fourthNum, fifthNum := 32, 33 // no panic
