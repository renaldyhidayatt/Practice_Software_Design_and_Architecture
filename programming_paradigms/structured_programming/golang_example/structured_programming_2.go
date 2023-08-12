package golang_example

import (
	"fmt"
)

type Contact struct {
	name        string
	phoneNumber string
}

func (c Contact) String() string {
	return "Name: " + c.name + ", Phone: " + c.phoneNumber
}

type ContactManager struct {
	contacts []Contact
}

func NewContactManager() *ContactManager {
	return &ContactManager{
		contacts: make([]Contact, 0),
	}
}

func (cm *ContactManager) AddContact(name, phoneNumber string) {
	newContact := Contact{name: name, phoneNumber: phoneNumber}
	cm.contacts = append(cm.contacts, newContact)
	fmt.Println("Contact added:", newContact)
}

func (cm *ContactManager) ListContacts() {
	fmt.Println("Contact List:")
	for _, contact := range cm.contacts {
		fmt.Println(contact)
	}
}

func main() {
	contactManager := NewContactManager()

	for {
		fmt.Println("Menu:")
		fmt.Println("1. Add Contact")
		fmt.Println("2. List Contacts")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var name, phoneNumber string
			fmt.Print("Enter name: ")
			fmt.Scan(&name)
			fmt.Print("Enter phone number: ")
			fmt.Scan(&phoneNumber)
			contactManager.AddContact(name, phoneNumber)
		case 2:
			contactManager.ListContacts()
		case 3:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please choose again.")
		}
	}
}
