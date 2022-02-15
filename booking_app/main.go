package main

import (
	"fmt"
	"booking_app/utils"
)


func main () {
	conference_name := "Go Conference"
	// above line is syntactic sugar equivalent of 
	// var conference_name = "Go Conference"
	const conference_tickets uint = 50
	var remaining_tickets uint = 50
	
	// use Slice (abstraction of array) insted of Array datatype for dynamic memory allocation
	// var bookings [] string   // <-- Slice
	// var bookings [50] string <-- Array

	// actually, we use map
	// var bookings = make([]map[string]string, 1)
	// actually not - we use custom data type
	var bookings = make([]utils.User_data_struct, 0)



	fmt.Printf("conference_tickets is %T, remaining_tickets is %T, conference_name is %T\n", conference_tickets, remaining_tickets, conference_name)

	utils.GreetUsers(conference_name, conference_tickets, remaining_tickets)

	for {
		user_first_name, user_last_name, email, user_tickets := utils.GetUserInput()

		name_is_valid, email_is_valid, num_tickets_is_valid := utils.IsUserInputValid(
			user_first_name, user_last_name, email, user_tickets, remaining_tickets)
		
		if name_is_valid && email_is_valid && num_tickets_is_valid {

			utils.BookTicket(user_tickets, user_first_name, user_last_name, email, 
				       remaining_tickets, bookings, conference_name)
			// "go" in Golang is ridiculously simple way of creating new thread 
			// and passing current line execution to it
			utils.Wait_group.Add(1)
			go utils.SendTicket(user_tickets, user_first_name, user_last_name, email)
			
			if remaining_tickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year")
				break
			}
		} else { 
			if !name_is_valid {
				fmt.Println("firt name or last name you entered is too short")
			}
			if !email_is_valid {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !num_tickets_is_valid {
				fmt.Println("number of tickets you entered is invalid")
			}
			continue
		}
	}
	utils.Wait_group.Wait()
}
