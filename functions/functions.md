# Functions

## Functions concept

Functions in Go can take zero or more arguments.
To make code easier to read, the variable type comes after the variable name.
For example, the following function:
func sub(x int, y int) int {
    return x-y
}
Accepts two integer parameters and returns another integer.
Here, func sub(x int, y int) int is known as the "function signature".

## Multiple Parameters

When multiple arguments are of the same type, and are next to each other in the function signature, the type only needs to be declared after the last argument.
Here are some examples:
func addToDatabase(hp, damage int) {
  // ...
}
func addToDatabase(hp, damage int, name string) {
  // ?
}
func addToDatabase(hp, damage int, name string, level int) {
  // ?
}

## Declaration Syntax
Developers often wonder why the declaration syntax in Go is different from the tradition established in the C family of languages.
C-Style Syntax
The C language describes types with an expression including the name to be declared, and states what type that expression will have.
int y;
The code above declares y as an int. In general, the type goes on the left and the expression on the right.
Interestingly, the creators of the Go language agreed that the C-style of declaring types in signatures gets confusing really fast - take a look at this nightmare.
int (*fp)(int (*ff)(int x, int y), int b)
Go-Style Syntax
Go's declarations are clear, you just read them left to right, just like you would in English.
x int
p *int
a [3]int
It's nice for more complex signatures, it makes them easier to read.

- Passing Variables by Value
Variables in Go are passed by value (except for a few data types we haven't covered yet). "Pass by value" means that when a variable is passed into a function, that function receives a copy of the variable. The function is unable to mutate the caller's original data.
func main() {
    x := 5
    increment(x)

    fmt.Println(x)
    // still prints 5,
    // because the increment function
    // received a copy of x
}

func increment(x int) {
    x++
}

- Ignoring Return Values
A function can return a value that the caller doesn't care about. We can explicitly ignore variables by using an underscore, or more precisely, the blank identifier _.
For example:
func getPoint() (x int, y int) {
    return 3, 4
}

// ignore y value
x, _ := getPoint()
Even though getPoint() returns two values, we can capture the first one and ignore the second. In Go, the blank identifier isn't just a convention; it's a real language feature that completely discards the value.
Why Might You Ignore a Return Value?
Maybe a function called getCircle returns the center point and the radius, but you only need the radius for your calculation. In that case, you can ignore the center point variable.
The Go compiler will return an error if you have any unused variable declarations in your code, so you need to ignore anything you don't intend to use.
- Named Return Values
Return values may be given names, and if they are, then they are treated the same as if they were new variables defined at the top of the function.
Named return values are best thought of as a way to document the purpose of the returned values.
According to the tour of go:
A return statement without arguments returns the named return values. This is known as a "naked" return. Naked return statements should be used only in short functions. They can harm readability in longer functions.
Named return values are what enable naked returns. Use naked returns only in short functions where the purpose of the returned values is obvious.
func getCoords() (x, y int) {
        // x and y are initialized with zero values

        return // automatically returns x and y
}

Is the same as:
func getCoords() (int, int) {
        var x int
        var y int
        return x, y
}
In the first example, x and y are the return values. At the end of the function, we could simply write return to return the values of those two variables, rather than writing return x,y.-
- The Benefits of Named Returns
Good for Documentation (Understanding)
Named return parameters are great for documenting a function. We know what the function is returning directly from its signature, no need for a comment.
Named return parameters are particularly important in longer functions with many return values.
func calculator(a, b int) (mul, div int, err error) {
    if b == 0 {
      return 0, 0, errors.New("can't divide by zero")
    }
    mul = a * b
    div = a / b
    return mul, div, nil
}
Which is easier to understand than:
func calculator(a, b int) (int, int, error) {
    if b == 0 {
      return 0, 0, errors.New("can't divide by zero")
    }
    mul := a * b
    div := a / b
    return mul, div, nil
}
We know the meaning of each return value just by looking at the function signature: func calculator(a, b int) (mul, div int, err error)
nil is the zero value of an error. More on this later.
Less Code (Sometimes)
If there are multiple return statements in a function, you don't need to write all the return values each time, though you probably should.
When you choose to omit return values, it's called a naked return. Naked returns should only be used in short and simple functions.
- Explicit Returns
Even though a function has named return values, we can still explicitly return values if we want to.
func getCoords() (x, y int) {
        return x, y // this is explicit
}
Using this explicit pattern we can even overwrite the return values:
func getCoords() (x, y int) {
    return 5, 6 // this is explicit, x and y are NOT returned
}
Otherwise, if we want to return the values defined in the function signature we can just use a naked return (blank return):
func getCoords() (x, y int) {
    return // implicitly returns x and y
}

- Early Returns
Go supports the ability to return early from a function. This is a powerful feature that can clean up code, especially when used as guard clauses.
Guard Clauses leverage the ability to return early from a function (or continue through a loop) to make nested conditionals one-dimensional. Instead of using if/else chains, we just return early from the function at the end of each conditional block.
func divide(dividend, divisor int) (int, error) {
        if divisor == 0 {
                return 0, errors.New("can't divide by zero")
        }
        return dividend/divisor, nil
}

- Functions As Values
Go supports first-class and higher-order functions, which are just fancy ways of saying "functions as values". Functions are just another type -- like ints and strings and bools.
Let's assume we have two simple functions:
func add(x, y int) int {
        return x + y
}

func mul(x, y int) int {
        return x * y
}
We can create a new aggregate function that accepts a function as its 4th argument:
func aggregate(a, b, c int, arithmetic func(int, int) int) int {
  firstResult := arithmetic(a, b)
  secondResult := arithmetic(firstResult, c)
  return secondResult
}
It calls the given arithmetic function (which could be add or mul, or any other function that accepts two ints and returns an int) and applies it to three inputs instead of two. It can be used like this:
func main() {
        sum := aggregate(2, 3, 4, add)
        // sum is 9
        product := aggregate(2, 3, 4, mul)
        // product is 24
}

- Anonymous Functions
Anonymous functions are true to form in that they have no name. They're useful when defining a function that will only be used once or to create a quick closure.
Let's say we have a function conversions that accepts another function, converter as input:
func conversions(converter func(int) int, x, y, z int) (int, int, int) {
        convertedX := converter(x)
        convertedY := converter(y)
        convertedZ := converter(z)
        return convertedX, convertedY, convertedZ
}
We could define a function normally and then pass it in by name... but it's usually easier to just define it anonymously:
func double(a int) int {
    return a + a
}

func main() {
    // using a named function
        newX, newY, newZ := conversions(double, 1, 2, 3)
        // newX is 2, newY is 4, newZ is 6

    // using an anonymous function
        newX, newY, newZ = conversions(func(a int) int {
            return a + a
        }, 1, 2, 3)
        // newX is 2, newY is 4, newZ is 6
}

- Defer
The defer keyword is a fairly unique feature of Go. It allows a function to be executed automatically just before its enclosing function returns. The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
Deferred functions are typically used to clean up resources that are no longer being used. Often to close database connections, file handlers and the like.
For example:
func GetUsername(dstName, srcName string) (username string, err error) {
        // Open a connection to a database
        conn, _ := db.Open(srcName)

        // Close the connection *anywhere* the GetUsername function returns
        defer conn.Close()

        username, err = db.FetchUser()
        if err != nil {
                // The defer statement is auto-executed if we return here
                return "", err
        }

        // The defer statement is auto-executed if we return here
        return username, nil
}
In the above example, the conn.Close() function is not called here:
defer conn.Close()
It's called:
// here
return "", err
// or here
return username, nil
Depending on whether the FetchUser function errored. (We'll cover errors later).
Defer is a great way to make sure that something happens before a function exits, even if there are multiple return statements, a common occurrence in Go.
Multiple Defers
The location of a defer statement inside a function matters. The deferred call is registered at the point where defer is executed, and it will run when the function returns. If you have multiple defer statements in a single function, they are executed in last-in, first-outorder (the last deferred call runs first).
For example, you'd want to close a file before trying to remove it:
func CreateTempFile() {
        f, _ := os.Create("temp-42.txt")
        defer os.Remove(f.Name()) // executed second
        defer f.Close()           // executed first

        fmt.Fprintln(f, "How many roads must a man walk down?")
}

- Block Scope
Unlike Python, Go is not function-scoped, it's block-scoped. Variables declared inside a block are only accessible within that block (and its nested blocks). There's also the package scope. We'll talk about packages later, but for now, you can think of it as the outermost, nearly global scope.
package main

// scoped to the entire "main" package (basically global)
var age = 19

func sendEmail() {
    // scoped to the "sendEmail" function
    name := "Jon Snow"

    for i := 0; i < 5; i++ {
        // scoped to the "for" body
        email := "snow@winterfell.net"
    }
}
Blocks are defined by curly braces {}. New blocks are created for:
    • Functions
    • Loops
    • If statements
    • Switch statements
    • Select statements
    • Explicit blocks
It's a bit unusual, but occasionally you'll see a plain old explicit block. It exists for no other reason than to create a new scope.
package main

import "fmt"

func main() {
    {
        age := 19
        // this is okay
        fmt.Println(age)
    }

    // this is not okay
    // the age variable is out of scope
    fmt.Println(age)
}

- Exercise: Processing Orders
Management thinks our branding is so creative that our SaaS customers will pay for Textio merch.
Assignment
Complete the placeOrder function.
It returns a bool indicating whether the order was successful (true is a success) and a float64 representing the user's balance after the order. The placeOrder function should always return the account balance regardless of whether it was adjusted.
The amountInStock and calcPrice functions can be used to look up the current stock and price of an item.
    • If the quantity is greater than the amount in stock, the order should be rejected.
    • If the user doesn't have enough money in their account, the order should be rejected.
    • Otherwise, the order should be accepted and you should return the new balance.

- Closures
A closure is a function that references variables from outside its own function body. The function may access and assign to the referenced variables.
In this example, the concatter() function returns a function that has reference to an enclosed doc value. Each successive call to harryPotterAggregator mutates that same doc variable.
func concatter() func(string) string {
        doc := ""
        return func(word string) string {
                doc += word + " "
                return doc
        }
}

func main() {
        harryPotterAggregator := concatter()
        harryPotterAggregator("Mr.")
        harryPotterAggregator("and")
        harryPotterAggregator("Mrs.")
        harryPotterAggregator("Dursley")
        harryPotterAggregator("of")
        harryPotterAggregator("number")
        harryPotterAggregator("four,")
        harryPotterAggregator("Privet")

        fmt.Println(harryPotterAggregator("Drive"))
        // Mr. and Mrs. Dursley of number four, Privet Drive
}

- Currying
Function currying is a concept from functional programming and involves partial applicationof functions. It allows a function with multiple arguments to be transformed into a sequence of functions, each taking a single argument.
Let's simulate this behavior. For example:
func main() {
  squareFunc := selfMath(multiply)
  doubleFunc := selfMath(add)

  fmt.Println(squareFunc(5))
  // prints 25

  fmt.Println(doubleFunc(5))
  // prints 10
}

func multiply(x, y int) int {
  return x * y
}

func add(x, y int) int {
  return x + y
}

func selfMath(mathFunc func(int, int) int) func (int) int {
  return func(x int) int {
    return mathFunc(x, x)
  }
}
In the example above:
    • selfMath(multiply) returns a new function and stores it in squareFunc; it doesn't run multiply yet.
    • squareFunc(5) runs the returned function, which calls multiply(5, 5).
Assignment
The Textio API needs a very robust error-logging system so we can see when things are going awry in the back-end system. We need a function that can create a custom "logger" (a function that prints to the console) given a specific formatter.
These errors are test data, not runtime failures.
Complete the getLogger function. It should take as input a formatter function and return a new function. The new logger function takes as input two strings and passes them to the formatter, then prints the result. Keep the order of the strings.