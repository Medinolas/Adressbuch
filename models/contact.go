// Package models contains the contact library models
package models

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

const FileName = "data.csv"

// Contact repository
var contacts []Contact

// Contact persistence
var filePersistence bool = false

// Contact as type
type Contact struct {
	Name        string
	StreetName  string
	HouseNumber string
	PostalCode  string
	CityName    string
	PhoneNumber string
	EMailAdress string
}

// EnableFilePersistence enables the file persistence
func EnableFilePersistence() {
	filePersistence = true
}

// DisableFilePersistence disables the file persistence
func DisableFilePersistence() {
	filePersistence = false
}

// Initialize does the initialization of the repository
func Initialize() {
	if filePersistence {
		contacts, _ = getDataFromFile()
	}
}

func getDataFromFile() ([]Contact, error) {
	// open file
	//
	file, err := os.Open(FileName)
	if err != nil {
		return nil, err
	}

	var readContacts []Contact

	// read csv values using csv.Reader
	//
	csvReader := csv.NewReader(file)
	for {
		records, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		contact := parseContactData(records)

		// Add contact to slice
		//
		readContacts = append(readContacts, contact)
	}

	// remember to close the file at the end
	//
	file.Close()

	return readContacts, nil
}

func parseContactData(rec []string) Contact {
	// Parse contact
	//
	name := rec[0]
	streetName := rec[1]
	houseNumber := rec[2]
	postalCode := rec[3]
	cityName := rec[4]
	phoneNumber := rec[5]
	eMailAdress := rec[6]

	// Create new contact based on parsed values
	//
	contact := Contact{Name: name, StreetName: streetName, HouseNumber: houseNumber, PostalCode: postalCode, CityName: cityName, PhoneNumber: phoneNumber, EMailAdress: eMailAdress}
	return contact
}

func updateDataInFile() {
	file, err := os.OpenFile(FileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	checkError("Cannot open file", err)
	writer := csv.NewWriter(file)

	for _, contact := range contacts {
		Contactserialized := []string{contact.Name, contact.StreetName, contact.HouseNumber, contact.PostalCode, contact.CityName, contact.PhoneNumber, contact.EMailAdress}
		err := writer.Write(Contactserialized)
		checkError("Cannot write to file", err)
	}

	writer.Flush()
	file.Close()
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

// FindAllContacts returns a copy of all Contacts
func FindAllContacts() []Contact {
	allContacts := make([]Contact, len(contacts))
	copy(allContacts, contacts)
	return allContacts
}

// ToInt converts a string to an integer value
func ToInt(info string) int {
	aInt, _ := strconv.Atoi(info)
	return aInt
}

// AddContact adds a contact to the list
func AddContact(contact Contact) bool {
	if contact.Name == "" || contact.StreetName == "" || contact.HouseNumber == "" || contact.PostalCode == "" || contact.CityName == "" || contact.PhoneNumber == "" || contact.EMailAdress == "" {
		return false
	}

	contacts = append(contacts, contact)

	if filePersistence {
		updateDataInFile()
	}

	return true
}

// RemoveContact removes a contact from the list
func RemoveContact(rowIndex string) {
	var tempContacts []Contact
	idAsInt := ToInt(rowIndex) - 1

	for index, currentContact := range contacts {
		if index != idAsInt {
			tempContacts = append(tempContacts, currentContact)
		}
	}

	contacts = tempContacts

	if filePersistence {
		updateDataInFile()
	}
}

// EditContact chooses the contact you want to edit
func EditContact(i string) {
	idAsInt := ToInt(i) - 1

	for index, currentContact := range contacts {
		if index == idAsInt {
			fmt.Println(currentContact)
		}
	}

}

// EditFieldContact changes a field in the contact information
func EditFieldContact(rowIndex, field, newInfo string) {
	idAsInt := ToInt(rowIndex) - 1

	for index, _ := range contacts {
		if index == idAsInt {
			switch {
			case field == "name":
				contacts[index].Name = newInfo
			case field == "adress":
				contacts[index].StreetName = newInfo
			case field == "house number":
				contacts[index].HouseNumber = newInfo
			case field == "postal code":
				contacts[index].PostalCode = newInfo
			case field == "city name":
				contacts[index].CityName = newInfo
			case field == "phone number":
				contacts[index].PhoneNumber = newInfo
			case field == "e-mail adress":
				contacts[index].EMailAdress = newInfo
			}
		}
	}
	if filePersistence {
		updateDataInFile()
	}
	fmt.Println("Contact updated.")
}
