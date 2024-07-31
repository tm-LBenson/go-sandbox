package main

import (
	"fmt"

)
type User struct{
	name string
	age int
	email string
	isAdmin bool
}
func main(){
	user1 := User{
		name:"Fred",
		age:18,
		email:"none",
		isAdmin:false,
	}

	fmt.Println(user1)	
}