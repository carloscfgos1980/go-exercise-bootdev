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

// Structs methods
type authenticationInfo struct {
	username string
	password string
}

func (auth authenticationInfo) GetBasicAuth() string {
	return "Authorization: Basic " + auth.username + ":" + auth.password
}

// Update Users
type User1 struct {
	Name string
	Membership
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

func NewUser(name string, membershipType string) User1 {
	if membershipType == "premium" {
		return User1{
			Name: name,
			Membership: Membership{
				Type:             membershipType,
				MessageCharLimit: 1000,
			},
		}
	}
	return User1{
		Name: name,
		Membership: Membership{
			Type:             membershipType,
			MessageCharLimit: 100,
		},
	}
}
