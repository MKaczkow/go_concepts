package main

import (
	"fmt"
	"strings"
)


func main () {
	conference_name := "Go Conference"
	// above line is syntactic sugar equivalent of 
	// var conference_name = "Go Conference"
	const conference_tickets uint = 50
	var remaining_tickets uint = 50

	fmt.Printf("conference_tickets is %T, remaining_tickets is %T, conference_name is %T\n", conference_tickets, remaining_tickets, conference_name)

	fmt.Printf("Welcome to %v booking application\n", conference_name)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conference_tickets, remaining_tickets)
	fmt.Println("Get your tickets here to attend")

	var user_first_name string
	var user_last_name string
	var email string
	var user_tickets uint
	// use Slice (abstraction of array) insted of Array datatype for dynamic memory allocation
	var bookings [] string   // <-- Slice
	// var bookings [50] string <-- Array

	for {
		fmt.Println("Enter your first name: ")
		fmt.Scan(&user_first_name)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&user_last_name)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter how many tickets you would like to buy: ")
		fmt.Scan(&user_tickets)
		
		if user_tickets > remaining_tickets {
			fmt.Printf("We only have %v tickets remaining\n", remaining_tickets)
			continue
		}

		remaining_tickets = remaining_tickets - user_tickets
		bookings = append(bookings, user_first_name + " " + user_last_name)

		fmt.Printf("Thank you %v %v for booking %v tickets. \nYou will receive confirmation via email at %v\n", user_first_name, user_last_name, user_tickets, email)
		fmt.Printf("%v tickets remainig for %v\n", remaining_tickets, conference_name)
		
		first_names := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			first_names = append(first_names, names[0])
		}

		fmt.Printf("These are all our bookings: %v\n", first_names)

		if remaining_tickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year")
			break
		}
	}
}
