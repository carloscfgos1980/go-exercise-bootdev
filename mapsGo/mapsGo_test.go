package mapsgo

import (
	"fmt"
	"testing"
)

func equalMaps(a, b map[string]user) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if vb, ok := b[k]; !ok || vb != v {
			return false
		}
	}
	return true
}

func TestGetUserMap(t *testing.T) {
	tests := []struct {
		names        []string
		phoneNumbers []int
		expected     map[string]user
		errString    string
	}{
		{
			[]string{"Eren", "Armin", "Mikasa"},
			[]int{14355550987, 98765550987, 18265554567},
			map[string]user{"Eren": {"Eren", 14355550987}, "Armin": {"Armin", 98765550987}, "Mikasa": {"Mikasa", 18265554567}},
			"",
		},
		{
			[]string{"Eren", "Armin"},
			[]int{14355550987, 98765550987, 18265554567},
			nil,
			"invalid sizes",
		},
		{
			[]string{"George", "Annie", "Reiner", "Sasha"},
			[]int{20955559812, 38385550982, 48265554567, 16045559873},
			map[string]user{"George": {"George", 20955559812}, "Annie": {"Annie", 38385550982}, "Reiner": {"Reiner", 48265554567}, "Sasha": {"Sasha", 16045559873}},
			"",
		},
		{
			[]string{"George", "Annie", "Reiner"},
			[]int{20955559812, 38385550982, 48265554567, 16045559873},
			nil,
			"invalid sizes",
		},
	}
	for _, tt := range tests {
		result, err := GetUserMap(tt.names, tt.phoneNumbers)
		fmt.Printf("GetUserMap(%v, %v) = %v, %v; want %v, %v\n", tt.names, tt.phoneNumbers, result, err, tt.expected, tt.errString)
		if err != nil {
			if err.Error() != tt.errString {
				t.Errorf("GetUserMap(%v, %v) returned unexpected error: got %v, want %v", tt.names, tt.phoneNumbers, err.Error(), tt.errString)
			}
			continue
		}
		if !equalMaps(result, tt.expected) {
			t.Errorf("GetUserMap(%v, %v) = %v; want %v", tt.names, tt.phoneNumbers, result, tt.expected)
		}
	}
}

func equalUser2Maps(a, b map[string]user2) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if vb, ok := b[k]; !ok || vb != v {
			return false
		}
	}
	return true
}
func TestDeleteIfNecessary(t *testing.T) {
	tests := []struct {
		users             map[string]user2
		name              string
		expectedErrString string
		expectedUsers     map[string]user2
		expectedDeleted   bool
	}{
		{
			map[string]user2{"Erwin": {"Erwin", 14355550987, true}, "Levi": {"Levi", 98765550987, true}, "Hanji": {"Hanji", 18265554567, true}},
			"Erwin",
			"",
			map[string]user2{"Levi": {"Levi", 98765550987, true}, "Hanji": {"Hanji", 18265554567, true}},
			true,
		},
		{
			map[string]user2{"Erwin": {"Erwin", 14355550987, false}, "Levi": {"Levi", 98765550987, false}, "Hanji": {"Hanji", 18265554567, false}},
			"Erwin",
			"",
			map[string]user2{"Erwin": {"Erwin", 14355550987, false}, "Levi": {"Levi", 98765550987, false}, "Hanji": {"Hanji", 18265554567, false}},
			false,
		},
		{
			map[string]user2{"Erwin": {"Erwin", 14355550987, true}, "Levi": {"Levi", 98765550987, true}, "Hanji": {"Hanji", 18265554567, true}},
			"Eren",
			"not found",
			map[string]user2{"Erwin": {"Erwin", 14355550987, true}, "Levi": {"Levi", 98765550987, true}, "Hanji": {"Hanji", 18265554567, true}},
			false,
		},
	}
	for _, tt := range tests {
		deleted, err := DeleteIfNecessary(tt.users, tt.name)
		fmt.Printf("DeleteIfNecessary(%v, %v) = %v, %v; want %v, %v\n", tt.users, tt.name, deleted, err, tt.expectedDeleted, tt.expectedErrString)
		if err != nil {
			if err.Error() != tt.expectedErrString {
				t.Errorf("DeleteIfNecessary(%v, %v) returned unexpected error: got %v, want %v", tt.users, tt.name, err.Error(), tt.expectedErrString)
			}
			continue
		}
		if deleted != tt.expectedDeleted {
			t.Errorf("DeleteIfNecessary(%v, %v) = %v; want %v", tt.users, tt.name, deleted, tt.expectedDeleted)
		}
		if !equalUser2Maps(tt.users, tt.expectedUsers) {
			t.Errorf("DeleteIfNecessary(%v, %v) modified users map incorrectly: got %v; want %v", tt.users, tt.name, tt.users, tt.expectedUsers)
		}
	}
}

func TestUpdateCounts(t *testing.T) {
	tests := []struct {
		messagedUsers []string
		validUsers    map[string]int
		expected      map[string]int
	}{
		{
			[]string{"Eren", "Armin", "Mikasa", "Eren", "Mikasa"},
			map[string]int{"Eren": 0, "Armin": 0, "Mikasa": 0},
			map[string]int{"Eren": 2, "Armin": 1, "Mikasa": 2},
		},
		{
			[]string{"Levi", "Erwin", "Hanji", "Levi", "Hanji"},
			map[string]int{"Levi": 0, "Erwin": 0, "Hanji": 0},
			map[string]int{"Levi": 2, "Erwin": 1, "Hanji": 2},
		},
	}
	for _, tt := range tests {
		UpdateCounts(tt.messagedUsers, tt.validUsers)
		fmt.Printf("UpdateCounts(%v, %v) = %v; want %v\n", tt.messagedUsers, tt.validUsers, tt.validUsers, tt.expected)
		if !equalIntMaps(tt.validUsers, tt.expected) {
			t.Errorf("UpdateCounts(%v, %v) modified validUsers map incorrectly: got %v; want %v", tt.messagedUsers, tt.validUsers, tt.validUsers, tt.expected)
		}
	}
}

func equalIntMaps(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if vb, ok := b[k]; !ok || vb != v {
			return false
		}
	}
	return true
}

func TestGetNameCounts(t *testing.T) {
	tests := []struct {
		names                    []string
		initial                  rune
		expectedNamesWithInitial int
		name                     string
		expectedNameCount        int
	}{
		{getNames(50), 'M', 4, "Matthew", 3},
		{getNames(100), 'G', 6, "George", 1},
		{getNames(300), '😊', 1, "😊", 1},
		{getNames(150), 'D', 13, "Drew", 4},
		{getNames(200), 'P', 9, "Philip", 4},
		{getNames(250), 'B', 16, "Bryant", 1},
	}
	for _, tt := range tests {
		result := GetNameCounts(tt.names)
		fmt.Printf(`---------------------------------
  len(names): %v
  initial: %c
  expected names with initial: %d
  actual names with initial: %d
  name: %s
  expected count: %d
  actual count: %d
`, len(tt.names), tt.initial, tt.expectedNamesWithInitial, len(result[tt.initial]), tt.name, tt.expectedNameCount, result[tt.initial][tt.name])
		if len(result[tt.initial]) != tt.expectedNamesWithInitial {
			t.Errorf("GetNameCounts(%v) returned unexpected number of names with initial %c: got %d; want %d", tt.names, tt.initial, len(result[tt.initial]), tt.expectedNamesWithInitial)
		}
		if result[tt.initial][tt.name] != tt.expectedNameCount {
			t.Errorf("GetNameCounts(%v) returned unexpected count for name %s: got %d; want %d", tt.names, tt.name, result[tt.initial][tt.name], tt.expectedNameCount)
		}
	}
}

func getNames(length int) []string {
	return []string{
		"Grant", "Eduardo", "Peter", "Matthew", "Matthew", "Matthew", "Peter", "Peter", "Henry", "Parker",
		"Parker", "Parker", "Collin", "Hayden", "George", "Bradley", "Mitchell", "Devon", "Ricardo", "Shawn",
		"Taylor", "Nicolas", "Gregory", "Francisco", "Liam", "Kaleb", "Preston", "Erik", "Alexis", "Owen",
		"Omar", "Diego", "Dustin", "Corey", "Fernando", "Clayton", "Carter", "Ivan", "Jaden", "Javier",
		"Alec", "Johnathan", "Scott", "Manuel", "Cristian", "Alan", "Raymond", "Brett", "Max", "Drew",
		"Andres", "Gage", "Mario", "Dawson", "Dillon", "Cesar", "Wesley", "Levi", "Jakob", "Chandler",
		"Martin", "Malik", "Edgar", "Sergio", "Trenton", "Josiah", "Nolan", "Marco", "Drew", "Peyton",
		"Harrison", "Drew", "Hector", "Micah", "Roberto", "Drew", "Brady", "Erick", "Conner", "Jonah",
		"Casey", "Jayden", "Edwin", "Emmanuel", "Andre", "Phillip", "Brayden", "Landon", "Giovanni", "Bailey",
		"Ronald", "Braden", "Damian", "Donovan", "Ruben", "Frank", "Gerardo", "Pedro", "Andy", "Chance",
		"Abraham", "Calvin", "Trey", "Cade", "Donald", "Derrick", "Payton", "Darius", "Enrique", "Keith",
		"Raul", "Jaylen", "Troy", "Jonathon", "Cory", "Marc", "Eli", "Skyler", "Rafael", "Trent",
		"Griffin", "Colby", "Johnny", "Chad", "Armando", "Kobe", "Caden", "Marcos", "Cooper", "Elias",
		"Brenden", "Israel", "Avery", "Zane", "Zane", "Zane", "Zane", "Dante", "Josue", "Zackary",
		"Allen", "Philip", "Mathew", "Dennis", "Leonardo", "Ashton", "Philip", "Philip", "Philip", "Julio",
		"Miles", "Damien", "Ty", "Gustavo", "Drake", "Jaime", "Simon", "Jerry", "Curtis", "Kameron",
		"Lance", "Brock", "Bryson", "Alberto", "Dominick", "Jimmy", "Kaden", "Douglas", "Gary", "Brennan",
		"Zachery", "Randy", "Louis", "Larry", "Nickolas", "Albert", "Tony", "Fabian", "Keegan", "Saul",
		"Danny", "Tucker", "Myles", "Damon", "Arturo", "Corbin", "Deandre", "Ricky", "Kristopher", "Lane",
		"Pablo", "Darren", "Jarrett", "Zion", "Alfredo", "Micheal", "Angelo", "Carl", "Oliver", "Kyler",
		"Tommy", "Walter", "Dallas", "Jace", "Quinn", "Theodore", "Grayson", "Lorenzo", "Joe", "Arthur",
		"Bryant", "Roman", "Brent", "Russell", "Ramon", "Lawrence", "Moises", "Aiden", "Quentin", "Jay",
		"Tyrese", "Tristen", "Emanuel", "Salvador", "Terry", "Morgan", "Jeffery", "Esteban", "Tyson", "Braxton",
		"Branden", "Marvin", "Brody", "Craig", "Ismael", "Rodney", "Isiah", "Marshall", "Maurice", "Ernesto",
		"Emilio", "Brendon", "Kody", "Eddie", "Malachi", "Abel", "Keaton", "Jon", "Shaun", "Skylar",
		"Ezekiel", "Nikolas", "Santiago", "Kendall", "Axel", "Camden", "Trevon", "Bobby", "Conor", "Jamal",
		"Lukas", "Malcolm", "Zackery", "Jayson", "Javon", "Roger", "Reginald", "Zachariah", "Desmond", "Felix",
		"Johnathon", "Dean", "Quinton", "Ali", "Davis", "Gerald", "Rodrigo", "Demetrius", "Billy", "Rene",
		"Reece", "Kelvin", "Leo", "Justice", "Chris", "Guillermo", "Matthew", "Matthew", "Matthew", "Kevon",
		"Steve", "Frederick", "Clay", "Weston", "Dorian", "Hugo", "Roy", "Orlando", "Terrance", "😊",
		"Kai", "Khalil", "Khalil", "Khalil", "Graham", "Noel", "Willie", "Nathanael", "Terrell",
	}[:length]
}

func TestCountDistinctWords(t *testing.T) {
	tests := []struct {
		messages []string
		expected int
	}{
		{
			[]string{"WTS Arcanite Bar! Cheaper than AH", "Do you need an Arcanite Bar!"},
			10,
		},
		{
			[]string{"Could you give me a number crunch real quick?", "Looks like we have a 32.33% (repeating of course) percentage of survival."},
			19,
		},
		{
			[]string{"LFG UBRS", "lfg ubrs", "LFG Ubrs"},
			2,
		},
		{
			[]string{"Alright time's up! Let's do this.", "Leroy Jenkins!", "Damn it Leroy"},
			10,
		},
		{
			[]string{"I'm out of range", "I'm out of mana"},
			5,
		},
		{
			[]string{
				"LF9M UBRS need all",
				"LF8M UBRS need all",
				"LF7M UBRS need all",
				"LF6M UBRS need tanks and heals",
				"LF5M UBRS need tanks and heals",
				"LF4M UBRS need tanks and heals",
				"LF3M UBRS need tanks and healer",
				"LF2M UBRS need tanks",
				"LF1M UBRS need tank",
				"Group is full thanks!",
			},
			21,
		},
		{
			[]string{""},
			0,
		},
	}

	for _, tt := range tests {
		result := CountDistinctWords(tt.messages)
		fmt.Printf("CountDistinctWords(%v) = %d; want %d\n", tt.messages, result, tt.expected)
		if result != tt.expected {
			t.Errorf("CountDistinctWords(%v) = %d; want %d", tt.messages, result, tt.expected)
		}
	}
}
