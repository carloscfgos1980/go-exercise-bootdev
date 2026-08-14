package interfaceGo

import (
	"fmt"
	"time"
)

// Interface in go
func SendMessage(msg message) (string, int) {
	message := msg.getMessage()
	cost := len(message) * 3
	return message, cost
}

type message interface {
	getMessage() string
}

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

// Interface implementation
func GetEmployeeSalary(e employee) int {
	return e.getSalary()
}

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hoursPerYear * c.hourlyPay
}

type fullTime struct {
	name   string
	salary int
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func (ft fullTime) getName() string {
	return ft.name
}

// Multiples interface
func SendMessage1(e expense1, f formatter1) (string, int) {
	return f.format(), e.cost()
}

func (e email1) cost() int {
	if !e.isSubscribed {
		return len(e.body) * 5
	}
	return len(e.body) * 2

}

func (e email1) format() string {
	subscribedText := "Subscribed"
	if !e.isSubscribed {
		subscribedText = "Not Subscribed"
	}
	return fmt.Sprintf("'%s' | %s", e.body, subscribedText)
}

type expense1 interface {
	cost() int
}

type formatter1 interface {
	format() string
}

type email1 struct {
	isSubscribed bool
	body         string
}
