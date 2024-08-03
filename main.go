package main
import "fmt"
import "strings"

func main(){
	conferenceName := "Go Conference"
	const totalTickets uint = 50
	var remainingTickets uint = 50
	
	greetUser(conferenceName, totalTickets, remainingTickets)

	users := []string{}
	users, remainingTickets  = bookingLogic(remainingTickets, users)
	fmt.Printf("Remaining Tickets after sale: %v\n", remainingTickets)
	fmt.Printf("Booked List: %v", strings.Join(users, "\n"))
}


func greetUser(conferenceName string, totalTickets uint, remainingTickets uint) {
	
	println("Welcome to",conferenceName,"the platform")
	println("Get you tickets here!")
	
	fmt.Printf("Total number of tickets %v available tickets %v \n", 
	totalTickets, remainingTickets)
}

func bookingLogic(remainingTickets uint, users []string) ([]string, uint) {
	
	var userName string
	var userTickets uint
	top:
	for{
		fmt.Println("Welcome Buy your Tickets here!!!!!")
		fmt.Print("Enter your name:")
		fmt.Scan(&userName)
		fmt.Print("Enter no of Tickets:")
		fmt.Scan(&userTickets)
		isValidTicketNumber := userTickets > 0

		if remainingTickets <=0{
			fmt.Println("Fully booked!!!!!!!!!!!!!!")
			fmt.Println("EXITING")
			break
		} else if (isValidTicketNumber && userTickets <= remainingTickets){
			remainingTickets -= userTickets
			fmt.Println("Successfull purchase!!")
			fmt.Printf("\nUser %v booked %v tickets\n", userName, userTickets)
			users = append(users, userName)
		} else
		{
			fmt.Printf("\nUser %v cannot book more tickets than available tickets: %v\n", 
			userName, remainingTickets)
		}
		var flag string
		fmt.Print("Do you want to continue:(y/n)")
		fmt.Scan(&flag)

		if (flag == string('n') || flag == string('N')){
			break top
		} else if (flag == string('y') || flag == string('Y')){
			fmt.Println("Make next booking --------------------->")
		} else {
			fmt.Println("Bad input")
			fmt.Println("Make next Booking -------------------->")
		}
	}
	return users, remainingTickets
}