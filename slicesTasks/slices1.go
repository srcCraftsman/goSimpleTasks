package slicesTasks

// 1. Создание слайса
//	  Напишите программу, которая создает слайс из первых N целых чисел и выводит его.

import (
	"fmt"
)

func slice1() {

	var sliceNum []int
	var inputCount, inputNum int // inputCount = N; inputNum = numbers added in slice

	fmt.Println("\nSet N:")
	fmt.Scan(&inputCount)
	fmt.Println("\nEnter the numbers:")

	// Add numbers in sliceNum

	for i := 0; i < inputCount; i++ {
		fmt.Scan(&inputNum)
		sliceNum = append(sliceNum, inputNum)

	}

	// Print result

	fmt.Println("\nSlice: ", sliceNum)
}
