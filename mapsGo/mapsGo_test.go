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
