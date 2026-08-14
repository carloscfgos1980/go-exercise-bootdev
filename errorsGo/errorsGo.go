package errorsgo

import (
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

// don't edit below this line

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
