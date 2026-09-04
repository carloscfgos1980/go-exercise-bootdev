# Loops

## Loops in Go
The basic loop in Go is written in standard C-like syntax:
for INITIAL; CONDITION; AFTER{
  // do something
}
INITIAL is run once at the beginning of the loop and can create
variables within the scope of the loop.
CONDITION is checked before each iteration. If the condition doesn't pass
then the loop breaks.
AFTER is run after each iteration.
For example:
for i := 0; i < 10; i++ {
  fmt.Println(i)
}
// Prints 0 through 9

## Omitting Conditions from a for Loop in Go
Loops in Go can omit sections of a for loop. For example, the CONDITION (middle part) can be omitted which causes the loop to run forever.
for INITIAL; ; AFTER {
  // do something forever
}

## There Is No While Loop in Go
Most programming languages have a concept of a while loop. Because Go allows for the omission of sections of a for loop, a while loop is just a for loop that only has a CONDITION.
for CONDITION {
  // do some stuff while CONDITION is true
}
For example:
plantHeight := 1
for plantHeight < 5 {
  fmt.Println("still growing! current height:", plantHeight)
  plantHeight++
}
fmt.Println("plant has grown to ", plantHeight, "inches")
Which prints:
still growing! current height: 1
still growing! current height: 2
still growing! current height: 3
still growing! current height: 4
plant has grown to 5 inches
Go also allows you to omit the condition entirely to create an infinite loop:
for {
  // do some stuff forever
}

## Exercise: Fizzbuzz
Go supports the standard modulo operator:
7 % 3 // 1
The AND logical operator:
true && false // false
true && true // true
As well as the OR operator:
true || false // true
false || false // false

Assignment
We're hiring engineers at Textio, so time to brush up on the classic "Fizzbuzz" game, a coding exercise that has been dramatically overused in coding interviews across the world.
Complete the fizzbuzz function that prints the numbers 1 to 100 inclusive each on their own line, but replace multiples of 3 with the text fizz and multiples of 5 with buzz. Print fizzbuzz for multiples of 3 AND 5.

## Continue & Break
Whenever we want to change the control flow of a loop we can use the continue and break keywords.
continue
The continue keyword stops the current iteration of a loop and continues to the next iteration. continue is a powerful way to use the guard clause pattern within loops.
for i := 0; i < 10; i++ {
  if i % 2 == 0 {
    continue
  }
  fmt.Println(i)
}
// 1
// 3
// 5
// 7
// 9
break
The break keyword stops the current iteration of a loop and exits the loop.
for i := 0; i < 10; i++ {
  if i == 5 {
    break
  }
  fmt.Println(i)
}
// 0
// 1
// 2
// 3
// 4

exercise: Connections
Textio has group chats that make communicating with multiple people much more efficient--if the chat doesn't descend into chaos. Instead of sending the message multiple times to individual people, you send one message to all of them at once.
Assignment
Complete the countConnections function that takes an integer groupSizerepresenting the number of people in the group chat and returns an integer representing the number of connections between them. For each additional person in the group, the number of new connections continues to grow. Use a for loop to accumulate the number of connections instead of directly using a mathematical formula.
To make sure you're picturing it right:
    • If there are two people, how many possible connections exist between them?
    • If you add a third person, how many new connections are created with the rest?