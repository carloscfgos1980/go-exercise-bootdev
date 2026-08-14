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

// Type assertion in Go
func GetExpenseReport(e expense) (string, float64) {
	email, ok := e.(email)
	if ok {
		return email.toAddress, e.cost()
	}
	sms, ok := e.(sms)
	if ok {
		return sms.toPhoneNumber, e.cost()
	}
	return "", 0.0
}

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (em email) cost() float64 {
	if !em.isSubscribed {
		return float64(len(em.body)) * .05
	}
	return float64(len(em.body)) * .01
}

func (sm sms) cost() float64 {
	if !sm.isSubscribed {
		return float64(len(sm.body)) * .1
	}
	return float64(len(sm.body)) * .03
}

func (inv invalid) cost() float64 {
	return 0.0
}

// Type Switches
func GetExpenseReport2(e expense) (string, float64) {
	switch v := e.(type) {
	case email:
		return v.toAddress, v.cost()
	case sms:
		return v.toPhoneNumber, v.cost()
	default:
		return "", 0.0
	}
}

// Message formatters
type formatter interface {
	format() string
}

type plainText struct {
	message string
}

func (p plainText) format() string {
	return p.message
}

type bold struct {
	message string
}

func (b bold) format() string {
	return "**" + b.message + "**"
}

type code struct {
	message string
}

func (c code) format() string {
	return "`" + c.message + "`"
}

func SendMessage2(format formatter) string {
	return format.format() // Adjusted to call Format without an argument
}
