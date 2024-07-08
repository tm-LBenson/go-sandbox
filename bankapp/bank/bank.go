package bank

import (
	"fmt"
	"os"
	"strconv"
)

func GetBalance() float64{
	data, readErr:= os.ReadFile("balance.txt")
	balanceString := string(data)
	balance, parseErr :=strconv.ParseFloat(balanceString, 64)
	if readErr != nil || parseErr != nil{
		return 1000.0
	}
	return balance
}

func SetBalance(amount float64){
	// data := strconv.FormatFloat(amount, 'f', -1, 64)
	data := fmt.Sprintf("%.1f", amount)
	os.WriteFile("balance.txt", []byte(data), 0644)
}




func Deposit(amount, balance float64) float64{
	newBalance := balance + amount
	if amount < 0 {
		fmt.Println("Invalid")
		return balance
	}else{
		SetBalance(newBalance)
		return newBalance
	}
}

func Withdraw(amount, balance float64) float64{
		newBalance := balance - amount
	if amount < 0 || newBalance <= 0{
		fmt.Println("Invalid")
		return balance
	}else{
		SetBalance(newBalance)
		return newBalance
	}
}