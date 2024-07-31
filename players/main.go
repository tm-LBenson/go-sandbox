package main

import (
	"fmt"
	"github.com/tm-lbenson/go-sandbox/players/player"
)
func main(){

	playerOne := player.New("Dave", 1)
	playerOne.ToString()
	fmt.Println(playerOne.Test())
}