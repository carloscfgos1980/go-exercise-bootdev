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

func TestAnalyzeMessage(t *testing.T) {
	tests := []struct {
		initialAnalytics Analytics
		newMessage       Message
		expected         Analytics
	}{
		{
			initialAnalytics: Analytics{MessagesTotal: 0, MessagesFailed: 0, MessagesSucceeded: 0},
			newMessage:       Message{Recipient: "mickey", Success: true},
			expected:         Analytics{MessagesTotal: 1, MessagesFailed: 0, MessagesSucceeded: 1},
		},
		{
			initialAnalytics: Analytics{MessagesTotal: 1, MessagesFailed: 0, MessagesSucceeded: 1},
			newMessage:       Message{Recipient: "minnie", Success: false},
			expected:         Analytics{MessagesTotal: 2, MessagesFailed: 1, MessagesSucceeded: 1},
		},
		{
			initialAnalytics: Analytics{MessagesTotal: 2, MessagesFailed: 1, MessagesSucceeded: 1},
			newMessage:       Message{Recipient: "goofy", Success: false},
			expected:         Analytics{MessagesTotal: 3, MessagesFailed: 2, MessagesSucceeded: 1},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("initial: %+v, new: %+v", tt.initialAnalytics, tt.newMessage), func(t *testing.T) {
			AnalyzeMessage(&tt.initialAnalytics, tt.newMessage)
			if tt.initialAnalytics != tt.expected {
				t.Errorf("expected %+v, got %+v", tt.expected, tt.initialAnalytics)
			}
		})
	}
}

func TestRemoveProfanitySafe(t *testing.T) {
	s1 := "English, motherfubber, do you speak it?"
	s2 := "English, mother****er, do you speak it?"
	s3 := "Does he look like a witch?"
	s4 := "Does he look like a *****?"
	testCases := []struct {
		messageIn *string
		expected  *string
	}{
		{
			&s1,
			&s2,
		},
		{
			nil,
			nil,
		},
		{
			&s3,
			&s4,
		},
		{
			nil,
			nil,
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("messageIn: %v", tc.messageIn), func(t *testing.T) {
			RemoveProfanitySafe(tc.messageIn)
			if tc.messageIn != nil && tc.expected != nil && *tc.messageIn != *tc.expected {
				t.Errorf("expected %q, got %q", *tc.expected, *tc.messageIn)
			}
		})
	}
}
