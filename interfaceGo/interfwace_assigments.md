# Interface assigments

## Interfaces in Go. Send messages
The birthdayMessage and sendingReport structs already have implementations of the getMessage method. The getMessage method returns a string, and any type that implements the method can be considered a message (meaning it implements the message interface).

Add the getMessage() method signature as a requirement on the message interface.
Complete the sendMessage function. It should return:
The content of the message.
The cost of the message, which is the length of the message multiplied by 3.
Notice that your code doesn't care at all about whether a specific message is a birthdayMessage or a sendingReport!

## Interface Implementation. Employee interface
At Textio we have full-time employees and contract employees. We have been tasked with making a more general employee interface so that dealing with different employee types is simpler.

Run the code. You should see an error indicating the contractor type does not fulfill the employee interface.
Implement the missing getSalary method for the contractor type so that it fulfills the employee interface.
A contractor's salary is their hourly pay multiplied by how many hours they work per year.

## Multiple Interfaces
A type can implement any number of interfaces in Go. For example, the empty interface, interface{}, is always implemented by every type because it has no requirements.

Assignment
Complete the required methods so that the email type implements both the expense and formatter interfaces.

Complete the cost() method:

If the email is not "subscribed", then the cost is 5 cents for each character in the body.
If it is, then the cost is 2 cents per character.
Return the total cost of the entire email in cents.
Complete the format() method.

It should return a string in this format:

'CONTENT' | Subscribed

If the email is not subscribed, change the second part to "Not Subscribed":

'CONTENT' | Not Subscribed

The single quotes are included in the string, and CONTENT is the email's body. For example:

'Hello, World!' | Subscribed

## Type Assertion in Go. Get expenses report
Implement the getExpenseReport function.

If the expense is an email, return the email's toAddress and the cost of the email.
If the expense is an sms, return the sms's toPhoneNumber and its cost.
If the expense has any other underlying type, return an empty string and 0.0 for the cost.

## Type switches. Get expenses report
After submitting our last snippet of code for review, a more experienced gopher told us to use a type switch instead of successive assertions. Let's make that improvement!

Implement the getExpenseReport function using a type switch.

If the expense is an email, return the email's toAddress and the cost of the email.
If the expense is an sms, return the sms's toPhoneNumber and its cost.
If the expense has any other underlying type, return an empty string and 0.0 for the cost.

## Message Formatter
As Textio evolves, the team has decided to introduce a new feature for custom message formats. Depending on the user's preferences, messages can be sent in different formats, such as plain text, markdown, code, or even encrypted text. To efficiently manage this, you'll implement a system using interfaces.

Assignment
Implement the formatter interface with a method format that returns a formatted string.
Define structs that satisfy the formatter interface: plainText, bold, code.
The structs must all have a message field of type string
plainText should return the message as is.
bold should wrap the message in two asterisks (**) to simulate bold text (e.g., **message**).
code should wrap the message in a single backtick (`) to simulate inline code (e.g., `message`)

## Process Notifications
Textio now has a system to process different types of notifications: direct messages, group messages, and system alerts. Each notification type has a unique way of calculating its importance score based on user interaction and content.

Assignment
Implement the importance methods for each message type. They should return the importance score for each message type.
For a directMessage the importance score is based on if the message isUrgent or not. If it is urgent, the importance score is 50 otherwise the importance score is equal to the DM's priorityLevel.
For a groupMessage the importance score is based on the message's priorityLevel
All systemAlert messages should return a 100 importance score.
Complete the processNotification function. It should identify the type and return different values for each type
For a directMessage, return the sender's username and importance score.
For a groupMessage, return the group's name and the importance score.
For a systemAlert, return the alert code and the importance score.
If the notification does not match any known type, return an empty string and a score of 0.