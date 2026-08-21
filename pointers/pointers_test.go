package pointers

import (
	"fmt"
	"testing"
)

func TestRemoveProfanity(t *testing.T) {
	testCases := []struct {
		messageIn string
		expected  string
	}{
		{
			"English, motherfubber, do you speak it?",
			"English, mother****er, do you speak it?",
		},
		{
			"Oh man I've seen some crazy ass shiz in my time...",
			"Oh man I've seen some crazy ass **** in my time...",
		},
		{
			"Does he look like a witch?",
			"Does he look like a *****?",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.messageIn, func(t *testing.T) {
			RemoveProfanity(&tc.messageIn)
			fmt.Printf("messageIn: %q, expected: %q\n", tc.messageIn, tc.expected)
			if tc.messageIn != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, tc.messageIn)
			}
		})
	}
}
