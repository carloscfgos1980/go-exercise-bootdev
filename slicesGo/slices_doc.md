# Slices

## Arrays in Go
Arrays are fixed-size groups of variables of the same type. For example, [4]string is an array of 4 values of type string.
To declare an array of 10 integers:
var myInts [10]int
or to declare an initialized literal:
primes := [6]int{2, 3, 5, 7, 11, 13}

## Slices in Go
99 times out of 100 you will use a slice instead of an array when working with ordered lists.
Arrays are fixed in size. Once you make an array like [10]int you can't add an 11th element.
A slice is a dynamically-sized, flexible view of the elements of an array.
The zero value of slice is nil.
Non-nil slices always have an underlying array, though it isn't always specified explicitly. To explicitly create a slice on top of an array we can do:
primes := [6]int{2, 3, 5, 7, 11, 13}
mySlice := primes[1:4]
// mySlice = {3, 5, 7}
The syntax is:
arrayname[lowIndex:highIndex]
arrayname[lowIndex:]
arrayname[:highIndex]
arrayname[:]
Where lowIndex is inclusive and highIndex is exclusive.
lowIndex, highIndex, or both can be omitted to use the entire array on that side of the colon.
Sequence Slicing
Adjust the slice controls and watch values move into the result.

## Slices Review
Slices wrap arrays to give a more general, powerful, and convenient interface to sequences of data. Except for items with explicit dimensions such as transformation matrices, most array programming in Go is done with slices rather than simple arrays.
Slices hold references to an underlying array, and if you assign one slice to another, both refer to the same array. If a function takes a slice argument, any changes it makes to the elements of the slice will be visible to the caller, analogous to passing a pointer (we'll cover pointers later) to the underlying array. A Read function can therefore accept a slice argument rather than a pointer and a count; the length within the slice sets an upper limit of how much data to read. Here is the signature of the Read() method of the File type in package os:
Questions
1. Which references the other?
Answer: Slices reference arrays
In Go, an array is a fixed-size block of contiguous memory that holds the actual data.
A slice, under the hood, is a small data structure (called a slice header) containing three things:
    1. A pointer to an underlying array
    2. A length (number of elements currently in the slice)
    3. A capacity (maximum number of elements the slice can grow to before reallocating)
Because the slice holds a pointer pointing to the array in memory, the slice references the array, not the other way around.

2. Multiple slices can point to the same array
Answer: True
Since slices only hold a reference (pointer) to an underlying array, you can create multiple slices that look at different (or overlapping) parts of the same underlying array.
numbers := [5]int{10, 20, 30, 40, 50}

// Both sliceA and sliceB reference parts of the 'numbers' array
sliceA := numbers[0:3] // [10, 20, 30]
sliceB := numbers[2:5] // [30, 40, 50]
Both sliceA and sliceB share the same underlying memory for the element at index 2 (30).

3. A function that only has access to a slice can modify the elements of the underlying array
Answer: True
When you pass a slice into a function, Go passes a copy of the slice header (the pointer, length, and capacity). However, the copied pointer still points to the same underlying array.
If the function changes an element via the slice index (e.g., s[0] = 99), it directly modifies the memory in that underlying array. Those changes will be visible to any caller referencing that same array.

## Make
Most of the time we don't need to think about the underlying array of a slice. We can create a new slice using the make function:
// func make([]T, len, cap) []T
mySlice := make([]int, 5, 10)

// the capacity argument is usually omitted and defaults to the length
mySlice := make([]int, 5)
Slices created with make will be filled with the zero value of the type.
If we want to create a slice with a specific set of values, we can use a slice literal:
mySlice := []string{"I", "love", "go"}
Notice the square brackets do not have a 3 in them. If they did, you'd have an array instead of a slice.
Length
The length of a slice is simply the number of elements it contains. It is accessed using the built-in len() function:
mySlice := []string{"I", "love", "go"}
fmt.Println(len(mySlice)) // 3
Capacity
The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice. It is accessed using the built-in cap() function:
mySlice := []string{"I", "love", "go"}
fmt.Println(cap(mySlice)) // 3
Generally speaking, unless you're hyper-optimizing the memory usage of your program, you don't need to worry about the capacity of a slice because it will automatically grow as needed.
Indexing
A programming concept you should already be familiar with is array indexing. You can access or assign a single element of an array or slice by using its index.
mySlice := []string{"I", "love", "go"}
fmt.Println(mySlice[2]) // go

mySlice[0] = "you"
fmt.Println(mySlice) // [you love go]

- Len and Cap Review
The length of a slice may be changed as long as it still fits within the limits of the underlying array; just assign it to a slice of itself. The capacity of a slice, accessible by the built-in function cap, reports the maximum length the slice may assume. Here is a function to append data to a slice. If the data exceeds the capacity, the slice is reallocated. The resulting slice is returned. The function uses the fact that len and cap are legal when applied to the nil slice, and return 0.
Referenced from Effective Go
func Append(slice, data []byte) []byte {
    l := len(slice)
    if l + len(data) > cap(slice) {  // reallocate
        // Allocate double what's needed, for future growth.
        newSlice := make([]byte, (l+len(data))*2)
        // The copy function is predeclared and works for any slice type.
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[0:l+len(data)]
    copy(slice[l:], data)
    return slice
}

Questions:
1. What does the cap() function return?
The maximum length of the slice before reallocation of the array is necessary
2. What does the len() function return?
The current length of the slice
3. What do len() and cap() do when a slice is nil?
Return 0

## Variadic
A variadic function takes an arbitrary number of final arguments. These are called variadic arguments and are received as a slice using the ... syntax in the function signature.
func concat(strs ...string) string {
    final := ""
    // strs is just a slice of strings
    for i := 0; i < len(strs); i++ {
        final += strs[i]
    }
    return final
}

func main() {
    final := concat("Hello ", "there ", "friend!")
    fmt.Println(final)
    // Output: Hello there friend!
}
The familiar fmt.Println() and fmt.Sprintf() are variadic, as are many in the standard library! fmt.Println() prints each element with space delimiters and a newline at the end.
func Println(a ...interface{}) (n int, err error)
Spread Operator
The spread operator allows us to pass a slice into a variadic function. The spread operator consists of three dots following the slice in the function call.
func printStrings(strings ...string) {
        for i := 0; i < len(strings); i++ {
                fmt.Println(strings[i])
        }
}

func main() {
    names := []string{"bob", "sue", "alice"}
    printStrings(names...)
}

## Append
The built-in append function is used to dynamically add elements to a slice:
func append(slice []Type, elems ...Type) []Type
If the underlying array is not large enough, append() will create a new underlying array and point the returned slice to it.
Notice that append() is variadic, the following are all valid:
slice = append(slice, oneThing)
slice = append(slice, firstThing, secondThing)
slice = append(slice, anotherSlice...)

## Range
Go provides syntactic sugar to iterate easily over elements of a slice:
for INDEX, ELEMENT := range SLICE {
}
The element is a copy of the value at that index.
For example:
fruits := []string{"apple", "banana", "grape"}
for i, fruit := range fruits {
    fmt.Println(i, fruit)
}
// 0 apple
// 1 banana
// 2 grape

## Slice of Slices
Slices can hold other slices, effectively creating a matrix, or a 2D slice.
rows := [][]int{}
rows = append(rows, []int{1, 2, 3})
rows = append(rows, []int{4, 5, 6})
fmt.Println(rows)
// [[1 2 3] [4 5 6]]

## Tricky Slices
The append() function changes the underlying array of its parameter AND returns a new slice. This means that using append() on anything other than itself is usually a BAD idea.
// don't do this!
someSlice = append(otherSlice, element)
Take a look at these head-scratchers:
Example 1: Works As Expected
a := make([]int, 3)
fmt.Println("len of a:", len(a))
fmt.Println("cap of a:", cap(a))
// len of a: 3
// cap of a: 3

b := append(a, 4)
fmt.Println("appending 4 to b from a")
fmt.Println("b:", b)
fmt.Println("addr of b:", &b[0])
// appending 4 to b from a
// b: [0 0 0 4]
// addr of b: 0x44a0c0

c := append(a, 5)
fmt.Println("appending 5 to c from a")
fmt.Println("addr of c:", &c[0])
fmt.Println("a:", a)
fmt.Println("b:", b)
fmt.Println("c:", c)
// appending 5 to c from a
// addr of c: 0x44a180
// a: [0 0 0]
// b: [0 0 0 4]
// c: [0 0 0 5]
With slices a, b, and c, 4 and 5 seem to be appended as we would expect. We can even check the memory addresses and confirm that b and c point to different underlying arrays.
Example 2: Something Fishy
i := make([]int, 3, 8)
fmt.Println("len of i:", len(i))
fmt.Println("cap of i:", cap(i))
// len of i: 3
// cap of i: 8

j := append(i, 4)
fmt.Println("appending 4 to j from i")
fmt.Println("j:", j)
fmt.Println("addr of j:", &j[0])
// appending 4 to j from i
// j: [0 0 0 4]
// addr of j: 0x454000

g := append(i, 5)
fmt.Println("appending 5 to g from i")
fmt.Println("addr of g:", &g[0])
fmt.Println("i:", i)
fmt.Println("j:", j)
fmt.Println("g:", g)
// appending 5 to g from i
// addr of g: 0x454000
// i: [0 0 0]
// j: [0 0 0 5]
// g: [0 0 0 5]
In this example, however, when 5 is appended to i (creating g) it overwrites j's fourth index because j and g point to the same underlying array. The append() function only creates a new array when there isn't any capacity left. We created i with a length of 3 and a capacity of 8, which means we can append 5 items before a new array is automatically allocated.
Again, to avoid bugs like this, you should always use the append function on the same slice the result is assigned to:
mySlice := []int{1, 2, 3}
mySlice = append(mySlice, 4)

Questions:
1. Why is 5 the final value in the last index of 'j'?
The array’s cap() is exceeded so a new underlaying array is allocated
2. How can you best avoid these kinds of bugs?
Always assign the result of the append() function back to the same slide
3. Why is 5 the final value in the last index of 'j'?
j and g point to the same underlying array so g’s append overwrote j

- Exercise: Message Filter
Textio is introducing a feature that allows users to filter their messages based on specific criteria. For this feature, messages are categorized into three types: TextMessage, MediaMessage, and LinkMessage. Users can filter their messages to view only the types they are interested in.
Assignment
Your task is to implement a function that filters a slice of messages based on the message type.
Complete the filterMessages function. It should take a slice of Message interfaces and a string indicating the desired type ("text", "media", or "link"). It should return a new slice of Message interfaces containing only messages of the specified type.

- Exercise: Password Strength
As part of improving security, Textio wants to enforce a new password policy. A valid password must meet the following criteria:
    • At least 5 characters long but no more than 12 characters.
    • Contains at least one uppercase letter.
    • Contains at least one digit.
A string is really just a read-only slice of bytes. This means that you can use the same techniques you learned in previous lessons to iterate over the characters of a string.
Assignment
Implement the isValidPassword function by looping through each character in the password string. Make sure the password is long enough and includes at least one uppercase letter and one digit.
Assume passwords consist of ASCII characters only.
Tip
Remember that characters in Go strings are really just bytes under the hood. You can compare a character to another character like 'A' or '0' to check if it falls within a certain range.

## Exercise: Message Tagger
Textio needs a way to tag messages based on specific criteria.
Assignment
    1. Complete the tagMessages function. It should take a slice of sms messages, and a function (that takes an sms as input and returns a slice of strings) as inputs. And it should return a slice of sms messages.
        ◦ It should loop through each message and set the tags to the result of the passed in function.
        ◦ Be sure to modify the messages of the original slice using bracket notation messages[i].
See the tip below on how the strings package could be used here.
    2. Complete the tagger function. It should take an sms message and return a slice of strings.
        ◦ Return an initialized slice, even if no tags match. No nil slices.
        ◦ For any message that contains "urgent" (regardless of casing) in the content, the Urgent tag should be applied first.
        ◦ For any message that contains "sale" (regardless of casing), the Promo tag should be applied second.
Regardless of casing just means that even "uRGent" or "SALE" should trigger the tag.
Example usage:
messages := []sms{
        {id: "001", content: "Urgent! Last chance to see!"},
        {id: "002", content: "Big sale on all items!"},
        // Additional messages...
}
taggedMessages := tagMessages(messages, tagger)
// `taggedMessages` will now have tags based on the content.
// 001 = [Urgent]
// 002 = [Promo]