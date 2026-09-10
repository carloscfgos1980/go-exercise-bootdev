# Generics assigments

## Generics in Go. Get last

At Textio we store all the emails for a campaign in memory as a slice. We store payments for a single user in the same way.

Complete the getLast() function. It should be a generic function that returns the last element from a slice, no matter the types stored in the slice. If the slice is empty, it should return the zero value of the type.

Tip: Zero Value of a Type
Creating a variable that's the zero value of a type is easy:

var myZeroInt int

It's the same with generics, we just have a variable that represents the type:

var myZero T

## Contraints. Charge for  Line Items

We have different kinds of "line items" that we charge our customer's credit cards for. Line items can be things like "subscriptions" or "one-time payments" for email usage.

Complete the chargeForLineItem function.

Check if the user has a balance with enough funds to be able to pay for the cost of the newItem.
If they don't, then return an "insufficient funds" error and zero values for the other return values.
If they do have enough funds:
Add the line item to the user's history by appending the newItem to the slice of oldItems. This new slice is your first return value.
Calculate the user's new balance by subtracting the cost of the new item from their balance. This is your second return value.

## Parametric Constraints

The chief architect at Textio has decided she wants to implement billing with generics. Specifically, she wants us to create a new biller interface. A biller is an interface that can be used to charge a customer, and it can also report its name.

There are two kinds of billers:

userBiller (cheaper)
orgBiller (more expensive)
A customer is either a user or an org. A user will be billed with a userBiller and an org with an orgBiller.

Create the new biller interface. It should have 2 methods:

Charge
Name
The good news is that the architect already wrote the userBiller and orgBiller types for us that fulfill this new biller interface. Use the definitions of those types and their methods to figure out how to write the biller interface definition.