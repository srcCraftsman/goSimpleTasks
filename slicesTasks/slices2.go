package slicesTasks

// 2. Поиск максимального элемента в слайсе
//	  Напишите программу, которая находит максимальный элемент в заданном слайсе.

import (
	"fmt"
)

func slice2() {

	var sliceNum []int = SliceGen()

	var max int // maximum number in slice

	// Search for maximum num in sliceNum

	for _, num := range sliceNum {

		if num >= max {
			max = num
		}
	}

	// Print result

	fmt.Printf("\nMaximum number: %d", max)
}
