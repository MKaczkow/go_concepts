package utils

import (
	"fmt"
	"strings"
)


func GreetUsers(conference_name string, conference_tickets uint, remaining_tickets uint) {
	fmt.Printf("Welcome to %v booking application\n", conference_name)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conference_tickets, remaining_tickets)
	fmt.Println("Get your tickets here to attend")
}


func GetFirstNames(bookings []string) []string{
	first_names := []string{}

	for _, booking := range bookings {
		var names = strings.Fields(booking)
		first_names = append(first_names, names[0])
	}

	return first_names
}


func GetUserInput() (string, string, string, uint) {
	var user_first_name string
	var user_last_name string
	var email string
	var user_tickets uint

	fmt.Println("Enter Your First Name: ")
	fmt.Scanln(&user_first_name)

	fmt.Println("Enter Your Last Name: ")
	fmt.Scanln(&user_last_name)

	fmt.Println("Enter Your Email: ")
	fmt.Scanln(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scanln(&user_tickets)

	return user_first_name, user_last_name, email, user_tickets
}


func IsUserInputValid(user_first_name string, user_last_name string, email string, user_tickets uint, remaining_tickets uint) (bool, bool, bool) {
	name_is_valid := len(user_first_name) >= 2 && len(user_last_name) >= 2
	email_is_valid := strings.Contains(email, "@")
	num_tickets_is_valid := user_tickets > 0 && user_tickets <= remaining_tickets

	return name_is_valid, email_is_valid, num_tickets_is_valid
}


func BookTicket(user_tickets uint, user_first_name string, user_last_name string, 
	email string, remaining_tickets uint, bookings []string, conference_name string) {

	remaining_tickets = remaining_tickets - user_tickets
	bookings = append(bookings, user_first_name + " " + user_last_name)

	fmt.Printf("Thank you %v %v for booking %v tickets. \nYou will receive confirmation via email at %v\n", user_first_name, user_last_name, user_tickets, email)
	fmt.Printf("%v tickets remainig for %v\n", remaining_tickets, conference_name)
	
	fmt.Printf("These are all our bookings: %v\n", GetFirstNames(bookings))
}
