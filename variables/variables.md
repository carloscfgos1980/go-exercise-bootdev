# Variables

- Short Delcaration
Use the Walrus Operator
The :=, (walrus operator) should be used instead of var style declarations basically anywhere possible. The limitation is that := can't be used outside of a function

- The Compilation Process
Computers need machine code, they don't understand English or even Go. We need to convert our high-level (Go) code into machine language, which is really just a set of instructions that some specific hardware can understand. In your case, your CPU.
The Go compiler's job is to take Go code and produce machine code, an .exe file on Windows or a standard executable on Mac/Linux.

- Go Program Structure
We'll go over all of this later in more detail, but to sate your curiosity:
    1. package main lets the Go compiler know that we want this code to compile and run as a standalone program, as opposed to being a library that's imported by other programs.
    2. import "fmt" imports the fmt (formatting) package from the standard library. It allows us to use fmt.Println to print to the console.
    3. func main() defines the main function, the entry point for a Go program.

- Two Kinds of Errors
Generally speaking, there are two kinds of errors in programming:
    1. Compilation errors. Occur when code is compiled. It's generally better to have compilation errors because they'll never accidentally make it into production. You can't ship a program with a compiler error because the resulting executable won't even be created.
    2. Runtime errors. Occur when a program is running. These are generally worse because they can cause your program to crash or behave unexpectedly.ompiled
Generally speaking, languages that compile directly to machine code produce programs that are faster than interpreted programs.

Go is one of the fastest programming languages, beating JavaScript, Python, and Ruby handily in most benchmarks.

However, Go programs don't run quite as fast as its compiled Rust, Zig, and C counterparts. That said, it compiles much faster than they do, which makes the developer experience super productive. Unfortunately, there are no swordfights on Go teams…

Blank 1: "Go code generally runs faster than interpreted languages"
The lesson explains that compiled languages produce programs that run faster than interpreted ones. This is because:
    • Compiled languages (like Go) translate your source code directly into machine code before you run it. When you execute the program, the CPU is reading instructions it can act on immediately.
    • Interpreted languages (like Python, Ruby, or JavaScript in many cases) translate code into machine instructions while the program is running, line by line. That translation step happens over and over as the program executes, adding overhead.

The lesson even says Go beats JavaScript, Python, and Ruby handily in most benchmarks - all three of those are commonly interpreted (or JIT-compiled at runtime), so this confirms "faster" for the first blank.

Blank 2: "and compiles faster than other compiled languages"
The lesson specifically calls out that Go doesn't run quite as fast as Rust, Zig, or C (which are also compiled), but it compiles much faster than they do. Compilation speed is a separate thing from runtime speed:
    • Runtime speed = how fast the final program executes.
    • Compile speed = how long it takes to turn your source code into that final program.

Go trades a small amount of runtime performance (compared to Rust/C/Zig) for a much snappier compile step, which is why the lesson says it makes for a productive developer experience - you can compile and rerun your program quickly while working on it.

So both blanks are "faster": Go is faster than interpreted languages at runtime, and faster than other compiled languages at compiling.

- Type Sizes
Integers, uints, floats, and complex numbers all have type sizes.
    • Signed integers (no decimal)
int  int8  int16  int32  int64
    • Unsigned integers (non-negative numbers/no decimal)
uint uint8 uint16 uint32 uint64 uintptr
    • Signed decimal numbers
float32 float64
    • Complex numbers (a complex number has a real and imaginary part)
complex64 complex128

What's the Deal With the Sizes?
The size (8, 16, 32, 64, 128, etc.) represents how many bits in memory will be used to store the variable. The "default" int and uint types refer to their respective 32 or 64-bit sizes depending on the environment of the user.

The "standard" types that should be used unless you have a specific performance need (e.g. using less memory) are:
    • int
    • uint
    • float64
    • complex128
Converting Between Types
Some types can be easily converted like this:
temperatureFloat := 88.26
temperatureInt := int(temperatureFloat)

- Go Is Statically Typed

Go enforces static typing meaning variable types are known before the code runs. That means your editor and the compiler can display type errors before the code is ever run, making development easier and faster.

Contrast this with most dynamically typed languages like JavaScript and Python... Dynamic typing often leads to subtle bugs that are hard to detect. The code must be run to catch syntax and type errors. (sometimes in production if you're unlucky 😨)

Languages also have strong or weak typing, meaning stricter or weaker type checking rules.

Concatenating Strings
Two strings can be concatenated with the + operator. But the compiler will not allow you to concatenate a string variable with an int or a float64.

- Compiled vs. Interpreted
You can run a compiled program without the original source code. You don't need the compiler anymore after it's done its job. That's how most video games are distributed! Players don't need to install the correct version of Go to run a PC game: they just download the executable game and run it.

With interpreted languages like Python and Ruby, the code is interpreted at runtime by a separate program known as the "interpreter". Distributing code for users to run can be a pain because they need to have an interpreter installed, and they need access to the source code.

Examples of Compiled Languages
Go, C, C++, Rust

Examples of Interpreted Languages
JavaScript (sometimes JIT-compiled, but a similar concept), Python, Ruby

Why Build Textio in a Compiled Language?
One of the most convenient things about using a compiled language like Go for Textio is that when we deploy our server we don't need to include any runtime language dependencies like Node or a Python interpreter. We just add the pre-compiled binary to the server and start it up!

- Same Line Declarations
You can declare multiple variables on the same line:
mileage, company := 80276, "Toyota"
The above is the same as:
mileage := 80276
company := "Toyota"

- Small Memory Footprint
Go programs are fairly lightweight. Each program includes a small amount of extra code that's included in the executable binary called the Go Runtime. One of the purposes of the Go runtime is to clean up unused memory at runtime. It includes a garbage collector that automatically frees up memory that's no longer in use.

Comparison
As a general rule, Java programs use more memory than comparable Go programs. There are several reasons for this, but one of them is that Java uses a virtual machine to interpret bytecode at runtime and typically allocates more on the heap.

On the other hand, Rust and C programs use slightly less memory than Go programs because more control is given to the developer to optimize the memory usage of the program. The Go runtime just handles it for us automatically.

Question 1: Generally speaking, which language uses the most memory?
Answer: Java
Why?
Java code runs on top of the Java Virtual Machine (JVM). The JVM manages execution, interpret bytecode, and handles heap allocations, which adds significant baseline memory overhead even for very simple programs.
In contrast:
    • C and Rust compile directly to machine code and give developers manual/strict control over every byte of memory, making their footprint tiny.
    • Go also compiles directly to machine code and includes a small runtime, making it far lighter than Java, though slightly larger than C or Rust.

Question 2: What's one of the purposes of the Go runtime?
Answer: To clean up unused memory
Why?
Go includes a small built-in library shipped inside every compiled binary called the Go Runtime. One of its primary responsibilities is running a garbage collector.
The garbage collector automatically tracks variables and objects in memory. When your program no longer needs or uses them, the garbage collector frees up that memory so it can be reused, saving you from having to manually manage memory like in C or C++.