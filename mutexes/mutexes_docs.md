# MUTEXES

## Mutexes in Go
Mutexes allow us to lock access to data. This ensures that we can control which goroutines can access certain data at which time.
Go's standard library provides a built-in implementation of a mutex with the sync.Mutextype and its two methods:
    • .Lock()
    • .Unlock()
We can protect a block of code by surrounding it with a call to Lock and Unlock as shown on the protected() function below.
It's good practice to structure the protected code within a function so that defer can be used to ensure that we never forget to unlock the mutex.
func protected(){
    mu.Lock()
    defer mu.Unlock()
    // the rest of the function is protected
    // any other calls to `mu.Lock()` will block
}
In the example above, mu is a mutex variable. In your assignment, the mutex is a field on your safeCounter struct (sc.mu), so you'll call sc.mu.Lock() and use defer sc.mu.Unlock() to protect the critical sections.
Mutexes are powerful. Like most powerful things, they can also cause many bugs if used carelessly.
Maps Are Not Thread-Safe
Maps are not safe for concurrent use! If you have multiple goroutines accessing the same map, and at least one of them is writing to the map, you must lock your maps with a mutex.

## Why Is It Called a “mutex”?
Mutex is short for mutual exclusion, and the conventional name for the data structure that provides it is "mutex", often abbreviated to "mu".
It's called "mutual exclusion" because a mutex excludes different threads (or goroutines) from accessing the same data at the same time.

## Mutex Review
The principal problem that mutexes help us avoid is the concurrent read/write problem. This problem arises when one thread is writing to a variable while another thread is reading from that same variable at the same time.
When this happens, a Go program will panic because the reader could be reading bad data while it's being mutated in place.
Questions:
1. How many goroutines can hold the same mutex lock at once?
1
2. Why would you use a mutex?
To savely access shared resources concurrently

## RW Mutex
The standard library also exposes a sync.RWMutex
In addition to these methods:
    • Lock()
    • Unlock()
The sync.RWMutex also has these methods for concurrent reads:
    • RLock()
    • RUnlock()
The sync.RWMutex improves performance in read-intensive processes. Multiple goroutines can safely read from the map simultaneously, as many RLock() calls can occur at the same time. However, only one goroutine can hold a Lock(), and during this time, all RLock() operations are blocked.

## Read/Write Mutex Review
Maps are safe for concurrent read access, just not concurrent read/write or write/write access. A read/write mutex allows all the readers to access the map at the same time, but a writer will still lock out all other readers and writers.
Questions
1. How many writers can access a RWMutex at once?
1
2. How many readers can access a RWMutex at once?
Infinite
3. Can readers and writers lock a RWMutex at the same time?
No