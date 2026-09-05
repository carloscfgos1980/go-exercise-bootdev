# Slices Assigments

## Arrays. Get messages with retries

When our clients don't respond to a message, they can be reminded with up to 2 additional messages.

Complete the getMessageWithRetries function. It takes three strings and returns:

An array of 3 strings
An array of 3 integers
The returned string array contains the original messages. The first is the primary message, the second is the first reminder, and the third is the last reminder.

The integers in the integer array represent the cost of sending each message. The cost of each message is equal to the length of the message, plus the length of any previous messages. For example:

"hello" costs 5
"world" costs 5, adding "hello" makes total cost 10 (5 + 5)
"!" costs 1, adding the previous messages makes total cost 11 (5 + 5 + 1)

## Slices in Go. Get messages with retries for plan

Retries are a premium feature now! Textio's free users only get 1 retry message, while pro members get an unlimited amount.

Complete the getMessageWithRetriesForPlan function. It takes a plan variable as input as well as an array of 3 messages. You've been provided with constants representing the plan types at the top of the file.

If the plan is a pro plan, return all the strings from the messages input in a slice.
If the plan is a free plan, return the first 2 strings from the messages input in a slice.
If the plan isn't either of those, return a nil slice and an error that says unsupported plan.

## Make. Get messages cost

We send a lot of text messages at Textio, and our API is getting slow and unresponsive.

If we know the rough size of a slice before we fill it up, we can make our program faster by creating the slice with that size ahead of time so that the Go runtime doesn't need to continuously allocate new underlying arrays of larger and larger sizes. By setting the length, the slice can still be resized later, but it means we can avoid all the expensive resizing since we know what we'll need.

Complete the getMessageCosts() function. It takes a slice of messages and returns a slice of message costs.

Preallocate a slice for the message costs of the same length as the messages slice.
Fill the costs slice with costs for each message. The cost in the cost slice should correspond to the message in the messages slice at the same index. The cost of a message is the length of the message multiplied by 0.01.

## Variadic. Sums

We need to sum up the costs of all individual messages so we can send an end-of-month bill to our customers.

Complete the sum function to return the sum of all inputs.

Take note of how the variadic inputs and the spread operator are used in the test suite.

## Append. Get Cost Day

We've been asked to add a feature that extracts costs for a given day.

Complete the getDayCosts() function. It accepts a slice of cost structs and a day int, and it returns a float64 slice containing that day's costs.

Create an empty, non-nil float64 slice.
Use append() to add each cost's value when its day matches the day argument.
Return the slice.

## Range Index of the first bad word

We need to be able to quickly detect bad words in the messages our system sends.

Complete the indexOfFirstBadWord function. It takes 2 arguments:

msg: a single message represented as a slice of words
badWords: a slice of bad words
If it finds any bad words in the message it should return the index of the first bad word in the msg slice. This will help us filter out naughty words from our messaging system. If no bad words are found, return -1 instead.

Use the range keyword.

Tip
Remember, you can ignore returned values with an underscore _ instead of creating an unused variable.

## Slices of slices. Create matrix

We support various visualization dashboards on Textio that display message analytics to our users. The UI for our graphs and charts is built on top of a grid system. Let's build some grid logic.

Complete the createMatrix function. It takes a number of rows and columns and returns a 2D slice of integers. The value of each cell is i * j where i and j are the indexes of the row and column respectively. Basically, we're building a multiplication chart.

For example, a 5x10 matrix, produced from calling createMatrix(5, 10), would look like this:

[0  0  0  0  0  0  0  0  0  0]
[0  1  2  3  4  5  6  7  8  9]
[0  2  4  6  8 10 12 14 16 18]
[0  3  6  9 12 15 18 21 24 27]
[0  4  8 12 16 20 24 28 32 36]

## Message Filter
Textio is introducing a feature that allows users to filter their messages based on specific criteria. For this feature, messages are categorized into three types: TextMessage, MediaMessage, and LinkMessage. Users can filter their messages to view only the types they are interested in.

Assignment
Your task is to implement a function that filters a slice of messages based on the message type.

Complete the filterMessages function. It should take a slice of Message interfaces and a string indicating the desired type ("text", "media", or "link"). It should return a new slice of Message interfaces containing only messages of the specified type.

## Password Strength
As part of improving security, Textio wants to enforce a new password policy. A valid password must meet the following criteria:

At least 5 characters long but no more than 12 characters.
Contains at least one uppercase letter.
Contains at least one digit.
A string is really just a read-only slice of bytes. This means that you can use the same techniques you learned in previous lessons to iterate over the characters of a string.
Assignment
Implement the isValidPassword function by looping through each character in the password string. Make sure the password is long enough and includes at least one uppercase letter and one digit.

Assume passwords consist of ASCII characters only.

## Message Tagger
Textio needs a way to tag messages based on specific criteria.

Assignment
Complete the tagMessages function. It should take a slice of sms messages, and a function (that takes an sms as input and returns a slice of strings) as inputs. And it should return a slice of sms messages.

It should loop through each message and set the tags to the result of the passed in function.
Be sure to modify the messages of the original slice using bracket notation messages[i].
See the tip below on how the strings package could be used here.
Complete the tagger function. It should take an sms message and return a slice of strings.

Return an initialized slice, even if no tags match. No nil slices.
For any message that contains "urgent" (regardless of casing) in the content, the Urgent tag should be applied first.
For any message that contains "sale" (regardless of casing), the Promo tag should be applied second.
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