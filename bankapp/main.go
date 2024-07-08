package main

import (
	"fmt"
	"strconv"
	"github.com/tm-lbenson/go-sandbox/bankapp/bank"
)

func main(){
	fmt.Print("\033[H\033[2J")
	fmt.Println("Welcome to the bank!")
	displayMenu()
	for{
	balance := bank.GetBalance()
	choice := getUserChoice()
	switch(choice){
	case 1:
		balanceAsString := strconv.FormatFloat(bank.GetBalance(), 'f', -1, 64)
		fmt.Println("Your current balance is $"+balanceAsString)
	case 2:
		amount := 0.0
		fmt.Print("How much would you like to deposit? ")
		fmt.Scan(&amount)
		balance = bank.Deposit(amount, balance)
	case 3:
		fmt.Print("How much would you like to withdraw? ")
		amount := 0.0
		fmt.Scan(&amount)
		balance = bank.Withdraw(amount, balance)
	case 4:
		fmt.Println("Goodbye")
		return
	default:
		fmt.Println("Invalid")

	}
	bank.SetBalance(balance)
}
}

func displayMenu(){
	fmt.Println("1. Balance")
	fmt.Println("2. Deposit")
	fmt.Println("3. Withdraw")
	fmt.Println("4. Exit")
}

func getUserChoice() (choice int){
	fmt.Print("Choose an option (1-4): ")
	fmt.Scan(&choice)
	return choice
}


