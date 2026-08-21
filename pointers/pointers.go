package pointers

import (
	"errors"
	"strings"
)

// Pointers dereference the value of a variable by using the * operator. This allows you to modify the value of a variable directly, rather than working with a copy of the value. In this example, we use pointers to remove profanity from a message string by modifying the original string in place.
func RemoveProfanity(message *string) {
	*message = strings.ReplaceAll(*message, "fubb", "****")
	*message = strings.ReplaceAll(*message, "shiz", "****")
	*message = strings.ReplaceAll(*message, "witch", "*****")
}

// Pay by reference allows you to pass a pointer to a variable instead of a copy of the variable's value. This can be more efficient for large data structures, as it avoids the overhead of copying the data. In this example, we use pointers to modify the original Analytics struct in place, rather than working with a copy of the struct.
type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

func AnalyzeMessage(a *Analytics, msg Message) {
	if msg.Success {
		a.MessagesSucceeded++
	} else {
		a.MessagesFailed++
	}
	a.MessagesTotal++
}

// Nil pointers
func RemoveProfanitySafe(message *string) {
	if message == nil {
		return
	}
	messageVal := *message
	messageVal = strings.ReplaceAll(messageVal, "fubb", "****")
	messageVal = strings.ReplaceAll(messageVal, "shiz", "****")
	messageVal = strings.ReplaceAll(messageVal, "witch", "*****")
	*message = messageVal
}

// Update Balance

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

func UpdateBalance(c *customer, t transaction) error {
	switch t.transactionType {
	case transactionDeposit:
		c.balance += t.amount
		return nil
	case transactionWithdrawal:
		if c.balance < t.amount {
			return errors.New("insufficient funds")
		} else {
			c.balance -= t.amount
			return nil
		}
	default:
		return errors.New("unknown transaction type")
	}
}
