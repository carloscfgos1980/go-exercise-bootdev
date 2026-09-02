# Structs Assigments

## Nested structs. Send message check
Textio has a bug, we've been sending texts that are missing critical bits of information! Before we send text messages in Textio, we must check to make sure the required fields have non-zero values.

Notice that the user struct is a nested struct within the messageToSend struct. Both sender and recipient are user struct types.

Complete the canSendMessage function. It should return true only if the sender and recipient fields each contain a name and a number. If any of the default zero values are present, return false instead.

## Struct Methods in Go. authentication info
Let's clean up Textio's authentication logic. We store our user's authentication data inside an authenticationInfo struct. We need a method that can take that data and return a basic authorization string.

The format of the string should be:

Authorization: Basic USERNAME:PASSWORD

Create a method on the authenticationInfo struct called getBasicAuth that returns the formatted string.

## Update Users
We need a way to differentiate between standard and premium users. When a new user is created, they need a membership type, and that type will determine the message character limit.

Assignment
Create a new struct called Membership, it should have:
A Type string field
A MessageCharLimit integer field
Update the User struct to embed a Membership.
Complete the newUser function. It should return a new User with all the fields set as you would expect based on the inputs. If the user is a "premium" member, the MessageCharLimit should be 1000, otherwise, it should only be 100.

## Send Message
Assignment
Create a SendMessage method for the User struct.

It should take a message string and messageLength int as inputs.

If the messageLength is within the user's MessageCharLimit, return the original message and true (indicating the message can be sent), otherwise, return an empty string and false.