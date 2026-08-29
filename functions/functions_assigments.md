# ASSIGMENTS

1. Functionns
We often will need to manipulate strings in our messaging app. For example, adding some personalization by using a customer's name within a template. The concat function should take two strings and smash them together.

hello + world = helloworld
Fix the function signature of concat to reflect its behavior.

2. Unit testing
Complete the getMonthlyPrice function. It accepts a tier (string) as input and returns the monthly price for that tier in pennies. Here are the prices in dollars:

"basic" - $100.00
"premium" - $150.00
"enterprise" - $500.00
Convert the prices from dollars to pennies. If the given tier doesn't match any of the above, return 0 pennies.

3. Passing Variables by Value
monthlyBillIncrease: Should return the increase in the bill from the previous to the current month. If the bill decreased, return a negative number.
getBillForMonth: Should return the total cost for the number of messages sent.
Fix the bugs in the monthlyBillIncrease and getBillForMonth functions. Looks like whoever wrote the functions didn't know the getBillForMonth function's bill parameter would be passed by value. It's not actually updating the lastMonthBill and thisMonthBill variables as intended so monthlyBillIncrease isn't returning the right result.

Drop the bill parameter from getBillForMonth, so it only takes 2 parameters.
Instead, simply return the total cost of the messages.
monthlyBillIncrease should use the result of calling getBillForMonth to calculate the increase between months.

4. Ignoring Return Values
Run the code as-is. You should get a compiler error.
Fix getProductMessage to ignore the unused return value.

5. Named Return Values
One of our clients likes us to send text messages reminding users of life events coming up.

Fix the bug by adding named return values to the function signature – the bare return at the end is already a naked return that will return them. The variables need to be automatically initialized. Order them as they appear in the code. Do not alter the body of the function.

6. Explicit Returns
Fix the bug in the code so that it returns the named values explicitly.

7. Functions As Values
Complete the reformat function. It takes a message string and a formatter function as input:

Apply the given formatter three times to the message
Add a prefix of TEXTIO: to the result
Return the final string
For example, if the message is "General Kenobi" and the given formatter adds a period to the end of the string, the final result should be

TEXTIO: General Kenobi...

8. Anonymous Functions
Complete the printReports function. It takes as input a sequence of messages, intro, body, outro. It should call printCostReport once for each message by making three separate calls (you don't need to use loops or arrays for this).

For each call of printCostReport, give it an anonymous function that returns the cost of a message as an integer. Here are the costs:

Intro: 2x the message length
Body: 3x the message length
Outro: 4x the message length
Use the built-in len() function to get the length of a string:

helloLen := len("hello")
// helloLen = 5

9. Defer
Complete the bootup function.

Be sure to print the following string just before the bootup function returns:

TEXTIO BOOTUP DONE

Use defer so that you only have to write this message once instead of before each return statement. The message should be printed on its own newline.

10. Block Scope
Run the code without changing anything: you should see a compilation error.
Fix the scoping issue in the function so that it runs as you'd expect.

11. Processing Orders
Management thinks our branding is so creative that our SaaS customers will pay for Textio merch.

Assignment
Complete the placeOrder function.

It returns a bool indicating whether the order was successful (true is a success) and a float64 representing the user's balance after the order. The placeOrder function should always return the account balance regardless of whether it was adjusted.

The amountInStock and calcPrice functions can be used to look up the current stock and price of an item.

If the quantity is greater than the amount in stock, the order should be rejected.
If the user doesn't have enough money in their account, the order should be rejected.
Otherwise, the order should be accepted and you should return the new balance.

12. Closures
Keeping track of how many texts we send is mission-critical at Textio. Complete the adder() enclosing function.

Create an enclosed sum value inside the adder() function.
Return a function from the adder() function that adds its input (an int) to the sum and returns the new value of sum. (In other words, it keeps a running total of the sum variable within a closure.)

13. Currying
The Textio API needs a very robust error-logging system so we can see when things are going awry in the back-end system. We need a function that can create a custom "logger" (a function that prints to the console) given a specific formatter.

These errors are test data, not runtime failures.
Complete the getLogger function. It should take as input a formatter function and return a new function. The new logger function takes as input two strings and passes them to the formatter, then prints the result. Keep the order of the strings.