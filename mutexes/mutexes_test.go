package mutexes

import (
	"fmt"
	"sync"
	"testing"
)

func TestSafeCounter(t *testing.T) {
	tests := []struct {
		email string
		count int
	}{
		{"norman@bates.com", 23},
		{"marion@bates.com", 67},
		{"lila@bates.com", 31},
		{"sam@bates.com", 453},
	}
	for _, tt := range tests {
		sc := safeCounter{
			counts: make(map[string]int),
			mu:     &sync.Mutex{},
		}
		var wg sync.WaitGroup
		for i := 0; i < tt.count; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				sc.inc(tt.email)
			}()
		}
		wg.Wait()
		got := sc.val(tt.email)
		fmt.Printf(`
		  	email: %v
  			count: %v
  			expected count: %v
  			actual count:   %v
		`, tt.email, tt.count, tt.count, got)
		if got != tt.count {
			t.Errorf("safeCounter for %s = %d; want %d", tt.email, got, tt.count)
		}
	}
}
