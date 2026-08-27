# Condicionals

## Conditionals
if statements in Go do not use parentheses around the condition:
if height > 4 {
    fmt.Println("You are tall enough!")
}
else if and else are supported as you might expect:
if height > 6 {
    fmt.Println("You are super tall!")
} else if height > 4 {
    fmt.Println("You are tall enough!")
} else {
    fmt.Println("You are not tall enough!")
}
Unlike other languages, you must put the opening brace on the same line as the condition and not on a new line.

## The Initial Statement of an If Block
An if conditional can have an "initial" statement. The variable(s) created in the initial statement are only defined within the scope of the if, else if, and else blocks.

if INITIAL_STATEMENT; CONDITION {
}

Why Would I Use This?
It has two valuable purposes:
    1. It's a bit shorter
    2. It limits the scope of the initialized variable(s) to the if statement
For example, instead of writing:
length := getLength(email)
if length < 10 {
    fmt.Printf("Email must be at least 10 characters, is %d\n", length)
}
We can do:
if length := getLength(email); length < 10 {
    fmt.Printf("Email must be at least 10 characters, is %d\n", length)
}
In the example above, length isn't available in the parent scope, which is nice because we don't need it there - we won't accidentally use it elsewhere in the function. It would still be available in any else if or else blocks attached to that if.

## Switch
Switch statements are a way to compare a value against multiple options. They are similar to if-else statements but are more concise and readable when the number of options is more than 2.
func getCreator(os string) string {
    var creator string
    switch os {
    case "linux":
        creator = "Linus Torvalds"
    case "windows":
        creator = "Bill Gates"
    case "mac":
        creator = "A Steve"
    default:
        creator = "Unknown"
    }
    return creator
}

Notice that in Go, the break statement is not required at the end of a case to stop it from falling through to the next case. The break statement is implicit in Go.
If you do want a case to fall through to the next case, you can use the fallthroughkeyword.
func getCreator(os string) string {
    var creator string
    switch os {
    case "linux":
        creator = "Linus Torvalds"
    case "windows":
        creator = "Bill Gates"

    // all three of these cases will set creator to "A Steve"
    case "macOS":
        fallthrough
    case "Mac OS X":
        fallthrough
    case "mac":
        creator = "A Steve"

    default:
        creator = "Unknown"
    }
    return creator
}
The default case does what you'd expect: it's the case that runs if none of the other cases match.