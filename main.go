package main
import "fmt"
import "strings"

func main(){
	conferenceName := "Go Conference"
	const totalTickets uint = 50
	var remainingTickets uint = 50
	
	println("Welcome to",conferenceName,"the platform")
	println("Get you tickets here!")
	
	fmt.Printf("Total number of tickets %v available tickets %v \n", 
	totalTickets, remainingTickets)

	users := []string{}
	var userName string
	var userTickets uint
	fmt.Println("Welcome Buy your Tickets here!!!!!")
	top:
	for{
		fmt.Print("Enter your name:")
		fmt.Scan(&userName)
		fmt.Print("Enter no of Tickets:")
		fmt.Scan(&userTickets)

		if (userTickets <= remainingTickets){
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

		if (flag == string('n')){
			fmt.Println("Booked list:")
			fmt.Println(strings.Join(users[:], ","))
			break top
		}
	}
}