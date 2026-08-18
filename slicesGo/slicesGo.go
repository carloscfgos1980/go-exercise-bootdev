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

// Variadic Functions
func Sum(nums ...int) int {
	var cost int
	for i := 0; i < len(nums); i++ {
		cost += nums[i]
	}
	return cost
}

// Append
type cost struct {
	day   int
	value float64
}

func GetDayCosts(costs []cost, day int) []float64 {
	dayCosts := make([]float64, 0)
	for i := 0; i < len(costs); i++ {
		if costs[i].day == day {
			dayCosts = append(dayCosts, costs[i].value)
		}
	}
	return dayCosts
}

// Range
func IndexOfFirstBadWord(msg []string, badWords []string) int {
	for i, word := range msg {
		for _, badWord := range badWords {
			if word == badWord {
				return i
			}
		}
	}
	return -1
}
