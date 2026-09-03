# Interface

## Interfaces in Go
Interfaces allow you to focus on what a type does rather than how it's built. They can help you write more flexible and reusable code by defining behaviors (like methods) that different types can share. This makes it easy to swap out or update parts of your code without changing everything else.
Interfaces are just collections of method signatures. A type "implements" an interface if it has methods that match the interface's method signatures.
In the following example, a "shape" must be able to return its area and perimeter. Both rectand circle fulfill the interface.
type shape interface {
  area() float64
  perimeter() float64
}

type rect struct {
    width, height float64
}
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perimeter() float64 {
    return 2*r.width + 2*r.height
}

type circle struct {
    radius float64
}
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
    return 2 * math.Pi * c.radius
}
When a type implements an interface, it can then be used as that interface type.
func printShapeData(s shape) {
        fmt.Printf("Area: %v - Perimeter: %v\n", s.area(), s.perimeter())
}
Because we say the input is of type shape, we know that any argument must implement the .area() and .perimeter() methods.
As an example, because the empty interface doesn't require a type to have any methods at all, every type automatically implements the empty interface, written as:
interface{}

## Interface Implementation
Interfaces are implemented implicitly.
A type never declares that it implements a given interface. If an interface exists and a type has the proper methods defined, then the type automatically fulfills that interface.
A quick way of checking whether a struct implements an interface is to declare a function that takes an interface as an argument. If the function can take the struct as an argument, then the struct implements the interface.

## Interfaces Are Implemented Implicitly
A type implements an interface by implementing its methods. Unlike in many other languages, there is no explicit declaration of intent, there is no "implements" keyword.
Implicit interfaces decouple the definition of an interface from its implementation. You may add methods to a type and in the process be unknowingly implementing various interfaces, and that's okay.

## Exercise: Multiple Interfaces
A type can implement any number of interfaces in Go. For example, the empty interface, interface{}, is always implemented by every type because it has no requirements.
Assignment
Complete the required methods so that the email type implements both the expenseand formatter interfaces.
Complete the cost() method:
    1. If the email is not "subscribed", then the cost is 5 cents for each character in the body.
    2. If it is, then the cost is 2 cents per character.
    3. Return the total cost of the entire email in cents.
Complete the format() method.
    1. It should return a string in this format:
'CONTENT' | Subscribed
    2. If the email is not subscribed, change the second part to "Not Subscribed":
'CONTENT' | Not Subscribed
The single quotes are included in the string, and CONTENT is the email's body. For example:
'Hello, World!' | Subscribed

## Name Your Interface Parameters
Consider the following interface:
type Copier interface {
  Copy(string, string) int
}
This is a valid interface, but based on the code alone, can you deduce what kinds of strings you should pass into the Copy function?
We know the function signature expects 2 string types, but what are they? Filenames? URLs? Raw string data? For that matter, what the heck is that int that's being returned?
Let's add some named parameters and return data to make it more clear.
type Copier interface {
  Copy(sourceFile string, destinationFile string) (bytesCopied int)
}
Much better. We can see what the expectations are now. The first parameter is the sourceFile, the second parameter is the destinationFile, and bytesCopied, an integer, is returned.

## Type Assertions in Go
When working with interfaces in Go, every once-in-awhile you'll need access to the underlying type of an interface value. You can cast an interface to its underlying type using a type assertion.
The example below shows how to safely access the radius field of s when it is not known that s is a circle:
type shape interface {
        area() float64
}

type circle struct {
        radius float64
}

func (c circle) area() float64 {
        // ...
}

func printShapeInfo(s shape) {
        c, ok := s.(circle)
        if ok {
                radius := c.radius
                fmt.Printf("s is a circle, radius: %v\n", radius)
                return
        }
}
In printShapeInfo:
    • We want to check if s is a circle in order to cast it into its underlying concrete type
    • We know (from the function signature) that s is an instance of the shape interface, but we do not know if it's also a circle
    • c is a new circle struct cast from s
    • ok is true if s is indeed a circle, or false if s is NOT a circle
The two values aren't duplicates: the assertion returns the concrete circle value and a separate boolean that says whether the assertion succeeded.

## Type Switches
A type switch makes it easy to do several type assertions in a series.
A type switch is similar to a regular switch statement, but the cases specify types instead of values.
func printNumericValue(num interface{}) {
        switch v := num.(type) {
        case int:
                fmt.Printf("%T\n", v)
        case string:
                fmt.Printf("%T\n", v)
        default:
                fmt.Printf("%T\n", v)
        }
}

func main() {
        printNumericValue(1)
        // prints "int"

        printNumericValue("1")
        // prints "string"

        printNumericValue(struct{}{})
        // prints "struct {}"
}
fmt.Printf("%T\n", v) prints the type of a variable.

- Clean Interfaces
Writing clean interfaces is hard. Frankly, any time you're dealing with abstractions in code, the simple can become complex very quickly if you're not careful. Let's go over some rules of thumb for keeping interfaces clean.
1. Keep Interfaces Small
If there is only one piece of advice that you take away from this lesson, make it this: keep interfaces small! Interfaces are meant to define the minimal behavior necessary to accurately represent an idea or concept.
Here is an example from the standard HTTP package of a larger interface that's a good example of defining minimal behavior:
type File interface {
    io.Closer
    io.Reader
    io.Seeker
    Readdir(count int) ([]os.FileInfo, error)
    Stat() (os.FileInfo, error)
}
Any type that satisfies the interface's behaviors can be considered by the HTTP package as a File. This is convenient because the HTTP package doesn't need to know if it's dealing with a file on disk, a network buffer, or a simple []byte.
2. Interfaces Should Have No Knowledge of Satisfying Types
An interface should define what is necessary for other types to classify as a member of that interface. They shouldn't be aware of any types that happen to satisfy the interface at design time.
For example, let's assume we are building an interface to describe the components necessary to define a car.
type car interface {
        Color() string
        Speed() int
        IsFiretruck() bool
}
Color() and Speed() make perfect sense. They are methods confined to the scope of a car. IsFiretruck() is an anti-pattern. We are forcing all cars to declare whether or not they are firetrucks. In order for this pattern to make any amount of sense, we would need a whole list of possible subtypes. IsPickup(), IsSedan(), IsTank()... where does it end??
Instead, the developer should have relied on the native functionality of type assertion to derive the underlying type when given an instance of the car interface. Or, if a sub-interface is needed, it can be defined as:
type firetruck interface {
        car
        HoseLength() int
}
Which inherits the required methods from car as an embedded interface and adds one additional required method to make the car a firetruck.
3. Interfaces Are Not Classes
    • Interfaces are not classes, they are slimmer.
    • Interfaces don't have constructors or destructors that require that data is created or destroyed.
    • Interfaces aren't hierarchical by nature, though there is syntactic sugar to create interfaces that happen to be supersets of other interfaces.
    • Interfaces define function signatures, but not underlying behavior. Making an interface often won't DRY up your code in regard to struct methods. For example, if five types satisfy the fmt.Stringer interface, they all need their own version of the String()function.

## Exercise: Message Formatter
As Textio evolves, the team has decided to introduce a new feature for custom message formats. Depending on the user's preferences, messages can be sent in different formats, such as plain text, markdown, code, or even encrypted text. To efficiently manage this, you'll implement a system using interfaces.
Assignment
    1. Implement the formatter interface with a method format that returns a formatted string.
    2. Define structs that satisfy the formatter interface: plainText, bold, code.
        ◦ The structs must all have a message field of type string
    • plainText should return the message as is.
    • bold should wrap the message in two asterisks (**) to simulate bold text (e.g., **message**).
    • code should wrap the message in a single backtick (`) to simulate inline code (e.g., `message`)

## Exercise: Process Notifications
Textio now has a system to process different types of notifications: direct messages, group messages, and system alerts. Each notification type has a unique way of calculating its importance score based on user interaction and content.
Assignment
    1. Implement the importance methods for each message type. They should return the importance score for each message type.
        1. For a directMessage the importance score is based on if the message isUrgent or not. If it is urgent, the importance score is 50 otherwise the importance score is equal to the DM's priorityLevel.
        2. For a groupMessage the importance score is based on the message's priorityLevel
        3. All systemAlert messages should return a 100 importance score.
    2. Complete the processNotification function. It should identify the type and return different values for each type
        1. For a directMessage, return the sender's username and importance score.
        2. For a groupMessage, return the group's name and the importance score.
        3. For a systemAlert, return the alert code and the importance score.
        4. If the notification does not match any known type, return an empty string and a score of 0.