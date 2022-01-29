// ZY, 29 Jan 2022, 11.57am
// Learning Go through JetBrains Academy free lessons
//Hexadecimal Numbers

//Anyone who spends enough time with a computer or other digital device is likely to come across mysterious records consisting of numbers and Latin letters, such as 10FEF10FEF10FEF. For uninitiated people, they may seem like some kind of cipher. What is behind these symbols? It turns out that these are simply numbers in HEX (Hexadecimal Number System). HEX is an alternative way to represent numbers.
//
//Note: HEX can also be written in lowercase (hex, an example of usage: Wikipedia article), but in this topic, we, for consistency, will use only uppercase HEX or "hexadecimal".
//Why HEX-a-decimal?
//
//For years before the invention of computers, people used the decimal (base-101010) system for counting – because it's convenient when you have 101010 fingers and 101010 toes.
//
//In computers, however, there are only 222 options: on and off, which created the idea of a binary digit or bit, for short.
//
//For convenience, computer engineers tended to group bits. In the 1960s, they could group 333 bits at a time: a group of 333 bits is base 2∗2∗2=2*2*2=2∗2∗2= base-888 number.
//
//As computers got more powerful, people learned to cluster bits by 444 instead of 333. A group of 444 bits, written as one symbol, can have 161616 values instead of 888 – in other words, two times more than a group of 333 bits. So, this system is more compact than decimal, octal (base-888), and binary systems!
//
//The only thing missing – a name for the new awesome number system.
//
//A single number can have 161616 values. 16=10+616=10+616=10+6
//
//Hmmmm... Hex (from Greek) means 666, and decimal (from Latin, but that's OK) means 101010 (also, we are already used to the decimal number system and the word decimal). So...
//
//HEX-a-decimal looks like the perfect name!
//
//Hexadecimal digits
//
//The hexadecimal number system uses numbers 0,1…90,1\dots90,1…9 and first six Latin letters: A,B,C,D,E,FA, B, C, D, E, FA,B,C,D,E,F (corresponding to decimals 10,11…1510,11\dots1510,11…15), as hexadecimal digits.
//
//Yes, in some numeral systems, digits can be letters, too!
//
//The table below compares how numbers 1,2…151,2\dots151,2…15 are written in the binary, the decimal, and the HEX number systems:
//Binary 	Decimal 	HEX
//000020000_200002​ 	0100_{10}010​ 	0160_{16}016​
//000120001_200012​ 	1101_{10}110​ 	1161_{16}116​
//001020010_200102​ 	2102_{10}210​ 	2162_{16}216​
//001120011_200112​ 	3103_{10}310​ 	3163_{16}316​
//010020100_201002​ 	4104_{10}410​ 	4164_{16}416​
//010120101_201012​ 	5105_{10}510​ 	5165_{16}516​
//011020110_201102​ 	6106_{10}610​ 	6166_{16}616​
//011120111_201112​ 	7107_{10}710​ 	7167_{16}716​
//100021000_210002​ 	8108_{10}810​ 	8168_{16}816​
//100121001_210012​ 	9109_{10}910​ 	9169_{16}916​
//101021010_210102​ 	101010_{10}1010​ 	A16A_{16}A16​
//101121011_210112​ 	111011_{10}1110​ 	B16B_{16}B16​
//110021100_211002​ 	121012_{10}1210​ 	C16C_{16}C16​
//110121101_211012​ 	131013_{10}1310​ 	D16D_{16}D16​
//111021110_211102​ 	141014_{10}1410​ 	E16E_{16}E16​
//111121111_211112​ 	151015_{10}1510​ 	F16F_{16}F16​
//
//So, as you see, in HEX, when we run out of digits, we start using letters. How is the decimal number 161616 represented in HEX? The answer is 101010!
//
//It's better to write numbers with their base as a lower index, so you don't confuse decimal 101010 with a hexadecimal 101010: 101610_{16}1016​
//
//If you are quick-thinking, you have already noticed those lower indices in the table above.
//
//Other ways to avoid confusion are writing a HEX number with h after it or 0x before it. For instance, 63h, 0x63 and 631663_{16}6316​ are three different ways to indicate that 636363 is a hexadecimal number.
//
//Binary to HEX and vice versa
//
//HEX is widely used in computer science because we can easily translate binary numbers into HEX. Moreover, HEX is (and has been created) as a simplified way to represent binary numbers.
//
//Remember: 161616 is 242^424. That means we can split (starting from the right, or, in other words, from the end) any binary number into a sequence of 444-digit binary numbers. These numbers are then mapped to HEX digits using the table above.
//
//Let's try converting the number 111001100111000121110011001110001_211100110011100012​ to HEX:
//
//11100110011100012=1110 0110 0111 00012=E671161110011001110001_2 = 1110\ 0110\ 0111\ 0001_2 = E671_{16}
//11100110011100012​=1110 0110 0111 00012​=E67116​
//
//Of course, everything is not as simple as it seems. If the length of our binary number is not divisible by 444 without remainder, then the leftmost number of our sequence will contain less than 444 digits:
//
//1011002=10 11002= ???16101100_2 = 10\ 1100_2 =\ ???_{16}
//1011002​=10 11002​= ???16​
//
//And again, if you are quick-thinking, you have already noticed and thought about such cases. Maybe, you've even found a solution!
//
//Did you notice that, in a decimal system, adding any numbers of zeroes to the left of a decimal number (leading zeroes) doesn't change the number's value? 42=042=000042=000…0004242=042=000042=000\dots0004242=042=000042=000…00042, and so on. Well, actually, not only decimal number system works like that:
//
//In any base-n (n > 1) number system adding any number of leading zeroes to the number does not change its value.
//
//Those zeroes are even called non-significant zeroes.
//
//How can it help us? Simply: remember that 102=0010210_2=0010_2102​=00102​, and, therefore:
//
//1011002=10 11002=2C16101100_2 = 10\ 1100_2 =2C_{16}
//1011002​=10 11002​=2C16​
//
//So, if the leftmost number in our split sequence of 444-digit binary numbers contains less than 444 digits, we add leading zeroes to it until it is a 444-digit binary number and then match it, along with the rest in the sequence, to the corresponding HEX digit using the table. The resulting sequence is ready to be mapped to HEX digits using the table above.
//
//Converting a hexadecimal number into a binary is even easier. Each HEX digit corresponds to the 444-digit binary sequence. All you have to do is map digits to sequences with the table.
//
//BE16=1011 11102=101111102BE_{16} = 1011\ 1110_2 = 10111110_2
//BE16​=1011 11102​=101111102​
//Applications of HEX
//
//We use the hexadecimal system to record error codes during the work of various software products. For example, operating system errors are encoded in this way. If you decode your error code, you'll find out the exact error that occurred. Besides, in URLs, character codes are written as hexadecimal pairs prefixed with %. You can see for yourself by googling a symbol @. The link in the address bar of your browser would look like that:
//
//https://www.google.com/search?q=%40
//
//The HEX numbers are also used for writing programs in low-level languages and in some encodings. For example, in Unicode (a computer standard for symbols encoding), every symbol is represented as a hexadecimal number. Even color schemes are encoded by HEX numbers. Thus, in RGB encoding, every color can be represented as 333 hexadecimal numbers, each for Red, Green, or Blue color components respectively. As you can see, HEX is used in many ways, so there is no doubt that this information will be of use to you in the future!
