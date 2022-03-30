package main

import (
	"fmt"
	"math"
)

func main() {
	//slices template
	room1 := []float64{20, 21, 23, 25, 22}
	room2 := []float64{27, 23, 25, 20, 30, 24}
	room3 := []float64{22, 23, 24, 22}

	//Insert code here
	var allRoom [][]float64
	var average []float64
	var total float64
	allRoom = append(allRoom, room1, room2, room3)
	// fmt.Println(allRoom)

	for _, room := range allRoom {
		total = 0
		for _, temp := range room {
			total += temp
		}
		// fmt.Println(total)
		average = append(average, math.Round(total*100/float64(len(room)))/100)
	}

	fmt.Println(average)

}
