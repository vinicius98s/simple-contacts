package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// FILE : the file name where contacts will be saved
const FILE string = "contacs.txt"

// StringConverter : interface that specify how an item can be converted to a string
type StringConverter interface {
	toString() string
}

// Contact : a structure that represent a contact
type Contact struct {
	name         string
	contactForm  string
	contactValue string
}

func (contact *Contact) toString() string {
	return fmt.Sprintf("%s|%s|%s \n", contact.name, contact.contactForm, contact.contactValue)
}

// ContactsManager : responsible for managing contacts
type ContactsManager struct{}

func (manager *ContactsManager) loadContacts() ([]Contact, error) {
	contacts := make([]Contact, 0)
	if _, e := os.Stat(FILE); !os.IsNotExist(e) {
		file, err := os.Open(FILE)
		if err != nil {
			return contacts, err
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			contactRow := scanner.Text()
			contactInfo := strings.Split(contactRow, "|")
			contacts = append(contacts, Contact{name: contactInfo[0], contactForm: contactInfo[1], contactValue: contactInfo[2]})
		}
	}
	return contacts, nil
}

func (manager *ContactsManager) saveContact(contact StringConverter) error {
	file, err := os.OpenFile(FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, e := file.WriteString(contact.toString()); e != nil {
		return e
	}
	return nil
}

func main() {
	manager := ContactsManager{}
	option := 0
	for true {
		fmt.Print("\nHello! What do you want to do?\n\n")
		fmt.Println("1 - List contacts")
		fmt.Println("2 - Create new contact")
		fmt.Printf("3 - Exit\n\n")
		fmt.Print("Choose your option and press Enter: ")
		fmt.Scanf("%d", &option)
		fmt.Print("\n\n")
		switch option {
		case 1:
			listContacts(&manager)
		case 2:
			createNewContact(&manager)
		}
		if option == 3 {
			break
		}
	}

	fmt.Println("Thanks for using :)")
}

func listContacts(manager *ContactsManager) {
	contacts, err := manager.loadContacts()
	if err != nil {
		fmt.Printf("Failed to load contacts: %s \n", err)
	} else {
		fmt.Print("---------")
		fmt.Print("\n\nContacts list: \n")
		for _, contact := range contacts {
			fmt.Printf("\n - %s, %s: %s \n", contact.name, contact.contactForm, contact.contactValue)
		}
		fmt.Print("\n---------\n")
	}
}

func createNewContact(manager *ContactsManager) {
	newContact := Contact{}
	fmt.Print("Contact name: ")
	fmt.Scanf("%s", &newContact.name)
	fmt.Print("Contact form: ")
	fmt.Scanf("%s", &newContact.contactForm)
	fmt.Print("Contact: ")
	fmt.Scanf("%s", &newContact.contactValue)

	err := manager.saveContact(&newContact)
	if err != nil {
		fmt.Printf("Failed to save contact: %s \n", err)
	}
}
