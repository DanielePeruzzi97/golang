package main

import (
	"fmt"
	"strings"
)

func main() {
	tickets := 50
	for {
		var firstName string
		var secondName string
		var email string
		var bookedTickets int
		var bookings []string

		fmt.Printf("%v tickets remaining.\n", tickets)

		fmt.Println("Insert your name:")
		fmt.Scan(&firstName)

		fmt.Println("Insert your surname:")
		fmt.Scan(&secondName)

		fmt.Println("Insert your email:")
		fmt.Scan(&email)

		fmt.Println("How many tickets do you want?:")
		fmt.Scan(&bookedTickets)
		tickets = tickets - bookedTickets

		bookings = append(bookings, firstName+" "+secondName+" "+email)

		firstNames := []string{}

		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("Names of the bookings are %v.", firstNames)
		fmt.Printf("Stato prenotazione %v.\n", bookings[0])
		fmt.Printf("Lunghezza slice: %v.\n", len(bookings))
		fmt.Printf("%v tickets remaining.\n", tickets)
	}
}
