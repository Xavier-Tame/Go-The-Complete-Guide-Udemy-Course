package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountbalanceFile = "balance.txt"

func main() {
	accountBalance, err := fileops.GetFloatFromFile(accountbalanceFile)
	if err != nil {
		fmt.Println(err)
		fmt.Println("--Starting with default balance of $1000.00--")
		// panic("Can't read balance from file")
	}

	var choice int
	fmt.Println("Hello, Welcome to Go Bank!")
	fmt.Println("Reach us 24/7 ", randomdata.PhoneNumber())

	for choice != 4 {

		presentOptions()

		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your account balance is: $%.2f\n", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Enter amount to deposit: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Deposit amount must be positive.")
				continue
			}
			accountBalance += depositAmount
			fmt.Printf("Deposited $%.2f. New balance is: $%.2f\n", depositAmount, accountBalance)
		case 3:
			var withdrawAmount float64
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient funds.")
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Printf("Withdrew $%.2f. New balance is: $%.2f\n", withdrawAmount, accountBalance)
		default:
			if choice != 4 {
				fmt.Println("Invalid choice, please try again.")
				continue
			}
			fmt.Println("Thank you for banking with us. Goodbye!")
			fileops.WriteFloatToFile(accountBalance, accountbalanceFile)
		}

	}
	// fmt.Println("Your choice: ", choice)
}
