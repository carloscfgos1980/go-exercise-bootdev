package mapsgo

import (
	"errors"
	"strings"
)

// Maps in Go
func GetUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}
	userMap := make(map[string]user)
	for i, name := range names {
		userMap[name] = user{
			name:        name,
			phoneNumber: phoneNumbers[i],
		}
	}
	return userMap, nil
}

type user struct {
	name        string
	phoneNumber int
}

// Mutations Maps in Go

func DeleteIfNecessary(users map[string]user2, name string) (deleted bool, err error) {
	u, ok := users[name]
	if !ok {
		return false, errors.New("not found")
	}
	if !u.scheduledForDeletion {
		return false, nil
	}
	delete(users, name)
	return true, nil
}

type user2 struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

// Count intances of a value in a map
func UpdateCounts(messagedUsers []string, validUsers map[string]int) {
	for _, user := range messagedUsers {
		if _, ok := validUsers[user]; ok {
			validUsers[user]++
		}
	}
}

// Nested Maps in Go
func GetNameCounts(names []string) map[rune]map[string]int {
	nameCounts := make(map[rune]map[string]int)
	for _, name := range names {
		runes := []rune(name)
		firstLetter := runes[0]

		if _, ok := nameCounts[firstLetter]; !ok {
			nameCounts[firstLetter] = make(map[string]int)
		}
		nameCounts[firstLetter][name]++
	}
	return nameCounts
}

// Distinct words
func CountDistinctWords(messages []string) int {
	distinctWordsStruct := make(map[string]struct{})
	for _, msg := range messages {
		for _, word := range strings.Fields(strings.ToLower(msg)) {
			if _, ok := distinctWordsStruct[word]; !ok {
				distinctWordsStruct[word] = struct{}{}
			}
		}
	}
	return len(distinctWordsStruct)
}
