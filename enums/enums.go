package enums

import "fmt"

// lack of enums in Go, we can use constants and iota to create a set of related values. Here's an example of how to define an enum-like type for email statuses:
func (a *analytics) handleEmailBounce(em email) error {
	err := em.recipient.updateStatus(em.status)
	if err != nil {
		return fmt.Errorf("error updating user status: %w", err)
	}
	err = a.track(em.status)
	if err != nil {
		return fmt.Errorf("error tracking user bounce: %w", err)
	}
	return nil
}

// iota is used to create a set of related constants. Here's an example of defining email statuses using iota:
type emailStatus int

const (
	EmailBounced emailStatus = iota
	EmailInvalid
	EmailDelivered
	EmailOpened
)

func getEmailStatusName(status emailStatus) string {
	switch status {
	case EmailBounced:
		return "EmailBounced"
	case EmailInvalid:
		return "EmailInvalid"
	case EmailDelivered:
		return "EmailDelivered"
	case EmailOpened:
		return "EmailOpened"
	default:
		return "Unknown"
	}
}
