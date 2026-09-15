# Enums assigments

## Lack of enums. Handle email bouncd

A lazy Go programmer wrote the handleEmailBounce function... just because the compiler doesn't force us to check errors doesn't mean we shouldn't!

Take a look at the updateStatus and track methods in the user.go file. Handle their errors properly, and use the fmt.Errorf function and the %w formatting verb to add useful context to the errors.

If updateStatus fails, return an error saying error updating user status: ERR
If track fails, return an error saying error tracking user bounce: ERR
Where ERR is the error returned by the method.

## Iota. Email status
Define an emailStatus type that uses iota syntax to represent the following states:

EmailBounced: 0
EmailInvalid: 1
EmailDelivered: 2
EmailOpened: 3