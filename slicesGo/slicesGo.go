package slicesgo

import "errors"

// Array in Go
func GetMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	messages := [3]string{primary, secondary, tertiary}
	msgsCosts := [3]int{len(primary), len(primary) + len(secondary), len(primary) + len(secondary) + len(tertiary)}
	return messages, msgsCosts
}

// Slices in Go
const (
	planFree = "free"
	planPro  = "pro"
)

func GetMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	if plan == planPro {
		return messages[:], nil
	}
	if plan == planFree {
		return messages[0:2], nil
	}
	return nil, errors.New("unsupported plan")

}

// Make
func GetMessageCosts(messages []string) []float64 {
	cost := make([]float64, len(messages))
	for i := 0; i < len(messages); i++ {
		cost[i] = float64(len(messages[i])) * 0.01
	}
	return cost
}
