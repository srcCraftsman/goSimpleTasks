package slicesTasks

// 3. Поиск минимального элемента в слайсе
//	  Напишите программу, которая находит минимальный элемент в заданном слайсе.

import (
	"fmt"
)

func slice3() {

	var sliceNum []int
	var inputCount, inputNum int // inputCount = N; inputNum = numbers added in slice
	var min, max int             // minimum and maximum number in slice
	fmt.Println("\nSet N:")
	fmt.Scan(&inputCount)
	fmt.Println("\nEnter the numbers:")

	// Add numbers in sliceNum

	for i := 0; i < inputCount; i++ {
		fmt.Scan(&inputNum)
		sliceNum = append(sliceNum, inputNum)

	}

	// Search for maximum num in sliceNum

	for _, num := range sliceNum {

		if num >= max {
			max = num
		}
	}

	// Search for minimum num in sliceNum

	for _, num := range sliceNum {
		if num < max {
			max = num
			min = num
		}
	}

	// Print result

	fmt.Printf("\nMinimum number: %d", min)
}
