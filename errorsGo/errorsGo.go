package errorsgo

import (
	"errors"
	"fmt"
)

// Error interface
func SendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	costMsgCostumer, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0, err
	}
	costMsgSpouse, err := sendSMS(msgToSpouse)
	if err != nil {
		return 0, err
	}
	return costMsgCostumer + costMsgSpouse, nil
}

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}

// Formating Strings Review
func GetSMSErrorString(cost float64, recipient string) string {
	return fmt.Sprintf("SMS that costs $%.2f to be sent to '%s' cannot be sent", cost, recipient)
}

// Custom error type

type divideError struct {
	dividend float64
}

func (de divideError) Error() string {
	return fmt.Sprintf("cannot divide %v by zero", de.dividend)
}

func Divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0.0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}

// Errors package
func Divide2(x, y float64) (float64, error) {
	if y == 0 {
		return 0.0, errors.New("no dividing by 0")
	}
	return x / y, nil
}
