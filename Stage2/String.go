// ZY, 16 Jan 2022, 8.37pm
// Learning Go through JetBrains Academy free lessons
// String theory

//Strings and string declaration
// The default value for string is "" or `` - empty string
var myFirstString string

// A string with special characters
var iAmSpecial string = "Hello\n\t"

/*
   Escape sequence Value
   \n A newline character
   \t A tab character
   \" Double quotation marks
   \\ A backslash
*/

fmt.Println("ABOBA\n\t") // This string consits of the word ABOBA, a new line and tabulation
/* Output:
   ABOBA


*/

fmt.Println(`ABOBA\n\t`) // This string consits of the word ABOBA\n\t
/* Output:
ABOBA\n\t
*/

//String contents are immutable in Golang, so when we concatenate two strings, it creates a new one in memory:

discover := "hello"               // the variable contains the "hello" string

discover = discover + " there"    // the variable contains a new "hello there" string;
// the "hello" string soon will be removed from memory

discover = "world"                // the variable again has a new value;
// the "hello there" string soon will be removed from memory

//UTF-8

//Remember we mentioned the scary word UTF-8? Shortly, it means that you can have a string with symbols from almost any alphabet.

//Side note: in standard UTF-8 Unicode, character representations occupy from 1 to 4 bytes.
// UTF-8 from box
var russian = "Привет, Мир!"

korean := "안녕하세요 월드입니다!"

var emoji string = "🙋🌍❗"

//As we've mentioned earlier, each string is a sequence of bytes. Hence, if you need to find out its byte length, you can use the function len(yourStringName).
asciiString := "ABCDE"
utf8String  := "БГДЖИ"

fmt.Println(len(asciiString)) // 5
fmt.Println(len(utf8String))  // 10

//Runes

//Go uses rune values to represent Unicode characters – it's all about UTF-8. Besides, you can assume that strings are not only sequences of bytes, but sequences of runes. Remember the example above? If you are interested in the length of a string in characters, use the function RuneCountInString from unicode/utf8 package:

utf8.RuneCountInString(asciiString) // 5
utf8.RuneCountInString(utf8String)  // 5

// Emoji example 🗿
len(emoji)                          // 11
utf8.RuneCountInString(emoji)       // 3

//Final remark: for safer work with strings, you should convert them into []rune , but we will cover it in detail in another topic.