// Package controllers contains the contact list controllers
package controllers

import (
	"bufio"
	"github.com/Medinolas/Adressbuch/models"
	"github.com/Medinolas/Adressbuch/views"
	"log"
	"os"
	"strings"
)

// Run does the running of the console application
func Run(enablePersistence bool) {
	if enablePersistence {
		models.EnableFilePersistence()
	} else {
		models.DisableFilePersistence()
	}

	models.Initialize()

	views.Clear()
	views.PrintMenu()

	for true {
		executeCommand()
	}
}

func executeCommand() {
	command := askForInput()
	parseCommand(command)
}

func askForInput() string {
	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	response = strings.TrimSpace(response)
	return response
}

func parseCommand(input string) {
	switch {
	case input == "1":
		// Add another contact
		//
		views.Clear()
		views.PrintContactInformation()
		response := askForInput()
		contact := createContact(response)
		models.AddContact(contact)
		views.Clear()
		views.PrintMenu()
		break
	case input == "2":
		// Remove a contact
		//
		views.Clear()
		contacts := models.FindAllContacts()
		views.PrintContactList(contacts)
		views.PrintRemoveInformation()
		response := askForInput()
		models.RemoveContact(response)
		views.Clear()
		views.PrintMenu()
		break
	case input == "3":
		// View all contacts
		//
		views.Clear()
		contacts := models.FindAllContacts()
		views.PrintContactList(contacts)
		break
	case input == "4":
		// Edit a contact
		//
		views.Clear()
		contacts := models.FindAllContacts()
		views.PrintContactList(contacts)
		views.EditContactInformation(contacts)
		responseRowIndex := askForInput()
		models.EditContact(responseRowIndex)
		views.EditFieldContact()
		responseField := askForInput()
		views.NewFieldContact()
		responseNewInfo := askForInput()
		models.EditFieldContact(responseRowIndex, responseField, responseNewInfo)
		views.PrintContinue()
		break
	case input == "c":
		// Clear view and print menu
		//
		views.Clear()
		views.PrintMenu()
		break
	case input == "q":
		// Terminate application
		//
		views.Clear()
		views.PrintGoodbye()
		views.ShutDown()
		break
	}
}

func createContact(input string) models.Contact {
	inputFormatted := strings.ReplaceAll(input, ", ", ",")
	elements := strings.Split(inputFormatted, ",")

	if len(elements) != 5 {
		log.Fatal("Input not as expected. 5 comma separated values needed!")
	}

	// Parse contact
	//
	name := elements[0]
	streetName := elements[1]
	houseNumber := elements[2]
	phoneNumber := elements[3]
	eMailAdress := elements[4]

	// Create new contact based on parsed values
	//
	contact := models.Contact{Name: name, StreetName: streetName, HouseNumber: houseNumber, PhoneNumber: phoneNumber, EMailAdress: eMailAdress}

	return contact
}
