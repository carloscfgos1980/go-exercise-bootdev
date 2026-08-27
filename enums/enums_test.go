package enums

import (
	"fmt"
	"testing"
)

func TestHandleEmailBounce(t *testing.T) {
	tests := []struct {
		email           email
		expectedError   string
		expectedStatus  string
		expectedBounces int
	}{
		{
			email: email{
				status:    "email_bounced",
				recipient: &user{email: "bugs@acme.inc"},
			},
			expectedError:   "<nil>",
			expectedStatus:  "email_bounced",
			expectedBounces: 1,
		},
		{
			email: email{
				status:    "email_failed",
				recipient: &user{email: "elmer@acme.inc"},
			},
			expectedError:   "error tracking user bounce: invalid event: email_failed",
			expectedStatus:  "email_failed",
			expectedBounces: 0,
		},
		{
			email: email{
				status:    "email_sent",
				recipient: &user{email: "daffy@acme.inc"},
			},
			expectedError:   "error updating user status: invalid status: email_sent",
			expectedStatus:  "",
			expectedBounces: 0,
		},
		{
			email: email{
				status:    "email_failed",
				recipient: &user{email: "porky@acme.inc"},
			},
			expectedError:   "error tracking user bounce: invalid event: email_failed",
			expectedStatus:  "email_failed",
			expectedBounces: 0,
		},
	}
	for _, tt := range tests {
		a := &analytics{}
		err := a.handleEmailBounce(tt.email)
		fmt.Printf("handleEmailBounce(%+v) = %v\n", tt.email, err)
		fmt.Printf(`
					Expected status: %s, got status: %s
					Expected bounces: %d, got bounces: %d
					Expected error: %v, got error: %v
				`,
			tt.expectedStatus, tt.email.recipient.status, tt.expectedBounces, a.totalBounces, tt.expectedError, err)
		if err != nil && err.Error() != tt.expectedError {
			t.Errorf("handleEmailBounce() error = %v, wantErr %v", err, tt.expectedError)
		}
		if tt.email.recipient.status != tt.expectedStatus {
			t.Errorf("handleEmailBounce() status = %v, want %v", tt.email.recipient.status, tt.expectedStatus)
		}
		if a.totalBounces != tt.expectedBounces {
			t.Errorf("handleEmailBounce() totalBounces = %v, want %v", a.totalBounces, tt.expectedBounces)
		}
	}
}

func TestGetEmailStatusName(t *testing.T) {
	tests := []struct {
		status   emailStatus
		expected string
	}{
		{EmailBounced, "EmailBounced"},
		{EmailInvalid, "EmailInvalid"},
		{EmailDelivered, "EmailDelivered"},
		{EmailOpened, "EmailOpened"},
		{17, "Unknown"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("status=%v", tt.status), func(t *testing.T) {
			got := getEmailStatusName(tt.status)
			if got != tt.expected {
				t.Errorf("getEmailStatusName() = %v, want %v", got, tt.expected)
			}
		})
	}
}
