package errorsgo

import (
	"fmt"
	"testing"
)

func TestSendSMSToCouple(t *testing.T) {
	tests := []struct {
		msgToCustomer string
		msgToSpouse   string
		expectedCost  int
		expectedErr   error
	}{
		{"Thanks for coming in to our flower shop today!", "We hope you enjoyed your gift.", 0, fmt.Errorf("can't send texts over 25 characters")},
		{"Thanks for joining us!", "Have a good day.", 76, nil},
		{"Thank you.", "Enjoy!", 32, nil},
		{"We loved having you in!", "We hope the rest of your evening is fantastic.", 0, fmt.Errorf("can't send texts over 25 characters")},
	}
	for _, tt := range tests {
		cost, err := SendSMSToCouple(tt.msgToCustomer, tt.msgToSpouse)
		fmt.Printf("Testing SendSMSToCouple(%q, %q), got cost %v and error %v\n", tt.msgToCustomer, tt.msgToSpouse, cost, err)
		if cost != tt.expectedCost || (err != nil && err.Error() != tt.expectedErr.Error()) {
			t.Errorf("SendSMSToCouple(%q, %q) = %v, %v; want %v, %v",
				tt.msgToCustomer, tt.msgToSpouse, cost, err, tt.expectedCost, tt.expectedErr)
		}
	}
}
