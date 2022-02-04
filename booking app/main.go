//declaring main pacakage for excuting

package main

// importing pacakages
import (
	"fmt"
	"strings"
)

func main() {
	// Variable & constants including explicit declaration for types
	for {
		var conferanceName = "Go refrence"
		var reaminingTickets uint = 50
		var userName string
		const conferanceTickets = 50

		// fmt package for text output (Println & Printf diffrence)
		fmt.Printf("conferanceTickets is an %T , reaminingTickets is an %T \n", conferanceTickets, reaminingTickets)
		fmt.Println("")
		fmt.Printf("Hello welcome to our conferance %v \n", conferanceName)
		fmt.Println("To join , please book your ticket")
		fmt.Println("Note there is only ", reaminingTickets, "Tickets Avaialble")

		//implicite assignements
		userName = "Abdo"

		// User Input with fmt package
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		fmt.Print("Please enter your username :")
		fmt.Scan(&userName)

		fmt.Print("Please enter your firstname :")
		fmt.Scan(&firstName)

		fmt.Print("Please enter your last name :")
		fmt.Scan(&lastName)

		fmt.Print("Please enter your email :")
		fmt.Scan(&email)
		for {
			fmt.Println("How much tickets you want to book ? :")
			fmt.Scan(&userTickets)
			if userTickets > reaminingTickets {
				fmt.Println("The amount of tickets is not avaialble")
				fmt.Printf("Only %v is availabe \n", reaminingTickets)
				continue
			} else {
				break
			}
		}

		reaminingTickets = reaminingTickets - userTickets

		fmt.Printf("Hello %v , You have book %v tickets , your will recive an email an email in a short time at : %v ", userName, userTickets, email)
		fmt.Printf("This many tickets remaning for this conferance : %v \n", reaminingTickets)

		// Arrays & slices
		var bookingClients [50]string

		// Initialize the first element of the arry with the user input and priting it
		fmt.Println("")

		bookingClients[0] = firstName + " " + lastName

		fmt.Printf("Type of bookingClients is %T\n", bookingClients)
		fmt.Printf("The first element is %v \n", bookingClients[0])
		fmt.Printf("The array length is %v \n", len(bookingClients))

		// Slices is an abstraction of an arry , but has a dynamic size ;)
		// Slices is genearaly better than arrays
		//Defining an array
		fmt.Println("")

		var SliceBooking []string
		SliceBooking = append(SliceBooking, firstName+" "+lastName)

		fmt.Printf("Type of SliceBooking is %T\n", SliceBooking)
		fmt.Printf("First slice element %v", SliceBooking[0])
		fmt.Printf("The Slice length is %v \n", len(SliceBooking))

		var firstNames []string
		for _, booking := range SliceBooking {
			name := strings.Fields(booking)
			firstName := name[0]
			firstNames = append(firstNames, firstName)
		}

		fmt.Println("Usernames  : ", firstNames)

		noUserTicketReamning := reaminingTickets == 0
		if noUserTicketReamning {
			// end of program
			fmt.Println("Thank you , our conferance is sold out")
			break
		} else {
			fmt.Printf("This many tickets remaning for this conferance : %v \n", reaminingTickets)
			continue
		}
	}
}
