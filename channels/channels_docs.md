# CHANNELS

## What Is Concurrency?
Concurrency is the ability to make progress on multiple tasks without waiting for each one to finish before starting the next. Typically, our code is executed one line at a time, one after the other. This is called sequential execution or synchronous execution.

If the computer we're running our code on has multiple cores, it can run multiple tasks at exactly the same time. That's called parallelism. On a single core, it switches between tasks very quickly instead. Either way, the code we write looks the same in Go and takes advantage of whatever resources are available.
How Does Concurrency Work in Go?
Go was designed to be concurrent, which is a trait fairly unique to Go. It excels at coordinating many tasks safely using a simple syntax.
There isn't a popular programming language in existence where spawning concurrent execution is quite as elegant, at least in my opinion.
Concurrency is as simple as using the go keyword when calling a function:
go doSomething()
In the example above, doSomething() will be executed concurrently with the rest of the code in the function. The go keyword is used to spawn a new goroutine.

## Channels I
Channels are a typed, thread-safe queue. Channels allow different goroutines to communicate with each other.
Create a Channel
Like maps and slices, channels must be created before use. They also use the same makekeyword:
ch := make(chan int)
Send Data to a Channel
ch <- 69
The <- operator is called the channel operator. Data flows in the direction of the arrow. This operation will block until another goroutine is ready to receive the value.
Receive Data from a Channel
v := <-ch
This reads and removes a value from the channel and saves it into the variable v. This operation will block until there is a value in the channel to be read. Channels are First-In-First-Out (FIFO), meaning values are received in the same order they are sent.
Reference Type
Like maps and slices, channels are reference types. Changes made inside a function affect the original.
func send(ch chan int) {
    ch <- 99
}

func main() {
    ch := make(chan int)
    go send(ch)
    fmt.Println(<-ch) // 99
}
Blocking and Deadlocks
A deadlock is when a group of goroutines are all blocking so none of them can continue. This is a common bug that you need to watch out for in concurrent programming.

## Channels II
In the previous lesson, we saw how we can receive values from channels like this:
v := <-ch
However, sometimes we don't care what is passed through a channel. We only care whenand if something is passed. In that situation, we can block and wait until something is sent on a channel using the following syntax.
<-ch
This will block until it pops a single item off the channel, then continue, discarding the item.
In cases like this, empty structs are often used as a unary value so that the sender communicates that this is only a "signal" and not data that is meant to be used.
Here is an example:
func downloadData() chan struct{} {
        downloadDoneCh := make(chan struct{})

        go func() {
                fmt.Println("Downloading data file...")
                time.Sleep(2 * time.Second) // simulate download time

                // after the download is done, send a "signal" to the channel
                downloadDoneCh <- struct{}{}
        }()

        return downloadDoneCh
}

func processData(downloadDoneCh chan struct{}) {
        // any code here can run normally
        fmt.Println("Preparing to process data...")

        // block until `downloadData` sends the signal that it's done
        <-downloadDoneCh

        // any code here can assume that data download is complete
        fmt.Println("Data download complete, starting data processing...")
}

processData(downloadData())
// Preparing to process data...
// Downloading data file...
// Data download complete, starting data processing...

## Closing Channels in Go
Channels can be explicitly closed by a sender:
ch := make(chan int)

// do some stuff with the channel

close(ch)
Checking If a Channel Is Closed
Similar to the ok value when accessing data in a map, receivers can check the ok value when receiving from a channel to test if a channel was closed.
v, ok := <-ch
ok is false if the channel is empty and closed.
Don't Send on a Closed Channel
Sending on a closed channel will cause a panic. A panic on the main goroutine will cause the entire program to crash, and a panic in any other goroutine will cause that goroutine to crash.
Closing isn't necessary. There's nothing wrong with leaving channels open, they'll still be garbage collected if they're unused. You should close channels to indicate explicitly to a receiver that nothing else is going to come across.

## Range
Similar to slices and maps, channels can be ranged over.
for item := range ch {
    // item is the next value received from the channel
}
This example will receive values over the channel (blocking at each iteration if nothing new is there) and will exit only when the channel is closed.

## Select
Sometimes we have a single goroutine listening to multiple channels and want to process data in the order it comes through each channel.
A select statement is used to listen to multiple channels at the same time. It is similar to a switch statement but for channels.
select {
case i, ok := <-chInts:
        if ok {
                fmt.Println(i)
        }
case s, ok := <-chStrings:
        if ok {
                fmt.Println(s)
        }
}
The first channel with a value ready to be received will fire and its body will execute. If multiple channels are ready at the same time then one is chosen randomly. In the example above, the ok variable is a boolean. It is true while the channel is open, and false when the channel is closed by the sender.

## Select Default Case
The default case in a select statement executes immediately if no other channel has a value ready. A default case stops the select statement from blocking.
select {
case v := <-ch:
    // use v
default:
    // receiving from ch would block
    // so do something else
}
Ignoring Channels
Sometimes you want to ignore a channel's value. You can do this by not binding it to a variable:
select {
case <-ch:
    // event received; value ignored
default:
    // so do something else
}
Alternatively, you can use the blank identifier _ to ignore the value:
select {
case _ = <-ch:
    // event received; value ignored
default:
    // so do something else
}
Tickers
    • time.Tick() is a standard library function that returns a channel that sends a value on a given interval.
    • time.After() sends a value once after the duration has passed.
    • time.Sleep() blocks the current goroutine for the specified duration of time.
The functions take a time.Duration as an argument. For example:
time.Tick(500 * time.Millisecond)
If you don't add time.Millisecond (or another unit), it will default to nanoseconds. That's – taking a wild guess here – probably faster than you want it to be.
Read-Only Channels
A channel can be marked as read-only by casting it from a chan to a <-chan type. For example:
func main() {
    ch := make(chan int)
    readCh(ch)
}

func readCh(ch <-chan int) {
    // ch can only be read from
    // in this function
}
Write-Only Channels
The same goes for write-only channels, but the arrow's position moves.
func writeCh(ch chan<- int) {
    // ch can only be written to
    // in this function
}


## Channels Review
Here are a few extra things you should understand about channels from Dave Cheney's awesome article.
A Declared but Uninitialized Channel Is Nil Just Like a Slice
var s []int       // s is nil
var c chan string // c is nil

var s = make([]int, 5) // s is initialized and not nil
var c = make(chan int) // c is initialized and not nil
A Send to a Nil Channel Blocks Forever
var c chan string        // c is nil
c <- "let's get started" // blocks
A Receive from a Nil Channel Blocks Forever
var c chan string // c is nil
fmt.Println(<-c)  // blocks
A Send to a Closed Channel Panics
var c = make(chan int, 100)
close(c)
c <- 1 // panic: send on closed channel
A Receive from a Closed Channel Returns the Zero Value
var c = make(chan int, 100)
close(c)
fmt.Println(<-c) // 0
Any buffered values are received first. Once the buffer is empty, receives return the zero value.

## Exercise: Ping Pong
Many tech companies play ping pong in the office during break time. At Textio, we play virtual ping pong using channels.
Assignment
Run the code as-is. You should see that it doesn't do anything interesting: no ping or pong messages are printed.
Fix the bug in the pingPong function.
Remember: if a program exits before its goroutines have completed, those goroutines will be killed silently. Which of the function calls probably shouldn't run in the background as a goroutine?