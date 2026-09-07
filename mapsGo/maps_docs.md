# MAPS DOCS

# Maps
Maps are similar to JavaScript objects, Python dictionaries, and Ruby hashes. Maps are a data structure that provides key->value mapping.
The zero value of a map is nil.
We can create a map by using the make() function:
ages := make(map[string]int)
ages["John"] = 37
ages["Mary"] = 24
ages["Mary"] = 21 // overwrites 24
Or by using a literal:
ages := map[string]int{
  "John": 37,
  "Mary": 21,
}
Map values can be structs too:
type car struct {
  registration string
  model        string
}

cars := map[string]car{
  "ABC-123": {registration: "ABC-123", model: "Civic"},
}

cars["XYZ-789"] = car{registration: "XYZ-789", model: "Accord"}
The len() function works on a map, it returns the total number of key/value pairs.
ages := map[string]int{
  "John": 37,
  "Mary": 21,
}
fmt.Println(len(ages)) // 2

# Mutations
Insert an Element
m[key] = elem
Get an Element
elem = m[key]
Delete an Element
delete(m, key)
Check If a Key Exists
elem, ok := m[key]
    • If key is in m, then ok is true and elem is the value as expected.
    • If key is not in the map, then ok is false and elem is the zero value for the map's element type.

# Key Types
Any type can be used as the value in a map, but keys are more restrictive.
Read the following section of the official Go blog:
As mentioned earlier, map keys may be of any type that is comparable. The language spec defines this precisely, but in short, comparable types are boolean, numeric, string, pointer, channel, and interface types, and structs or arrays that contain only those types. Notably absent from the list are slices, maps, and functions; these types cannot be compared using ==, and may not be used as map keys.
Questions:
1. What makes a type qualify to be used as a map key?
The type is comparable

# Exercise: Count Instances
Remember that you can check if a key is already present in a map by using the second return value from the index operation.
You can combine an if statement with an assignment operation to use the variables inside the if block:
names := map[string]int{}
missingNames := []string{}

if _, ok := names["Denna"]; !ok {
    // if the key doesn't exist yet,
    // append the name to the missingNames slice
    missingNames = append(missingNames, "Denna")
}
Assignment
Each time a user is sent a message, their username is logged in a slice. We want a more efficient way to count how many messages each user received.
Implement the updateCounts function. It takes as input:
    • messagedUsers: a slice of strings.
    • validUsers: a map of string -> int.
It should update the validUsers map with the number of times each user has received a message. Each string in the slice is a username, but they may not be valid. Only update the message count of valid users.
So, if "benji" is in the map and appears in the slice 3 times, the key "benji" in the map should have the value 3.

# Effective Go
Read the following paraphrased sections from effective Go regarding maps:
Like Slices, Maps Hold References
Like slices, maps hold references to an underlying data structure. If you pass a map to a function that changes the contents of the map, the changes will be visible in the caller.
Map Literals
Maps can be constructed using the usual composite literal syntax with colon-separated key-value pairs, so it's easy to build them during initialization.
var timeZone = map[string]int{
    "UTC":  0*60*60,
    "EST": -5*60*60,
    "CST": -6*60*60,
    "MST": -7*60*60,
    "PST": -8*60*60,
}
Missing Keys
An attempt to fetch a map value with a key that is not present in the map will return the zero value for the type of the entries in the map. For instance, if the map contains integers, looking up a non-existent key will return 0. A set can be implemented as a map with value type bool. Set the map entry to true to put the value in the set, and then test it by simple indexing.
attended := map[string]bool{
    "Ann": true,
    "Joe": true,
    ...
}

if attended[person] { // will be false if person is not in the map
    fmt.Println(person, "was at the meeting")
}
Sometimes you need to distinguish a missing entry from a zero value. Is there an entry for "UTC" or is that 0 because it's not in the map at all? You can discriminate with a form of multiple assignment.
var seconds int
var ok bool
seconds, ok = timeZone[tz]
For obvious reasons, this is called the "comma ok" idiom. In this example, if tz is present, seconds will be set appropriately and ok will be true; if not, seconds will be set to zero and ok will be false. Here's a function that puts it together with a nice error report:
func offset(tz string) int {
    if seconds, ok := timeZone[tz]; ok {
        return seconds
    }
    log.Println("unknown time zone:", tz)
    return 0
}
Now we can redo the first example of a set by using a more efficient map with empty structs, which don't take up any space in memory:
attended := map[string]struct{}{
    "Ann": {},
    "Joe": {},
    ...
}

if _, ok := attended[person]; ok {
    fmt.Println(person, "was at the meeting")
}
Deleting Map Entries
To delete a map entry, use the built-in delete function, whose arguments are the map and the key to be deleted. It's safe to do this even if the key is already absent from the map.
delete(timeZone, "PDT")  // Now on Standard Time

Questions:
1. Maps can have at most __1__ value(s) associated with the same key
2. Attempting to get a value from a map where the key doesn't exist
Return zero value
3. A function can mutate the values stored in a map and those changes __affect__ the original values
4. What does the second return value from a retrieve operation in a map indicate?
A boolean that indicates whether the key exists

# Nested
Maps can contain maps, creating a nested structure. For example:
map[string]map[string]int
map[rune]map[string]int
map[int]map[string]map[string]int

- Exercise: Distinct Words
Complete the countDistinctWords function using a map. It should take a slice of strings and return the total count of distinct words across all the strings. Assume words are separated by spaces. Casing should not matter. (e.g., "Hello" and "hello" should be considered the same word).
For example:
messages := []string{"Hello world", "hello there", "General Kenobi"}
count := countDistinctWords(messages)
count should be 5 as the distinct words are "hello", "world", "there", "general" and "kenobi" irrespective of casing.