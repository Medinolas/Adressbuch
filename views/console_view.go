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
	#  4. EDIT CONTACT					            #
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
	for i, contact := range contactsToPrint {
		fmt.Println(i+1, "| Name:", contact.Name+",", "Adress:", contact.StreetName+",", "House Number:",
			contact.HouseNumber+",", "Phone Number:", contact.PhoneNumber+",", "E-Mail Adress:", contact.EMailAdress+",")
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

func PrintContactInformation() {
	fmt.Println("Please enter the new contact with the following format: Full name, Adress, House Number, Phone Number, E-Mail Adress")
}

func PrintRemoveInformation() {
	fmt.Println("Please enter the name of the contact you want to remove.")
}

func EditContactInformation([]models.Contact) {
	fmt.Println("Please enter the index number of the contact you want to edit.")
}

func EditFieldContact() {
	fmt.Println("Do you want to change name, adress, house number, phone number or e-mail adress?")
}

func NewFieldContact() {
	fmt.Println("Please enter new information.")
}
