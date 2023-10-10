package main

import (
	"booking-application/validation"
	"fmt"
	"sync"
	"time"
)

const eventTickets int = 50

var eventName = "Irvine event"
var remainingTickets uint = 50

// create an empty list of map
// create a slice with make
// initial 0 because it will increase
var bookings = make([]UserData, 0)

// using struct for mixed data type
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// waits for the launched goroutine to finish
// package "sync" provides basic synchronization functionality
var wg = sync.WaitGroup{}

func main() {

	greeting()

	firstName, lastName, email, userTickets := getUserInput()

	isValidUserName, isValidEmail, isValidTicketNumber := validation.ValidUserInputs(firstName, lastName, email, userTickets, remainingTickets)

	if isValidUserName && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, email)
		// Add: sets the number of goroutines to wait for (increases the counter by the provided number)
		wg.Add(1)
		// with the go keyword, it starts a new goroutine
		// a goroutine is a lightweight thread managed by the Go runtime
		// running seperate threads in the background
		// Go is using "Green Thread", it is an Abstraction of an actual thread
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		noTicketsRemaining := remainingTickets == 0
		if noTicketsRemaining {
			// end the application
			fmt.Println("The event is fully booked, please come back next year!")
		}
	} else {
		if !isValidUserName {
			fmt.Printf("Your first name or last name is too short\n")
		}
		if !isValidEmail {
			fmt.Printf("Your email address doen not contain @ sign\n")
		}
		if !isValidTicketNumber {
			fmt.Printf("Your number of tickets is invalid\n")
		}
	}
	// Blocks until the WaitGroup counter is 0
	wg.Wait()
}

func greeting() {
	fmt.Printf("This is the booking application of %v.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", eventName, eventTickets, remainingTickets)
}

func getFirstNames() []string {
	firstNames := []string{}
	// ignore a variable of index
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// user input
	fmt.Println("Enter Your First Name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter Your Last Name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter Your Email: ")
	fmt.Scanln(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scanln(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// // using map
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. The booking confirmation email will be sent to %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("There are %v tickets remaining for %v\n", remainingTickets, eventName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	// sleep for 30 secs
	// stop or block the current thread
	// block the further processing for 30 seconds
	time.Sleep(30 * time.Second)
	var ticket = fmt.Sprintf("There are %v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Right now sending ticket:\n %v \nto email address: %v\n", ticket, email)
	fmt.Println("#################")
	// Decrements the WaitGroup counter by 1, it is called by the goroutine to indicate that it's finished
	wg.Done()
}
