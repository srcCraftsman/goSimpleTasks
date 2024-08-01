package slicesTasks

// 3. Поиск минимального элемента в слайсе
//	  Напишите программу, которая находит минимальный элемент в заданном слайсе.

import (
	"fmt"
	"math"
)

func slice3() {

	var sliceNum []int = SliceGen()

	var min int = sliceNum[0] // minimum number in slice

	// Search for minimum num in sliceNum

	for _, num := range sliceNum {
		min = int(math.Min(float64(min), float64(num)))
	}

	// Print result

	fmt.Printf("\nMinimum number: %d", min)
}
