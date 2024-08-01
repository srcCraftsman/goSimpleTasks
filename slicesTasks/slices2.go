package slicesTasks

// 2. Поиск максимального элемента в слайсе
//	  Напишите программу, которая находит максимальный элемент в заданном слайсе.

import (
	"fmt"
)

func slice2() {

	var sliceNum []int
	var inputCount, inputNum int // inputCount = N; inputNum = numbers added in slice
	var max int                  // maximum number in slice
	fmt.Println("Set N:")
	fmt.Scan(&inputCount)
	fmt.Println("\nEnter the numbers:")

	// Added numbers in sliceNum

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

	// Print result

	fmt.Printf("\nMaximum number: %d", max)
}
