# ERRORS

## The Error Interface
Go programs express errors with error values. An Error is any type that implements the simple built-in error interface:
type error interface {
    Error() string
}
When something can go wrong in a function, that function should return an error as its last return value. Any code that calls a function that can return an error should handle errors by testing whether the error is nil.
Atoi
Let's look at how the strconv.Atoi function uses this pattern. The function signature of Atoi is:
func Atoi(s string) (int, error)
This means Atoi takes a string argument and returns two values: an integer and an error. If the string can be successfully converted to an integer, Atoi returns the integer and a nil error. If the conversion fails, it returns zero and a non-nil error.
Here's how you would safely use Atoi:
// Atoi converts a stringified number to an integer
i, err := strconv.Atoi("42b")
if err != nil {
    fmt.Println("couldn't convert:", err)
    // because "42b" isn't a valid integer, we print:
    // couldn't convert: strconv.Atoi: parsing "42b": invalid syntax
    // Note:
    // 'parsing "42b": invalid syntax' is returned by the .Error() method
    return
}
// if we get here, then the
// variable i was converted successfully
A nil error denotes success; a non-nil error denotes failure.

## Formatting Strings Review
A convenient way to format strings in Go is by using the standard library's fmt.Sprintf()function. It's a string interpolation function, similar to Python's f-strings. The %v substring uses the type's default formatting, which is often what you want.
Default Values
const name = "Kim"
const age = 22
s := fmt.Sprintf("%v is %v years old.", name, age)
// s = "Kim is 22 years old."
The equivalent Python code:
name = "Kim"
age = 22
s = f"{name} is {age} years old."
# s = "Kim is 22 years old."
Rounding Floats
s := fmt.Sprintf("I am %f years old", 10.523)
// s = I am 10.523000 years old

// The ".2" rounds the number to 2 decimal places
s := fmt.Sprintf("I am %.2f years old", 10.523)
// s = I am 10.52 years old

## Custom Error interface's
Because errors are just interfaces, you can build your own custom types that implement the error interface. Here's an example of a userError struct that implements the errorinterface:
type userError struct {
    name string
}

func (e userError) Error() string {
    return fmt.Sprintf("%v has a problem with their account", e.name)
}
It can then be used as an error:
func sendSMS(msg, userName string) error {
    if !canSendToUser(userName) {
        return userError{name: userName}
    }
    ...
}


## Errors Quiz
Go programs express errors with error values. Error-values are any type that implements the simple built-in error interface.
Keep in mind that the way Go handles errors is fairly unique. Most languages treat errors as something special and different. For example, Python raises exception types and JavaScript throws and catches errors. In Go, an error is just another value that we handle like any other value - however we want! There aren't any special keywords for dealing with them.
Questions:
1. What is the underlying type of an error?
Interfaces
2. Can a type be an error and also fulfill another interface?
YES

## The Errors Package
The Go standard library provides an "errors" package that makes it easy to deal with errors.
Read the godoc for the errors.New() function, but here's a simple example:
var err error = errors.New("something went wrong")

## Panic
As we've seen, the proper way to handle errors in Go is to make use of the error interface. Pass errors up the call stack, treating them as normal values:
func enrichUser(userID string) (User, error) {
    user, err := getUser(userID)
    if err != nil {
        // fmt.Errorf is GOATed: it wraps an error with additional context
        return User{}, fmt.Errorf("failed to get user: %w", err)
    }
    return user, nil
}
However, there is another way to deal with errors in Go: the panic function. When a function calls panic, the program crashes and prints a stack trace.
As a general rule, do not use panic!
The panic function yeets control out of the current function and up the call stack until it reaches a function that defers a recover. If no function calls recover, the goroutine (often the entire program) crashes.
func enrichUser(userID string) User {
    user, err := getUser(userID)
    if err != nil {
        panic(err)
    }
    return user
}

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("recovered from panic:", r)
        }
    }()

    // this panics, but the defer/recover block catches it
    // a truly astonishingly bad way to handle errors
    enrichUser("123")
}
Sometimes new Go developers look at panic/recover, and think, "This is like try/catch! I like this!" Don't be like them.
I use error values for all "normal" error handling, and if I have a truly unrecoverable error, I use log.Fatal to print a message and exit the program.
Questions
1. Panic and recover should be used instead of errors…
Basically never
2. If you want your program to cleanly exit in an unrecoverable way, which is a good alternative to panic?
Log.Fatal()

## Exercise: User Input
In Textio, users can set their profile status to communicate their current activity to those that choose to read it... However, there are some restrictions on what these statuses can contain. Your task is to implement a function that validates a user's status update. The status update cannot be empty and must not exceed 140 characters.
Assignment
Complete the validateStatus function. It should return an error when the status update violates any of the rules:
    • If the status is empty, return an error that reads status cannot be empty.
    • If the status exceeds 140 characters, return an error that says status exceeds 140 characters.
