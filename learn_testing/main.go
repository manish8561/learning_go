package main

import (
	"fmt"

	"example.com/demo/player"
)

func main() {
	// res := math.Add(1,2)

	// fmt.Println("result: ", res)

	s := player.Stats{Name: "player 1", Minutes: 25.1, Points: 21, Assits: 3, TrunOvers: 7, Rebounds: 10}

	fmt.Println(player.HadAGoodGame(&s))
}
