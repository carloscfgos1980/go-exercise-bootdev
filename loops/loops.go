package loops

// loops in go
func BulkSend(numMessages int) float64 {
	totalCost := 0.0
	for i := 0; i < numMessages; i++ {
		totalCost += 1.0 + (float64(i) * 0.01)
	}
	return totalCost
}
