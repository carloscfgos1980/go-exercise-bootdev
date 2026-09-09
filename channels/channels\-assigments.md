# Channels assigments

## Concurrency. Send email

At Textio we send a lot of network requests. Each email we send must go out over the internet. To serve our millions of customers, we need a single Go program to be capable of sending thousands of emails at once.

Edit the sendEmail() function to execute its anonymous function concurrently so that the "received" message prints after the "sent" message.

## Channels. Checking email age

Run the program. You'll see that it deadlocks and never exits. The sendIsOld function is trying to send on a channel, but no other goroutines are running that can accept the value from the channel.

Fix the deadlock by spawning a goroutine to send the "is old" values.

## Channels Ii. Number of databases

Our Textio server isn't able to boot up until it receives the signal that its databases are all online, and it learns about them being online by waiting for tokens (empty structs) on a channel.

Run the code. It never exits! The channel passed to waitForDBs stays blocked, because it's only popping the first value off the channel.

Fix the waitForDBs function. It should pause execution until it receives a token for each of the numDBs databases from the dbChan channel. Each time waitForDBs reads a token, the getDBsChannel goroutine will print a message to the console for you. Look at the test code to see how these functions are used so you can understand the control flow.

## buffered channels

We want to be able to send emails in batches. A writing goroutine will write an entire batch of email messages to a buffered channel, and later, once the channel is full, a reading goroutine will read all of the messages from the channel and send them out to our clients.

Complete the addEmailsToQueue function. It should create a buffered channel with a buffer large enough to store all of the emails it's given. It should then write the emails to the channel in order, and finally return the channel.

## closing channels. Count report

At Textio we're all about keeping track of what our systems are up to with great logging and telemetry.

The sendReports function sends out a batch of reports to our clients and reports back how many were sent across a channel. It closes the channel when it's done.

Complete the countReports function. It should:

Use an infinite for loop to read from the channel:
If the channel is closed, break out of the loop
Otherwise, keep a running total of the number of reports sent
Return the total number of reports sent

## Range. Fibonacci

At Textio we're all about keeping track of what our systems are up to with great logging and telemetry.

The sendReports function sends out a batch of reports to our clients and reports back how many were sent across a channel. It closes the channel when it's done.

Complete the countReports function. It should:

Use an infinite for loop to read from the channel:
If the channel is closed, break out of the loop
Otherwise, keep a running total of the number of reports sent
Return the total number of reports sent

## Select. Log Messages


Complete the logMessages function. It should:

Use an infinite for loop.
Use a select statement to read from chEmails and chSms in the order messages arrive.
Log messages with logEmail and logSms.
return when one of the two channels closes, whichever is first.

## Select default case. Save backups

Like all good back-end engineers, we frequently save backup snapshots of the Textio database.

Complete the saveBackups function.

It should use a loop to read values from the snapshotTicker and saveAfter channels simultaneously and continuously.

If a value is received from snapshotTicker, call takeSnapshot()
If a value is received from saveAfter, call saveSnapshot() and return from the function: you're done.
If neither channel has a value ready, call waitForData() and then time.Sleep() for 500 milliseconds. After all, we want to show in the logs that the snapshot service is running.

## Ping Pong
Many tech companies play ping pong in the office during break time. At Textio, we play virtual ping pong using channels.

Assignment
Run the code as-is. You should see that it doesn't do anything interesting: no ping or pong messages are printed.

Fix the bug in the pingPong function.

Remember: if a program exits before its goroutines have completed, those goroutines will be killed silently. Which of the function calls probably shouldn't run in the background as a goroutine?