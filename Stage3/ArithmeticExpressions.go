// ZY, 30 Jan 2022, 11.31am
// Learning Go through JetBrains Academy free lessons
//Theory: Arithmetic expressions in programming

//We are familiar with arithmetics since childhood and it seems as simple as ABC. Let's say you have 30 coins, and you want to treat five of your friends to some ice cream that costs 2 coins. 2∗5=102 * 5 = 102∗5=10, so it's definitely enough. In case you also want to wash it down with a Coke at 3 coins per cup, you will get (2+3)∗5=25(2 + 3) * 5 = 25(2+3)∗5=25, and this is enough too.
//
//
//You notice the last equation differs from the previous one because we added parentheses. Without them 2+3∗5=2+15=172 + 3 * 5 = 2 + 15 = 172+3∗5=2+15=17 due to the precedence of operators. So, the rules exist even for basic arithmetic since some expressions can be quite complex. We've learned these rules a long time ago. Now let's find out how we can explain them to a computer.
//Operands and operators, precedence of operators
//
//When we were figuring out if we could throw a party and treat all our friends, we were using operators: the addition operator (+++) and the multiplication operator (⋅\cdot⋅). The division operator (///) and the subtraction operator (−-−) will help us find out if we can invite someone else: (30–25)/5=1(30–25)/5=1(30–25)/5=1, so another friend can definitely join in.
//
//We can also perform operations with letters. Imagine all your friends want different ice creams at different prices. Here we will try to stay within 30 coins budget and not offend anyone. We can write down our predicament in mathematical language:
//
//A+B+C+D+E=30A + B + C + D + E = 30A+B+C+D+E=30
//
//The letters A,B,C,DA, B, C, DA,B,C,D, and EEE are variables that indicate the price of a chosen ice cream. They have different values depending on your friends' choices.
//
//Variables and numbers are called operands in arithmetic expressions.
//
//Each operator has its own precedence or order of use. Much like a driver gives way to a pedestrian at a crosswalk, the addition and subtraction operators give way to multiplication and division. It is very important to use higher-precedence operators before lower-precedence ones. For example, 80+100/2=13080+100/2=13080+100/2=130, not 909090. If your salary was in the question, you would probably be angry with the wrong prioritization, wouldn't you?
//
//For sure, you have already encountered these operators, and below is a table of operation priorities.
//
//Only the parentheses can change the order of precedence. We use them to distinguish special calculations and give the correct meaning to our expressions. In the expression (2+3)⋅5=25(2 + 3)\cdot5 = 25(2+3)⋅5=25, we first count how much a treat of ice cream and Coke will cost a person.
//Forms of arithmetic operations' notations
//
//We already know a lot about operations and their notations. Now, look at the following expressions. Do you think all of them make sense?
//
//    (A+B)⋅C(A + B) \cdot C(A+B)⋅C
//    C/B−7C / B - 7C/B−7
//    AB+C⋅AB+C\cdotAB+C⋅
//    +B/CD+B/CD+B/CD
//
//You have probably spotted that there is no operator between the operands A and B in the third expression, and the operator +++ stands before the operand BBB in the fourth one. It may look unusual to you, as you are probably used to certain rules of notation. Usually, each operator is placed between two operands and is read from left to right. Such notation is called an infix notation. With its help, we know when and with which operands this operator works.
//
//In reality, there are other ways to express and perform arithmetic operations. We will consider two of them: postfix and prefix notations. Have you already guessed what is the difference between them? Of course, the position of operators between operands. In the prefix or Polish notation, the operator precedes the operands. In the postfix (reverse Polish notation or RPN) it follows the operands. Have a look at examples in the table below:
//
//Sometimes programmers may decide on either prefix or Polish notation. They do it because in these notations the order of the operations is fully determined by the position of the operators and nothing else, so there are no parentheses. It means the computer program does not need to search for special operations and try to identify their priority to start the calculation.
//
//So, here we briefly discussed that there can be three different operation's notations forms. We will learn about this in more detail in later topics.
//Conclusion
//
//To sum up,
//
//    arithmetic expressions are composed of operators and operands,
//    operators are arithmetic operations, like +, -, *, and so on,
//    operands are objects of arithmetic operations, like variables and numbers,
//    each operator has its own precedence or order of use,
//    there can be infix, postfix, and prefix operator notations.
