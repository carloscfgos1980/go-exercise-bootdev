# CONSTANTS

- Constant concept
Constants are declared with the const keyword. They can't use the := short declaration syntax.
const pi = 3.14159
Constants can be primitive types like strings, integers, booleans and floats. They cannot be more complex types like slices, maps and structs, which are types we will explain later.
As the name implies, the value of a constant can't be changed after it has been declared.

- Computed Constants
Constants must be known at compile time. They are usually declared with a static value:
const myInt = 15

However, constants can be computed as long as the computation can happen at compile time.
For example, this is valid:
const firstName = "Lane"
const lastName = "Wagner"
const fullName = firstName + " " + lastName

That said, you cannot declare a constant that can only be computed at run-time like you can in JavaScript. This breaks:
// the current time can only be known when the program is running
const currentTime = time.Now()

- Formatting Strings in Go
Go follows the printf tradition from the C language. In my opinion, string formatting/interpolation in Go is less elegant than Python's f-strings, unfortunately.
    • fmt.Printf() - Prints a formatted string to standard out.
    • fmt.Sprintf() - Returns the formatted string

These following "formatting verbs" work with the formatting functions above:
Default Representation
The %v variant prints any value in a default format. It can be used as a catchall.
s := fmt.Sprintf("I am %v years old", 10)
// I am 10 years old

s := fmt.Sprintf("I am %v years old", "way too many")
// I am way too many years old
If you want to print in a more specific way, you can use the following formatting verbs:
String
s := fmt.Sprintf("I am %s years old", "way too many")
// I am way too many years old
Integer
s := fmt.Sprintf("I am %d years old", 10)
// I am 10 years old
Float
s := fmt.Sprintf("I am %f years old", 10.523)
// I am 10.523000 years old

// The ".2" rounds the number to 2 decimal places
s := fmt.Sprintf("I am %.2f years old", 10.523)
// I am 10.52 years old

- Runes and String Encoding
In many programming languages (cough, C, cough), a "character" is a single byte. Using ASCIIencoding, the standard for the C programming language, we can represent 128 characters with 7 bits. This is enough for the English alphabet, numbers, and some special characters.

In Go, strings are just sequences of bytes: they can hold arbitrary data. However, Go also has a special type, rune, which is an alias for int32. This means that a rune is a 32-bit integer, which is large enough to hold any Unicode code point.
When you're working with strings, you need to be aware of the encoding (bytes -> representation). Go uses UTF-8 encoding, which is a variable-length encoding.
UTF-8 Text
Type text to see its code points and bytes. Joining adjacent emoji inserts a zero width joiner (U+200D). 

What Does This Mean?
There are 2 main takeaways:
    1. When you need to work with individual characters in a string, you should use the runetype. It breaks strings up into their individual characters, which can be more than one byte long.
    2. We can include a wide variety of Unicode characters in our strings, such as emojis and Chinese characters, and Go will handle them just fine.

FORMAT PRACTICE:
package main

import "fmt"

func main() {
	fname := "Dalinar"
	lname := "Kholin"
	age := 45
	messageRate := 0.5
	isSubscribed := false
	message := "Sometimes a hypocrite is nothing more than a man in the process of changing."

	// Don't touch above this line

	userLog := fmt.Sprintf("Name: %s %s, Age: %d, Rate: %.1f, Is Subscribed: %t, Message: %s", fname, lname, age,messageRate, isSubscribed, message)

	fmt.Println(userLog)
}
