# POINTERS

## Introduction to Pointers
As we have learned, a variable is a named location in memory that stores a value. We can manipulate the value of a variable by assigning a new value to it or by performing operations on it. When we assign a value to a variable, we are storing that value in a specific location in memory.
x := 42
// "x" is the name of a location in memory. That location is storing the integer value of 42
A Pointer Is a Variable
A pointer is a variable that stores the memory address of another variable. This means that a pointer "points to" the location of where the data is stored, not the actual data itself.
The * syntax defines a pointer:
var p *int
The & operator generates a pointer to its operand.
myString := "hello"
myStringPtr := &myString

## References
It's possible to define an empty pointer. For example, an empty pointer to an integer:
var p *int

fmt.Printf("value of p: %v\n", p)
// value of p: <nil>
Its zero value is nil, which means it doesn't point to any memory address. Empty pointers are also called "nil pointers".
Instead of starting with a nil pointer, it's common to use the & operator to get a pointer to its operand:
myString := "hello"      // myString is just a string
myStringPtr := &myString // myStringPtr is a pointer to myString's address

fmt.Printf("value of myStringPtr: %v\n", myStringPtr)
// value of myStringPtr: 0x140c050
Dereference
The * operator dereferences a pointer to get the original value.
*myStringPtr = "world"                              // set myString through the pointer
fmt.Printf("value of myString: %s\n", *myStringPtr) // read myString through the pointer
// value of myString: world
Unlike C, Go has no pointer arithmetic (which is covered in our Learn Memory Management course if you haven't taken it already).

## Pass by Reference
Functions in Go generally pass variables by value, meaning that functions receive a copy of most non-composite types:
func increment(x int) {
    x++
    fmt.Println(x)
    // 6
}

func main() {
    x := 5
    increment(x)
    fmt.Println(x)
    // 5
}
The main function still prints 5 because the increment function received a copy of x.
One of the most common use cases for pointers in Go is to pass variables by reference, meaning that the function receives the address of the original variable, not a copy of the value. This allows the function to update the original variable's value.
func increment(x *int) {
    *x++
    fmt.Println(*x)
    // 6
}

func main() {
    x := 5
    increment(&x)
    fmt.Println(x)
    // 6
}
Fields of Pointers
When your function receives a pointer to a struct, you might try to access a field like this and encounter an error:
msgTotal := *analytics.MessagesTotal
Instead, access it – like you'd normally do – using a selector expression.
msgTotal := analytics.MessagesTotal
This approach is the recommended, simplest way to access struct fields in Go, and is shorthand for:
(*analytics).MessagesTotal

## Pointers Quiz
package main

func main() {
        var x int = 50
        var y *int = &x
        *y = 100
}

Questions:
1. What is the value of *y after the code on the left executes?
100
2. What is the value of x after the entire code block on the left executes?
100

- Nil Pointers
Pointers can be very dangerous.
If a pointer points to nothing (the zero value of the pointer type is nil) then dereferencing it will cause a runtime error (a panic) that crashes the program. Generally speaking, whenever you're dealing with pointers you should check if it's nil before trying to dereference it.

- Pointer Receivers
A receiver type on a method can be a pointer.
Methods with pointer receivers can modify the value to which the receiver points. Since methods often need to modify their receiver, pointer receivers are more common than value receivers. However, methods with pointer receivers don't require that a pointer is used to call the method. The pointer will automatically be derived from the value.
Pointer Receiver
type car struct {
        color string
}

func (c *car) setColor(color string) {
        c.color = color
}

func main() {
        c := car{
                color: "white",
        }
        c.setColor("blue")
        fmt.Println(c.color)
        // prints "blue"
}
Non-Pointer Receiver
type car struct {
        color string
}

func (c car) setColor(color string) {
        c.color = color
}

func main() {
        c := car{
                color: "white",
        }
        c.setColor("blue")
        fmt.Println(c.color)
        // prints "white"
}
The non-pointer receiver example prints "white" instead of "blue" because the method receives a copy of the struct. Without using a pointer receiver, any changes made inside the method only affect that copy, not the original.

## Pointer Receiver Code
Methods with pointer receivers don't require that a pointer is used to call the method. The pointer will automatically be derived from the value.
type circle struct {
        x int
        y int
    radius int
}

func (c *circle) grow() {
    c.radius *= 2
}

func main() {
    c := circle{
        x: 1,
        y: 2,
        radius: 4,
    }

    // notice c is not a pointer in the calling function
    // but the method still gains access to a pointer to c
    c.grow()
    fmt.Println(c.radius)
    // prints 8
}

## Pointer Performance
Occasionally, new Go developers hear "pointers don't pass copies" and take that to a logical extreme, concluding:
Pointers are always faster because copying is slow. I'll always use pointers!
No. Bad. Stop.
Here are my rules of thumb:
    1. First, worry about writing clear, correct, maintainable code.
    2. If you have a performance problem, fix it.
Before even thinking about using pointers to optimize your code, use pointers when you need a shared reference to a value; otherwise, just use values.
If you do have a performance problem, consider:
    1. Stack vs. Heap
    2. Copying
Interestingly, local non-pointer variables are generally faster to pass around than pointers because they're stored on the stack, which is faster to access than the heap. Even though copying is involved, the stack is so fast that it's no big deal.
Once the value becomes large enough that copying is the greater problem, it can be worth using a pointer to avoid copying. That value will probably go to the heap, so the gain from avoiding copying needs to be greater than the loss from moving to the heap.
One of the reasons Go programs tend to use less memory than Java and C# programs is that Go tends to allocate more on the stack.
Questions:
Question 1: Where do values pointed to by pointers usually live? (Heap)
In Go, when you take a pointer to a value, that value often can't stay on the stack. Here's why: the stack is organized like a stack of plates, tightly tied to function calls. When a function returns, its stack frame gets popped off and reused. But if you have a pointer to a variable inside that function, and something outside the function might still use that pointer after the function returns, the variable can't disappear when the stack frame does.
Go's compiler does what's called "escape analysis" to detect this situation. If a value's lifetime might outlast the function that created it (because something holds a pointer to it), the compiler makes that value "escape" to the heap instead, where it can stick around independently of any particular function call.
The heap is slower to allocate to and read from than the stack, but it's flexible: things there persist as long as something references them.
Question 2: Which is typically faster to pass to a function, especially for small data? (Value)
Passing a value means copying it, and copying small data on the stack is extremely fast. It also avoids pushing that data to the heap. Passing a pointer instead means:
    1. The value it points to likely escapes to the heap (slower allocation).
    2. You add a layer of indirection (the CPU has to follow the pointer to get the actual data).
For small values, the cost of copying on the stack is cheaper than the cost of heap allocation plus indirection. Pointers only start winning performance-wise when the data being copied is large enough that copying costs more than the heap/indirection overhead.
This is the core message of the lesson: pointers aren't a free performance win. They're a tool for sharing references, and only become a performance optimization once your data is large enough that the copy cost outweighs the heap cost.

## Exercise: Update Balance
Textio needs a new way to update user's account balance.
Assignment
Implement the updateBalance function. It should take a customer pointer and a transaction, and return an error. Depending on the transactionType, it should either add or subtract the amount from the customer's balance. If the customer does not have enough money, it should return the error insufficient funds and leave the balance unchanged. If the transactionType isn't a withdrawal or deposit, it should return the error unknown transaction type. Otherwise, it should process the transaction and return nil.
alice := customer{id: 1, balance: 100.0}
deposit := transaction{customerID: 1, amount: 50, transactionType: transactionDeposit}

updateBalance(&alice, deposit)
// id: 1 balance: 150