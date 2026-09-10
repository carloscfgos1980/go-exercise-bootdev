# Mutexes Assigments

## Mutexes in Go. Save Count Increment and Value

We send emails across many different goroutines at Textio. To keep track of how many we've sent to a given email address, we use an in-memory map.

Our safeCounter struct is unsafe! Update the inc() and val() methods so that they utilize the safeCounter's mutex to ensure that the map is not accessed by multiple goroutines at the same time.

Note: Wasm Is Single-Threaded
Now, it's worth pointing out that our execution engine on Boot.dev uses web assembly to run the code you write in your browser. Web assembly is single-threaded, which awkwardly means that maps are thread-safe in web assembly. I've simulated a multi-threaded environment with the slowIncrement and slowVal functions.

In reality, any Go code you write may or may not run on a single-core machine, so it's always best to write your code so that it is safe no matter which hardware it runs on.

## RW Mutex

Let's update our same code from the last assignment, but this time we can speed it up by allowing readers to read from the map concurrently.

Run the new test suite. You'll notice that it hangs forever and you'll need to cancel it.

Update the val method to only lock the mutex for reading.