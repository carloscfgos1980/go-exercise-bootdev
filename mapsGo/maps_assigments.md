# Maps assigments


## Maps. Get user map

We can speed up our contact-info lookups by using a map!

Key-based map lookup: O(1)
Slice brute-force search: O(n)
Complete the getUserMap function. It takes a slice of names and a slice of phone numbers, and returns a map of name -> user structs and an error. A user struct just contains a user's name and phone number. The first element in the names slice pairs with the first phone number, and so on.

If the length of names and phoneNumbers is not equal, return an error with the string "invalid sizes".

## Mutations. Deleteif necessary

It's important to keep up with privacy regulations and to respect our user's data. We need a function that will delete user records.

Complete the deleteIfNecessary function. The user struct has a scheduledForDeletion field that determines if they are scheduled for deletion or not.

If the user doesn't exist in the map, return false and the error not found.
If they exist but aren't scheduled for deletion, return deleted as false with no errors.
If they exist and are scheduled for deletion, return deleted as true with no errors and delete their record from the map.
Like slices, maps are reference types. When a map is passed into a function, we can make changes to the original – we don't have a copy.

## Count instances

Each time a user is sent a message, their username is logged in a slice. We want a more efficient way to count how many messages each user received.

Implement the updateCounts function. It takes as input:

messagedUsers: a slice of strings.
validUsers: a map of string -> int.
It should update the validUsers map with the number of times each user has received a message. Each string in the slice is a username, but they may not be valid. Only update the message count of valid users.

So, if "benji" is in the map and appears in the slice 3 times, the key "benji" in the map should have the value 3.

## Nesteed. Get names count

Because Textio is a glorified customer database, we have a lot of internal logic for sorting and dealing with customer names.

Complete the getNameCounts function. It takes a slice of strings names and returns a nested map. The parent map's keys are all the unique first characters (see runes) of the names, the nested maps keys are all the names themselves, and the value is the count of each name.

For example:

billy
billy
bob
joe

Creates the following nested map:

b: {
    billy: 2,
    bob: 1
},
j: {
    joe: 1
}

## Distinct Words
Complete the countDistinctWords function using a map. It should take a slice of strings and return the total count of distinct words across all the strings. Assume words are separated by spaces. Casing should not matter. (e.g., "Hello" and "hello" should be considered the same word).

For example:

messages := []string{"Hello world", "hello there", "General Kenobi"}
count := countDistinctWords(messages)

count should be 5 as the distinct words are "hello", "world", "there", "general" and "kenobi" irrespective of casing.

Tips
Go's strings package can be very helpful here. Specifically, the Fields and ToLower functions.
Since all that matters is counting distinct words, we don't care about the value of each key in the map. You can use struct{}{} as the value of your map key. This empty struct uses no memory. For example, both map[string]bool and map[string]struct{} can track unique words, but they use different amounts of memory:
distinctWordsBool := make(map[string]bool)
distinctWordsStruct := make(map[string]struct{})

distinctWordsBool["hello"] = true         // Uses 1 byte for the bool value
distinctWordsStruct["hello"] = struct{}{} // Uses 0 bytes for the empty struct