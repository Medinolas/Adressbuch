// Package views contains the book library views
package views

import (
	"fmt"
	"github.com/Medinolas/Adressbuch/models"
	"os"
	"os/exec"
)

// PrintMenu prints the menu to the console
func PrintMenu() {
	fmt.Println(`
	#################################################
	#                                               #
	#         WELCOME TO YOUR CONTACT LIST          #
	#          CHOOSE YOUR OPTION BELOW             #
	# --------------------------------------------- #
	#  1. ADD ANOTHER CONTACT IN CONTACT LIST       #
	#  2. REMOVE A CONTACT FROM THE LIST            #
	#  3. VIEW ALL CONTACTS                         #
	#  4. EDIT CONTACT                              #
	#  c. CLEAR VIEW AND PRINT MENU                 #
	#  q. TERMINATE CONTACT LIST APP                #
	#                                               #
	#################################################
	`)
}

// Clear clears the console view
func Clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

// PrintContactList prints all the contacts in the contact list to the console
func PrintContactList(contactsToPrint []models.Contact) {

	fmt.Println("--------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Printf("%-4s | %-20s | %-20s | %-11s | %-11s | %-15s | %-12s | %-25s\n", "ID", "Name", "Adress", "House number", "Postal code", "City name", "Phone number", "E-Mail adress")
	fmt.Println("--------------------------------------------------------------------------------------------------------------------------------------")

	for i, contact := range contactsToPrint {

		fmt.Printf("%-4d | %-20s | %-20s | %-12s | %-11s | %-15s | %-12s | %-25s\n", i+1, contact.Name, contact.StreetName, contact.HouseNumber, contact.PostalCode, contact.CityName, contact.PhoneNumber, contact.EMailAdress)
	}
}

// PrintContinue prints the continuation information to the console
func PrintContinue() (int, error) {
	return fmt.Println("Enter c to go back.")
}

// PrintGoodbye prints a goodbye message to the console
func PrintGoodbye() {
	fmt.Println("Goodbye!")
}

// ShutDown terminates the application
func ShutDown() {
	os.Exit(0)
}

// PrintContactInformation asks you to enter the new contact information
func PrintContactInformation() {
	fmt.Println("Please enter the new contact with the following format: Full name, Adress, House Number, Postal Code, City Name, Phone Number, E-Mail Adress")
}

// PrintRemoveInformation asks for the name of the contact you want to remove from the list
func PrintRemoveInformation() {
	fmt.Println("Please enter the index number of the contact you want to remove.")
}

// EditContactInformation asks which contact you want to edit
func EditContactInformation() {
	fmt.Println("Please enter the index number of the contact you want to edit.")
}

// EditFieldContact asks for the field of the contact information you want to edit
func EditFieldContact() {
	fmt.Println("Do you want to change name, adress, house number, postal code, city name, phone number or e-mail adress?")
}

// NewFieldContact asks for the new information
func NewFieldContact() {
	fmt.Println("Please enter new information.")
}
