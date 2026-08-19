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
