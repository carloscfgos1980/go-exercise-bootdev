# LOOPS ASSIGMENTS

## At Textio we have a dynamic formula for determining how much a batch of bulk messages costs to send. Complete the bulkSend() function.

It should return the total cost (as a float64) to send a batch of numMessages messages. Each message costs 1.0, plus an additional fee. The fee structure is:

1st message: 1.0 + 0.00
2nd message: 1.0 + 0.01
3rd message: 1.0 + 0.02
4th message: 1.0 + 0.03
...
Use a loop to calculate the total cost and return it.

## There Is No While Loop in Go. Get max message to send

We have an interesting new cost structure from our SMS vendor. They charge exponentially more money for each consecutive text we send! Let's write a function that calculates how many messages we can send in a given batch given a costMultiplier and a budgetInPennies.

In a nutshell, the first message costs a penny, and each message after that costs the same as the previous message multiplied by the costMultiplier.

There is a bug in the code! Add a condition to the for loop to fix the bug. The loop should stop when balance is equal to or less than 0. So what condition should the for loop evaluate?

## Fizzbuzz

We're hiring engineers at Textio, so time to brush up on the classic "Fizzbuzz" game, a coding exercise that has been dramatically overused in coding interviews across the world.

Complete the fizzbuzz function that prints the numbers 1 to 100 inclusive each on their own line, but replace multiples of 3 with the text fizz and multiples of 5 with buzz. Print fizzbuzz for multiples of 3 AND 5.

## Continue & Break. Prime numbers

As an Easter egg, we decided to reward our users with a free text message if they send a prime number of text messages this year.

Complete the printPrimes function. It should print all of the prime numbers up to and including max. It should skip any numbers that are not prime.

Here's the pseudocode:

printPrimes(max):
  for n in range(2, max+1):
    if n is 2:
      n is prime, print it and skip to next n
    if n is even:
      n is not prime, skip to next n
    isPrime = true
    for i in range (3, sqrt(n) + 1, 2):
      if n can be evenly divided by i:
        isPrime = false
        break from inner loop
    if isPrime:
      print n

Breakdown
This is a primality test.

We skip even numbers because they can't be prime
We only check up to the square root of n. A factor higher than the square root of n must multiply with a factor lower than the square root of n, meaning we only need to check up to the square root of n for potential factors.
In your code, you can set the stop condition as i * i <= n
We start checking at 2 because 1 is not prime

## Connections
Textio has group chats that make communicating with multiple people much more efficient--if the chat doesn't descend into chaos. Instead of sending the message multiple times to individual people, you send one message to all of them at once.

Assignment
Complete the countConnections function that takes an integer groupSize representing the number of people in the group chat and returns an integer representing the number of connections between them. For each additional person in the group, the number of new connections continues to grow. Use a for loop to accumulate the number of connections instead of directly using a mathematical formula.

To make sure you're picturing it right:

If there are two people, how many possible connections exist between them?
If you add a third person, how many new connections are created with the rest?