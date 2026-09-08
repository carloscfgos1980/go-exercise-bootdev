# Pointers assigments

## Reference. Remove profanity


Complete the removeProfanity function.

It should use the strings.ReplaceAll function to replace all instances of the following words in the input message with asterisks.

Word	Replacement
fubb	****
shiz	****
witch	*****
It should update the value in the pointer and return nothing.

Do not alter the function signature.

## Pass by reference. Analyze message

Write an analyzeMessage function. It should accept a pointer to an Analytics struct and a Message struct (not a pointer).

It should look at the Success field of the Message struct and, based on that, increment the MessagesTotal, MessagesSucceeded, or MessagesFailed fields of the Analytics struct as appropriate.

## nil pointers. Remove profanity

Let's make our profanity checker safe. Update the removeProfanity function. If message is nil, return early to avoid a panic. After all, there are no bad words to remove.

## Pointer receiber code. Set Message

Fix the bug in the code so that setMessage sets the message field of the given email struct, and the new value persists outside the scope of the setMessage method.

## Update Balance
Textio needs a new way to update user's account balance.

Assignment
Implement the updateBalance function. It should take a customer pointer and a transaction, and return an error. Depending on the transactionType, it should either add or subtract the amount from the customer's balance. If the customer does not have enough money, it should return the error insufficient funds and leave the balance unchanged. If the transactionType isn't a withdrawal or deposit, it should return the error unknown transaction type. Otherwise, it should process the transaction and return nil.

alice := customer{id: 1, balance: 100.0}
deposit := transaction{customerID: 1, amount: 50, transactionType: transactionDeposit}

updateBalance(&alice, deposit)
// id: 1 balance: 150