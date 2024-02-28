package contacts

import (
	"fmt"
	"strconv"
)

type Contact struct {
	ID    string
	Name  string
	Email string
}

type Contacts []Contact

func NewContacts(start int, num int) Contacts {
	var contacts Contacts
	for i := start; i <= num; i++ {
		contacts = append(contacts, Contact{
			ID:    strconv.Itoa(i),
			Name:  fmt.Sprintf("Agent Smith %d", i),
			Email: fmt.Sprintf("agent.smith%d@example.com", i),
		})
	}
	return contacts
}
