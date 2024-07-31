package player

import "fmt"

type player struct {
  name string
	level int
}
func (player) Test () string{
	return "TEST"
}


func New(name string, level int) player {
	return player{
		name: name,
		level: level,
	}
}

func (p player) ToString(){
	fmt.Println("Name:", p.name)
	fmt.Println("Level:", p.level)
}