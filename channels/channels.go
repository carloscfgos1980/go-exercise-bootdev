package channels

import (
	"fmt"
	"time"
)

// Concurrency is a powerful feature in Go that allows you to run multiple tasks simultaneously. In this example, we will demonstrate how to use goroutines and channels to send and receive messages concurrently.
func sendEmail(message string) {
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
	}()
	fmt.Printf("Email sent: '%s'\n", message)
}

func Test(message string) {
	sendEmail(message)
	time.Sleep(time.Millisecond * 500)
	fmt.Println("========================")
}

// Channels are a way to communicate between goroutines. They allow you to send and receive values of a specific type between different parts of your program. In this example, we will demonstrate how to use channels to send and receive messages concurrently.
type email struct {
	body string
	date time.Time
}

func CheckEmailAge(emails [3]email) [3]bool {
	isOldChan := make(chan bool)

	go sendIsOld(isOldChan, emails)

	isOld := [3]bool{}
	isOld[0] = <-isOldChan
	isOld[1] = <-isOldChan
	isOld[2] = <-isOldChan
	return isOld
}

// don't touch below this line

func sendIsOld(isOldChan chan<- bool, emails [3]email) {
	for _, e := range emails {
		if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
			isOldChan <- true
			continue
		}
		isOldChan <- false
	}
}
