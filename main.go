package main
import "fmt"

func main(){
	conferenceName := "Go Conference"
	const totalTickets int = 50
	var remainingTickets int = 50
	
	println("Welcome to",conferenceName,"the platform")
	println("Get you tickets here!")
	
	fmt.Printf("Total number of tickets %v available tickets %v \n", 
	totalTickets, remainingTickets)

	var userName string
	var userTickets int

	fmt.Print("Enter your name:")
	fmt.Scan(&userName)
	fmt.Print("Enter no of Tickets:")
	fmt.Scan(&userTickets)

	if (userTickets <= remainingTickets){
		remainingTickets -= userTickets
		fmt.Println("Successfull purchase!!")
		fmt.Printf("\nUser %v booked %v tickets\n", userName, userTickets)
	} else
	{
		fmt.Printf("\nUser %v cannot book more tickets than available tickets: %v\n", 
		userName, remainingTickets)
	}
}