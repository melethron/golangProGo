package main

import (
	"fmt"
	"math"
)

func converting() {
	fmt.Println("Converting")
	kayak := 275
	soccerBall := 19.50
	total := soccerBall + float64(kayak)
	fmt.Println("Int to float:", total)
	totalFloatToInt := kayak + int(soccerBall)
	fmt.Println("Float to int:", totalFloatToInt)
	fmt.Println("Int to int8:", int8(totalFloatToInt))
	totalRound := kayak + int(math.Round(soccerBall))
	fmt.Println("float64 math.Round:", totalRound)
	fmt.Println()
}
