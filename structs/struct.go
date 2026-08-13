package structs

// Nested structs
type MessageToSend struct {
	Message   string
	Sender    User
	Recipient User
}

type User struct {
	Name   string
	Number int
}

func CanSendMessage(mToSend MessageToSend) bool {
	if mToSend.Sender.Name == "" || mToSend.Sender.Number == 0 {
		return false
	}
	if mToSend.Recipient.Name == "" || mToSend.Recipient.Number == 0 {
		return false
	}
	return true
}
