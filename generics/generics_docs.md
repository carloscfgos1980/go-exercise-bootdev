# GENERICS

## Generics in Go
As we've mentioned, Go does not support classes. For a long time, that meant that Go code couldn't easily be reused in many cases. For example, imagine some code that splits a slice into 2 equal parts. The code that splits the slice doesn't care about the types of valuesstored in the slice. Before generics, we needed to write the same code for each type, which is a very un-DRY thing to do.
func splitIntSlice(s []int) ([]int, []int) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}
func splitStringSlice(s []string) ([]string, []string) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}
In Go 1.18 however, support for generics was released, effectively solving this problem!
Type Parameters
Put simply, generics allow us to use variables to refer to specific types. This is an amazing feature because it allows us to write abstract functions that drastically reduce code duplication.
func splitAnySlice[T any](s []T) ([]T, []T) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}
In the example above, T is the name of the type parameter for the splitAnySlicefunction, and we've said that it must match the any constraint, which means it can be anything. This makes sense because the body of the function doesn't care about the types of things stored in the slice.
firstInts, secondInts := splitAnySlice([]int{0, 1, 2, 3})
fmt.Println(firstInts, secondInts)

## Why Generics?
Generics Reduce Repetitive Code
You should care about generics because they mean you don't have to write as much code! It can be frustrating to write the same logic over and over again, just because you have some underlying data types that are slightly different.
Generics Are Used More Often in Libraries and Packages
Generics give Go developers an elegant way to write amazing utility packages. While you will see and use generics in application code, I think it will be much more common to see generics used in libraries and packages. Libraries and packages contain importable code intended to be used in many applications, so it makes sense to write them in a more abstract way. Generics are often the way to do just that!
Why Did It Take So Long to Get Generics?
Go places an emphasis on simplicity. In other words, Go has purposefully left out many features to provide its best feature: being simple and easy to work with.
According to historical data from Go surveys, Go's lack of generics has always been listed as one of the top three biggest issues with the language. At a certain point, the drawbacks associated with the lack of a feature like generics justify adding complexity to the language.
Question 1: Which code would generics be most likely to help with?
Answer: A binary tree
    • Why: A binary tree is a data structure whose internal logic (inserting nodes, traversing, balancing) does not depend on the specific type of data stored inside each node. Without generics, you would either have to write a separate IntBinaryTree, StringBinaryTree, and UserBinaryTree, or rely on interface{} / any with type assertions. With generics, you can write a single Tree[T] that works safely across any type.
    • Why not the others:
        ◦ Calculating the area of a circle always deals with numbers (specifically float64 or int), so there is no need to generalize across arbitrary types.
        ◦ Checking if a string contains a substring inherently operates only on string types.

Question 2: Go's approach to language design is...
Answer: Resist adding new features unless they're extremely important
    • Why: Go prioritizes simplicity, readability, and maintainability. The designers intentionally left out many features common in other languages (like inheritance, operator overloading, and for a long time, generics) to keep the language small and straightforward. Features are only introduced when the lack of them causes significant enough pain across the ecosystem to justify the added complexity.

Question 3: Generics will probably be used more heavily in...
Answer: Library packages
    • Why: Libraries and packages provide reusable utility code—such as custom data structures, collection manipulation tools, or algorithms—intended to be used across many different applications with various types. Application code (often in main packages) usually models concrete domain concepts (like User, Order, or Invoice) where the exact types are already known.

## Constraints
Sometimes you need your generic function to know something about the types it operates on. The example we used in the first exercise didn't need to know anything about the types in the slice, so we used the built-in any constraint:
func splitAnySlice[T any](s []T) ([]T, []T) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}
Constraints are just interfaces that allow us to write generics that only operate within the constraints of a given interface type. In the example above, the any constraint is the same as the empty interface because it means the type in question can be anything.
Creating a Custom Constraint
Let's take a look at the example of a concat function. It takes a slice of values and concatenates the values into a string. This should work with any type that can represent itself as a string, even if it's not a string under the hood. For example, a user struct can have a .String() that returns a string with the user's name and age.
type stringer interface {
    String() string
}

func concat[T stringer](vals []T) string {
    result := ""
    for _, val := range vals {
        // this is where the .String() method
        // is used. That's why we need a more specific
        // constraint instead of the any constraint
        result += val.String()
    }
    return result
}


## Parametric Constraints
Your interface definitions, which can later be used as constraints, can accept type parameters as well.
// The store interface represents a store that sells products.
// It takes a type parameter P that represents the type of products the store sells.
type store[P product] interface {
        Sell(P)
}

type product interface {
        Price() float64
        Name() string
}

type book struct {
        title  string
        author string
        price  float64
}

func (b book) Price() float64 {
        return b.price
}

func (b book) Name() string {
        return fmt.Sprintf("%s by %s", b.title, b.author)
}

type toy struct {
        name  string
        price float64
}

func (t toy) Price() float64 {
        return t.price
}

func (t toy) Name() string {
        return t.name
}

// The bookStore struct represents a store that sells books.
type bookStore struct {
        booksSold []book
}

// Sell adds a book to the bookStore's sold slice.
func (bs *bookStore) Sell(b book) {
        bs.booksSold = append(bs.booksSold, b)
}

// The toyStore struct represents a store that sells toys.
type toyStore struct {
        toysSold []toy
}

// Sell adds a toy to the toyStore's sold slice.
func (ts *toyStore) Sell(t toy) {
        ts.toysSold = append(ts.toysSold, t)
}

// sellProducts takes a store and a slice of products and sells
// each product one by one.
func sellProducts[P product](s store[P], products []P) {
        for _, p := range products {
                s.Sell(p)
        }
}

func main() {
        bs := bookStore{
                booksSold: []book{},
        }

    // By passing in "book" as a type parameter, we can use the sellProducts function to sell books in a bookStore
        sellProducts[book](&bs, []book{
                {
                        title:  "The Hobbit",
                        author: "J.R.R. Tolkien",
                        price:  10.0,
                },
                {
                        title:  "The Lord of the Rings",
                        author: "J.R.R. Tolkien",
                        price:  20.0,
                },
        })
        fmt.Println(bs.booksSold)

    // We can then do the same for toys
        ts := toyStore{
                toysSold: []toy{},
        }
        sellProducts[toy](&ts, []toy{
                {
                        name:  "LEGO bricks",
                        price: 10.0,
                },
                {
                        name:  "Barbie",
                        price: 20.0,
                },
        })
        fmt.Println(ts.toysSold)
}


## Naming Generic Types
Let's look at this simple example again:
func splitAnySlice[T any](s []T) ([]T, []T) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}
Remember, T is just a variable name, We could have named the type parameter anything. T happens to be a fairly common convention for a type variable, similar to how i is a convention for index variables in loops.
This is just as valid:
func splitAnySlice[MyAnyType any](s []MyAnyType) ([]MyAnyType, []MyAnyType) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}
