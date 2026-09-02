# STRUCTS

## Structs in Go
We use structs in Go to represent structured data. It's often convenient to group different types of variables together. For example, if we want to represent a car we could do the following:
type car struct {
        brand      string
        model      string
        doors      int
        mileage    int
}
This creates a new struct type called car. All cars have a brand, model, doors and mileage.
To create a car, use a struct literal:
myCar := car{
        brand:   "Toyota",
        model:   "Camry",
        doors:   4,
        mileage: 5000,
}
Structs in Go are often used to represent data that you might use a dictionary or object for in other languages.

## Nested Structs in Go
Structs can be nested to represent more complex entities:
type car struct {
  brand string
  model string
  doors int
  mileage int
  frontWheel wheel
  backWheel wheel
}

type wheel struct {
  radius int
  material string
}
The fields of a struct can be accessed using the dot . operator.
myCar := car{}
myCar.frontWheel.radius = 5

## Anonymous Structs in Go
An anonymous struct is just like a normal struct, but it is defined without a name and therefore cannot be referenced elsewhere in the code.
To create an anonymous struct, just instantiate the instance immediately using a second pair of brackets after declaring the type:
myCar := struct {
  brand string
  model string
} {
  brand: "Toyota",
  model: "Camry",
}
You can even nest anonymous structs as fields within other structs:
type car struct {
  brand string
  model string
  doors int
  mileage int
  // wheel is a field containing an anonymous struct
  wheel struct {
    radius int
    material string
  }
}

var myCar = car{
  brand:   "Rezvani",
  model:   "Vengeance",
  doors:   4,
  mileage: 35000,
  wheel: struct {
    radius   int
    material string
  }{
    radius:   35,
    material: "alloy",
  },
}
When Should You Use an Anonymous Struct?
In general, prefer named structs. Named structs make it easier to read and understand your code, and they have the nice side-effect of being reusable. I sometimes use anonymous structs when I know I won't ever need to use a struct again. For example, sometimes I'll use one to create the shape of some JSON data in HTTP handlers.
If a struct is only meant to be used once, then it makes sense to declare it in such a way that developers down the road won't be tempted to accidentally use it again.

## Embedded Structs
Go is not an object-oriented language. However, embedded structs provide a kind of data-only inheritance that can be useful at times. Keep in mind, Go doesn't support classes or inheritance in the complete sense, but embedded structs are a way to elevate and share fields between struct definitions.
type car struct {
  brand string
  model string
}

type truck struct {
  // "car" is embedded, so the definition of a
  // "truck" now also additionally contains all
  // of the fields of the car struct
  car
  bedSize int
}
Embedded vs. Nested
    • Unlike nested structs, an embedded struct's fields are accessed at the top level like normal fields.
    • Like nested structs, you assign the promoted fields with the embedded struct in a composite literal.
lanesTruck := truck{
  bedSize: 10,
  car: car{
    brand: "Toyota",
    model: "Tundra",
  },
}

fmt.Println(lanesTruck.brand) // Toyota
fmt.Println(lanesTruck.model) // Tundra
In the example above, car is an embedded struct within truck. You can see that both brand and model are accessible from the top-level, while the nested equivalent to this object would require you to access these fields via a nested car struct: lanesTruck.car.brand or lanesTruck.car.model.

## Struct Methods in Go
While Go is not object-oriented, it does support methods that can be defined on structs. Methods are just functions that have a receiver. A receiver is a special parameter that syntactically goes before the name of the function.
type rect struct {
  width int
  height int
}

// area has a receiver of (r rect)
// rect is the struct
// r is the placeholder
func (r rect) area() int {
  return r.width * r.height
}

var r = rect{
  width: 5,
  height: 10,
}

fmt.Println(r.area())
// prints 50
A receiver is just a special kind of function parameter. In the example above, the r in (r rect) could just as easily have been rec or even x, y or z. By convention, Go code will often use the first letter of the struct's name.
Receivers are important because they will, as you'll learn in the exercises to come, allow us to define interfaces that our structs (and other types) can implement.


## Empty Struct Memory
Empty structs are Go's smallest possible type: they take up zero bytes of memory.
Compare their size with the other types below:
MEMORY USAGE
struct {}{} - 0 bytes
bool – 1 byte
uint16 – 2 bytes
int64 – 8 bytes

## Exercise: Update Users
We need a way to differentiate between standard and premium users. When a new user is created, they need a membership type, and that type will determine the message character limit.
Assignment
    1. Create a new struct called Membership, it should have:
        ◦ A Type string field
        ◦ A MessageCharLimit integer field
    2. Update the User struct to embed a Membership.
    3. Complete the newUser function. It should return a new User with all the fields set as you would expect based on the inputs. If the user is a "premium" member, the MessageCharLimit should be 1000, otherwise, it should only be 100.

## Exercise: Send Message
Assignment
Create a SendMessage method for the User struct.
It should take a message string and messageLength int as inputs.
If the messageLength is within the user's MessageCharLimit, return the original message and true (indicating the message can be sent), otherwise, return an empty string and false.