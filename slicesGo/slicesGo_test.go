package slicesgo

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestGetMessageWithRetries(t *testing.T) {
	tests := []struct {
		messages         []string
		expectedMessages [3]string
		expectedCosts    [3]int
	}{
		{
			[]string{
				"Hello sir/madam can I interest you in a yacht?",
				"Please I'll even give you an Amazon gift card?",
				"You're missing out big time",
			},
			[3]string{
				"Hello sir/madam can I interest you in a yacht?",
				"Please I'll even give you an Amazon gift card?",
				"You're missing out big time",
			},
			[3]int{46, 92, 119},
		},
		{
			[]string{"It's the spring fling sale!", "Don't miss this event!", "Last chance."},
			[3]string{"It's the spring fling sale!", "Don't miss this event!", "Last chance."},
			[3]int{27, 49, 61},
		},
	}
	for _, tt := range tests {
		resultMessages, resultCosts := GetMessageWithRetries(tt.messages[0], tt.messages[1], tt.messages[2])
		fmt.Printf("GetMessageWithRetries(%v) = %v, %v; want %v, %v\n", tt.messages, resultMessages, resultCosts, tt.expectedMessages, tt.expectedCosts)
		if resultMessages != tt.expectedMessages || resultCosts != tt.expectedCosts {
			t.Errorf("GetMessageWithRetries(%v) = %v, %v; want %v, %v", tt.messages, resultMessages, resultCosts, tt.expectedMessages, tt.expectedCosts)
		}
	}
}

func TestGetMessageWithRetriesForPlan(t *testing.T) {
	tests := []struct {
		plan             string
		messages         [3]string
		expectedMessages []string
		expectedErr      error
	}{
		{
			planFree,
			[3]string{
				"Hello sir/madam can I interest you in a yacht?",
				"Please I'll even give you an Amazon gift card?",
				"You're missing out big time",
			},
			[]string{"Hello sir/madam can I interest you in a yacht?", "Please I'll even give you an Amazon gift card?"},
			nil,
		},
		{
			planPro,
			[3]string{
				"Hello sir/madam can I interest you in a yacht?",
				"Please I'll even give you an Amazon gift card?",
				"You're missing out big time",
			},
			[]string{
				"Hello sir/madam can I interest you in a yacht?",
				"Please I'll even give you an Amazon gift card?",
				"You're missing out big time",
			},
			nil,
		},
		{
			"invalid plan",
			[3]string{
				"You can get a good look at a T-bone by sticking your head up a bull's ass, but wouldn't you rather take the butcher's word for it?",
				"Wouldn't you?",
				"Wouldn't you???",
			},
			nil,
			fmt.Errorf("unsupported plan"),
		},
	}
	for _, tt := range tests {
		resultMessages, resultErr := GetMessageWithRetriesForPlan(tt.plan, tt.messages)
		fmt.Printf("GetMessageWithRetriesForPlan(%v, %v) = %v, %v; want %v, %v\n", tt.plan, tt.messages, resultMessages, resultErr, tt.expectedMessages, tt.expectedErr)
		if !equalSlices(resultMessages, tt.expectedMessages) || !equalErrors(resultErr, tt.expectedErr) {
			t.Errorf("GetMessageWithRetriesForPlan(%v, %v) = %v, %v; want %v, %v", tt.plan, tt.messages, resultMessages, resultErr, tt.expectedMessages, tt.expectedErr)
		}
	}

}

func equalSlices[S ~[]E, E any](a, b S) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !equalValues(a[i], b[i]) {
			return false
		}
	}
	return true
}

func equalValues[T any](a, b T) bool {
	return equalReflectValues(reflect.ValueOf(a), reflect.ValueOf(b))
}

func equalReflectValues(a, b reflect.Value) bool {
	if !a.IsValid() || !b.IsValid() {
		return !a.IsValid() && !b.IsValid()
	}
	if a.Kind() != b.Kind() {
		return false
	}
	if a.Kind() == reflect.Slice {
		if a.Len() != b.Len() {
			return false
		}
		for i := 0; i < a.Len(); i++ {
			if !equalReflectValues(a.Index(i), b.Index(i)) {
				return false
			}
		}
		return true
	}
	if a.Kind() == reflect.Float32 || a.Kind() == reflect.Float64 {
		fa := a.Float()
		fb := b.Float()
		if math.IsNaN(fa) || math.IsNaN(fb) {
			return false
		}
		return fa == fb
	}
	return reflect.DeepEqual(a.Interface(), b.Interface())
}

func equalErrors(a, b error) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Error() == b.Error()
}

func TestGetMessageCosts(t *testing.T) {
	tests := []struct {
		messages    []string
		expected    []float64
		expectedCap int
	}{
		{
			[]string{"Welcome to the movies!", "Enjoy your popcorn!"},
			[]float64{0.22, 0.19},
			2,
		},
		{
			[]string{"I don't want to be here anymore", "Can we go home?", "I'm hungry", "I'm bored"},
			[]float64{0.31, 0.15, 0.1, 0.09},
			4,
		},
		{[]string{}, []float64{}, 0},
		{[]string{""}, []float64{0}, 1},
		{[]string{"Hello", "Hi", "Hey"}, []float64{0.05, 0.02, 0.03}, 3},
	}
	for _, tt := range tests {
		cost := GetMessageCosts(tt.messages)
		fmt.Printf("GetMessageCosts(%v) = %v; want %v; want capacity = %v, got %v\n", tt.messages, cost, tt.expected, tt.expectedCap, cap(cost))
		if !equalSlices(tt.expected, cost) || cap(cost) != tt.expectedCap {
			t.Errorf("GetMessageCosts(%v) = %v; want %v; cap = %v; want cap = %v", tt.messages, cost, tt.expected, cap(cost), tt.expectedCap)
		}
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 55},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 120},
		{[]int{}, 0},
		{[]int{5}, 5},
	}
	for _, tt := range tests {
		result := Sum(tt.nums...)
		fmt.Printf("Sum(%v) = %v; want %v\n", tt.nums, result, tt.expected)
		if result != tt.expected {
			t.Errorf("Sum(%v) = %v; want %v", tt.nums, result, tt.expected)
		}
	}
}

func TestGetDayCosts(t *testing.T) {
	tests := []struct {
		costs    []cost
		day      int
		expected []float64
	}{
		{
			costs: []cost{
				{0, 1.0},
				{1, 2.0},
				{1, 3.1},
				{5, 2.5},
				{2, 3.6},
				{1, 2.7},
				{1, 3.3},
			},
			day: 1,
			expected: []float64{
				2.0,
				3.1,
				2.7,
				3.3,
			},
		},
	}
	for _, tt := range tests {
		result := GetDayCosts(tt.costs, tt.day)
		fmt.Printf("GetDayCosts(%v, %v) = %v; want %v\n", tt.costs, tt.day, result, tt.expected)
		if !equalSlices(result, tt.expected) {
			t.Errorf("GetDayCosts(%v, %v) = %v; want %v", tt.costs, tt.day, result, tt.expected)
		}
	}
}

func TestIndexOfFirstBadWord(t *testing.T) {
	tests := []struct {
		msg      []string
		badWords []string
		expected int
	}{
		{[]string{"hey", "there", "john"}, []string{"crap", "shoot", "frick", "dang"}, -1},
		{[]string{"ugh", "oh", "my", "frick"}, []string{"crap", "shoot", "frick", "dang"}, 3},
		{[]string{"what", "the", "shoot", "I", "hate", "that", "crap"}, []string{"crap", "shoot", "frick", "dang"}, 2},
		{[]string{"crap", "shoot", "frick", "dang"}, []string{""}, -1},
		{[]string{""}, nil, -1},
	}
	for _, tt := range tests {
		result := IndexOfFirstBadWord(tt.msg, tt.badWords)
		fmt.Printf("IndexOfFirstBadWord(%v, %v) = %v; want %v\n", tt.msg, tt.badWords, result, tt.expected)
		if result != tt.expected {
			t.Errorf("IndexOfFirstBadWord(%v, %v) = %v; want %v", tt.msg, tt.badWords, result, tt.expected)
		}
	}
}

func TestCreateMatrix(t *testing.T) {
	tests := []struct {
		rows, cols int
		expected   [][]int
	}{
		{3, 3, [][]int{
			{0, 0, 0},
			{0, 1, 2},
			{0, 2, 4},
		}},
		{4, 4, [][]int{
			{0, 0, 0, 0},
			{0, 1, 2, 3},
			{0, 2, 4, 6},
			{0, 3, 6, 9},
		}},
		{5, 7, [][]int{
			{0, 0, 0, 0, 0, 0, 0},
			{0, 1, 2, 3, 4, 5, 6},
			{0, 2, 4, 6, 8, 10, 12},
			{0, 3, 6, 9, 12, 15, 18},
			{0, 4, 8, 12, 16, 20, 24},
		}},
		{0, 0, [][]int{}},
	}
	for _, tt := range tests {
		result := CreateMatrix(tt.rows, tt.cols)
		fmt.Printf("CreateMatrix(%v, %v) = %v; want %v\n", tt.rows, tt.cols, result, tt.expected)
		if !equalSlices(result, tt.expected) {
			t.Errorf("CreateMatrix(%v, %v) = %v; want %v", tt.rows, tt.cols, result, tt.expected)
		}
	}
}

func TestFilterMessages(t *testing.T) {
	messages := []Message{
		TextMessage{"Alice", "Hello, World!"},
		MediaMessage{"Bob", "image", "A beautiful sunset"},
		LinkMessage{"Charlie", "http://example.com", "Example Domain"},
		TextMessage{"Dave", "Another text message"},
		MediaMessage{"Eve", "video", "Cute cat video"},
		LinkMessage{"Frank", "https://boot.dev", "Learn Coding Online"},
	}
	tests := []struct {
		filterType    string
		expectedCount int
		expectedType  string
	}{
		{"text", 2, "text"},
		{"media", 2, "media"},
		{"link", 2, "link"},
		{"unsupported", 0, "unsupported"},
	}
	for _, tt := range tests {
		result := FilterMessages(messages, tt.filterType)
		fmt.Printf("FilterMessages(%v, %v) = %v; want count %v and type %v\n", messages, tt.filterType, result, tt.expectedCount, tt.expectedType)
		if len(result) != tt.expectedCount {
			t.Errorf("FilterMessages(%v, %v) = %v; want count %v", messages, tt.filterType, result, tt.expectedCount)
		}
		for _, msg := range result {
			if msg.Type() != tt.expectedType {
				t.Errorf("FilterMessages(%v, %v) = %v; want type %v", messages, tt.filterType, result, tt.expectedType)
			}
		}
	}
}
