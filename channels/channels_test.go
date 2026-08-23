package channels

import (
	"fmt"
	"testing"
	"time"
)

func TestCheckEmailAge(t *testing.T) {
	tests := []struct {
		emails [3]email
		isOld  [3]bool
	}{
		{[3]email{
			{
				body: "Words are pale shadows of forgotten names. As names have power, words have power.",
				date: time.Date(2019, 2, 0, 0, 0, 0, 0, time.UTC),
			},
			{
				body: "It's like everyone tells a story about themselves inside their own head.",
				date: time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC),
			},
			{
				body: "Bones mend. Regret stays with you forever.",
				date: time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC),
			},
		}, [3]bool{true, false, false}},
		{[3]email{
			{
				body: "Music is a proud, temperamental mistress.",
				date: time.Date(2018, 0, 0, 0, 0, 0, 0, time.UTC),
			},
			{
				body: "Have you heard of that website Boot.dev?",
				date: time.Date(2017, 0, 0, 0, 0, 0, 0, time.UTC),
			},
			{
				body: "It's awesome honestly.",
				date: time.Date(2016, 0, 0, 0, 0, 0, 0, time.UTC),
			},
		}, [3]bool{true, true, true}},
		{[3]email{
			{
				body: "I have stolen princesses back from sleeping barrow kings.",
				date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC),
			},
			{
				body: "I burned down the town of Trebon",
				date: time.Date(2019, 6, 6, 0, 0, 0, 0, time.UTC),
			},
			{
				body: "I have spent the night with Felurian and left with both my sanity and my life.",
				date: time.Date(2022, 7, 0, 0, 0, 0, 0, time.UTC),
			},
		}, [3]bool{true, true, false}},
	}
	for _, tt := range tests {
		got := CheckEmailAge(tt.emails)
		fmt.Printf("CheckEmailAge(%v) = %v; want %v\n", tt.emails, got, tt.isOld)
		for i := range got {
			if got[i] != tt.isOld[i] {
				t.Errorf("CheckEmailAge(%v) = %v; want %v", tt.emails, got, tt.isOld)
			}
		}
	}
}

func TestWaitForDBs(t *testing.T) {
	tests := []struct {
		numDBs int
	}{
		{numDBs: 3},
		{numDBs: 5},
		{numDBs: 10},
	}
	for _, tt := range tests {
		dbChan, countPtr := getDBsChannel(tt.numDBs)
		WaitForDBs(tt.numDBs, dbChan)
		if len(dbChan) != 0 && *countPtr != tt.numDBs {
			t.Errorf("WaitForDBs(%d) = %d; want %d. Len(dbChan) = %d", tt.numDBs, *countPtr, tt.numDBs, len(dbChan))
		}
	}
}
