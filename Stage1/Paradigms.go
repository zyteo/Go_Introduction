// ZY, 25 Jan 2022, 8.42pm
// Learning Go through JetBrains Academy free lessons
//Paradigms

//Different programming tasks can be solved in different ways: you might need to write a function, or create a separate class with methods, and so on. All such variants are combined into different programming approaches, which are also called paradigms. Below we will analyze two main paradigms: imperative and declarative and their types.
//
//Almost all modern languages are multi-paradigm. They easily combine the capabilities of the imperative and declarative approaches. Nevertheless, in the section about each paradigm, we will give a list of languages where, among other paradigms, the one discussed in the section is also implemented.
//
//Imperative paradigm
//
//The imperative paradigm is one of the oldest programming paradigms. It is closely related to machine architecture. The imperative program is similar to the orders expressed by the imperative mood in natural languages. It is a sequence of instructions that the processor must execute step by step. The main focus of this paradigm is on how to achieve the goal. The paradigm consists of several statements, and after executing them all, the result is stored.
//
//For example, you want to display the phrase "Hello, <username>!" on the screen. How should the program do it? Through the following steps:
//
//    ask the username
//    read and remember the username
//    display the result
//
//Imperative programming is divided into three broad categories: Procedural programming paradigm, Object-oriented programming, and Parallel processing approach. Let's talk about them in more detail.
//Procedural programming paradigm
//
//The procedural programming paradigm is based upon the concept of procedure calls, in which statements are structured into procedures also known as subroutines or functions. They are a list of instructions to tell the computer what to do step by step. In other words, the computer takes input data and changes it sequentially, remembering each new change. There is no difference between procedural and imperative approaches.
//
//Procedural Programming is suitable for general-purpose programming in order to complete common tasks. So, this can be a small computational problem, such as calculating a factorial, or finding the area of a figure, or displaying some information/phrase like "Hello, world!". Also, the code can be reused in different parts of the program, without the need to copy it.
//
//The algorithm written in this paradigm is very simple to implement, but it is rather slow and cannot solve a complex problem.
//
//Programming languages that have implemented the Procedural Programming paradigm are C, Java, C++, ColdFusion, Pascal.
//Object-oriented programming
//
//Object-oriented programming or OOP is the paradigm where the program is written as a collection of classes. Each class has its instances called objects.
//
//A class is a way of describing an entity in general, defining the usual state and behavior that depends on that state, as well as the usual rules for interacting with this entity. Formally, a class is viewed as a set of data like fields, attributes, class members, and functions, i.e. methods for working with them.
//
//For example, we have an entity cat and we want to describe it using a class. So, the cat will be an object of the corresponding Cat class. A cat has some attributes, for example, a tail, paws, claws, muzzle, ears, whiskers. A cat's behavior is what it usually does, for example, it can run, jump, meow, eat, and rip off the wallpaper. All of these will be cat methods.
//
//OOP can handle almost all kinds of common real-life problems where you need to model typical objects and work with them.
//
//Programming languages that have implemented the OO paradigm are Ruby, Java, C++, Python, Simula (the first OOP language), Smalltalk, Visual Basic .NET, Objective-C.
//Parallel processing approach
//
//Parallel processing helps reduce instruction execution time. It does this by sharing or parallelizing instructions across multiple processors. The meaning of the approach can be summarized in one phrase: "divide and conquer". Examples are NESL (one of the oldest) and C / C ++ (also supported due to some library functions).
//Declarative paradigm
//
//Declarative programming is a programming paradigm in which it is important to specify the problem and the expected result of its solution. That is, in contrast to the imperative paradigm, where it is necessary to answer the question "how to do this?" you need to ask the questions "What needs to be done?" and "What will be the result of the work?". So, rather than providing step-by-step instruction, you tell the system what you need and let it try to come up with a solution.
//
//Declarative programming is divided into Logic, Functional, and Database paradigm types. We will describe them below.
//Logic programming paradigm
//
//Logic programming is a programming paradigm that is heavily based on formal logic. Any program written in a logical programming language is a set of sentences in a logical form that express facts and rules about a certain problem area.
//
//So, the basic statements of logic programming are as follows:
//
//    Facts are fundamental assertions about the problem domain, like "Socrates is a man."
//    Rules are inferences about the facts in the domain ("All men are mortal.")
//    Queries are questions about that domain ("Is Socrates mortal?")
//
//In general, the task here is to find the answer to the query based on facts and rules.
//
//The major families of logic programming languages include Prolog, Answer Set Programming (ASP), and Datalog.
//Functional programming paradigm
//
//Functional programming is a programming paradigm, in which the computation process is interpreted as the computation of the values of functions. The function, in this case, is similar to a mathematical one. That is, a function in which input is an array that is not changed, and the output is a new array with new data. This makes a mathematical function different from a function in procedural programming, where a function is a sequence of actions that change the original data.
//
//Here's a simple example: you might have a function that takes a list of numbers as input and returns a new list with the squares of those numbers. This does not change the original list of numbers.
//
//Programming languages that have implemented the Functional programming paradigm are JavaScript, Haskell, Scala, Erlang, Lisp, ML, Clojure, OCaml, Common Lisp, and F#.
//Database programming paradigm
//
//This programming methodology is based on working with data. The data is stored in the database and queries are made to this database in a special language, for example, SQL. With these languages, you can access the data for filtering, transformations, calculating statistics, and so on. Program statements are data-defined rather than a hard-coded series of steps.
//
//The database program is the heart of the business information system, allowing for file creation, data entry, updating, querying, and reporting functions.
//Conclusion
//
//To sum up,
//
//    Different approaches to creating programs are called paradigms.
//    There are two main programming paradigms: imperative and declarative.
//    The imperative paradigm focuses on achieving a result using step-by-step instructions that change the data sequentially.
//    Imperative paradigm includes Procedural programming paradigm, Object-oriented programming, and Parallel processing approach.
//    The declarative paradigm focuses on the task and tries to get an expected result.
//    Declarative paradigm includes Logic, Functional, and Database paradigms.