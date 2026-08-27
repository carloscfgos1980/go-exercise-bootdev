# Go concepts by boot.dev 
## CLI to run program and tests
// ce dirctory where is located the test
go test  -run '^TestMonthlyBillIncrease$' -v

// example to run function from root directory
go run . product-message enterprise



## Chacpters

1. Variables:
-  Short Delcaration
- The Compilation Process
- Go Program Structure
- Two Kinds of Errors
- Type
- Go Is Statically Typed
- Compiled vs. Interpreted
- Same Line Declarations

2. Constants
- Constant concept
- Computed Constants
- Formatting Strings in Go
- Runes and String Encoding
- Practicing formatting

3. Condicionals
- Conditionals
- The Initial Statement of an If Block
- Switch

4. Functions
- Functions concept
- Multiple Parameters
- Declaration Syntax
- Passing Variables by Value
- Ignoring Return Values
- Named Return Values
- The Benefits of Named Returns
- Explicit Returns
- Early Returns
- Functions As Values
- Anonymous Functions
- Exercise: Processing Orders
- Closures
- Currying

5. STRUCTS
- Structs in Go
- Nested Structs in Go
- Anonymous Structs in Go
- Embedded Structs
- Struct Methods in Go
- Empty Struct Memory
- Exercise: Update Users
- Exercise: Send Message

6. Interface
- Interfaces in Go
- Interface Implementation
- Interfaces Are Implemented Implicitly
- Exercise: Multiple Interfaces
- Name Your Interface Parameters
- Type Assertions in Go
- Type Switches
- Clean Interfaces
- Exercise: Message Formatter
- Exercise: Process Notifications

7. Errors
- The Error Interface
- Formatting Strings Review
- Custom Error interface's
- Errors Quiz
- The Errors Package
- Panic
- Exercise: User Input

8. Loops
- Loops in Go
- Omitting Conditions from a for Loop in Go
- There Is No While Loop in Go
- Exercise: Fizzbuzz
- Continue & Break (Prime numbers)
- exercise: Connections

9. Slices
- Arrays in Go
- Slices in Go
- Slices Review
- Make
- Len and Cap Review
- Variadic
- Append
- Range
- Slice of Slices
- Tricky Slices
- Exercise: Message Filter
- Exercise: Password Strength
- Exercise: Message Tagger

10. Maps
- Maps
- Mutations
- Key Types
- Exercise: Count Instances
- Effective Go
- Nested
- Exercise: Distinct Words

11. POINTERS
- Introduction to Pointers
- References
- Pass by Reference
- Pointers Quiz
- Nil Pointers
- Pointer Receivers
- Pointer Receiver Code
- Pointer Performance
- Exercise: Update Balance

12. PACKAGES
- Packages
- Package Naming
- Modules
- Go Run
- Go Build
- Go Install
- Custom Package
- Custom Package Continued
- Remote Packages
- Clean Packages

13. CHANNELS
- What Is Concurrency?
- Channels I
- Channels II
- Closing Channels in Go
- Range
- Select
- Channels Review
- Exercise: Ping Pong

14. MUTEXES
- Mutexes in Go
- Why Is It Called a “mutex”?
- Mutex Review
- RW Mutex
- Read/Write Mutex Review

15. GENERICS
- Generics in Go
- Why Generics?
- Constraints
- Parametric Constraints
- Naming Generic Types

16. ENUMS
- Lack of Enums
- Type Definitions e
- Iota