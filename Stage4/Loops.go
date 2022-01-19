package main

// ZY, 19 Jan 2022, 8.42am
// Learning Go through JetBrains Academy free lessons
//Loops theory

func main() {

	//To create a loop in Go, we use the for keyword, and then we specify a condition. After this, we put a block of code inside curly brackets. We call it the body of the loop. One run of this block of code we call an iteration. Before every iteration, Go will check the condition. If it's true, execution goes inside the body of the loop, and if it's false, it leaves the loop. The structure is exemplified as follows:
	var condition bool

	for condition {
		// here is the body of the loop
	}

	//	We can use the true constant as the condition to create an infinite loop:
	//
	//for true {
	//    // here is the body of an infinite loop
	//}

	//	 Using the for keyword without any condition will do the same.
	//
	//for {
	//    // here is the body of an infinite loop
	//}

	//We can use as a condition anything that returns bool. For example, the result of the comparison operator:
	//
	//var a, b int
	//
	//for a < b {
	//    // here is the body of the loop
	//}

	//Let's consider a task: we want to run a code 10 times. What do we need? A counter, a way to increment it after each iteration, and a condition. Here is how we can do it:
	//
	//var i int           // initializing the i variable as a counter
	//
	//for i < 10 {        // initializing a loop with the condition i < 10
	//    // some code
	//    i++             // incrementing the counter at the end of the body
	//}

	//	 We can also use an extended loop declaration:
	//
	//for i := 0; i < 10; i++ {
	//    // some code
	//}

	//	It saves a few lines, but there is one difference: in the former code example above, the i variable is visible outside the loop, whereas in the latter example it is not visible outside, only inside. Keep in mind that within the scope of the loop we'll have access only to the inner variable i if we declare two of them in inner and outer scopes:

	//	var i = 10    // Declaring the i variable in the outer scope
	//
	//for i := 0; i < 1; i++ {  // Declaring the i variable in the inner(loop) scope
	//    fmt.Print(i)          // Will print 0, not 10
	//}

	//	Flow of control
	//
	//In some cases, we need to exit the loop or jump to the next iteration on command. For these purposes, there are two keywords: break and continue. Consider the following example:
	//
	//var s int             // step 1
	//var i int             // step 2
	//
	//for ; i < 50; i++ {   // step 3, 4, 8; you can have a sparse loop declaration
	//    if s > 100 {      // step 5
	//        break         // step 5
	//    }
	//    if i % 2 == 0 {   // step 6
	//        continue      // step 6
	//    }
	//    s += i            // step 7
	//}                     // step 9
	//                      // step 10
	//
	//It represents this algorithm:
	//
	//    Initialize the s variable as a sum;
	//    Initialize the i variable as a current number;
	//    Loop begins here;
	//    If i is NOT less than 50, go to step 10 (if the condition is false, we leave the loop);
	//    If s is greater than 100, go to step 10 (the body of the loop);
	//    If the remainder of the division of i by 2 is equal to 0, go to step 8 (the body of the loop);
	//    Add i to s (the body of the loop);
	//    Increment i by 1 (i++; we don't consider this as part of the body of the loop, since we have defined it inside the for declaration);
	//    Go to step 3;
	//    Loop ends here.
	//
	//This algorithm sums all odd numbers between 0 and 49, and it stops if this sum is greater than 100.

	//	As you can see, executing the break command will terminate the loop. Executing the continue command will terminate only the current iteration. After this, Go will increment the counter and begin the next iteration if the condition is still true.

}
