package main

import (
	"fmt"

	"github.com/tm-lbenson/go-sandbox/players/player"
)
func main(){

	playerOne := player.Player{
		Name: "Dave",
		Level: 1,
	}
	fmt.Println(playerOne)
}