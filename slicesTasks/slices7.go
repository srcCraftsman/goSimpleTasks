package slicesTasks

// 7. Удаление элемента из слайса
//    Напишите программу, которая удаляет элемент из слайса по заданному индексу.

import (
	"fmt"
)

func slice7() {
	var sliceNum []int
	var inputCount, inputNum int // inputCount = N; inputNum = numbers added in slice
	fmt.Println("Set N:")
	fmt.Scan(&inputCount)
	fmt.Println("\nEnter the numbers:")

	// Add numbers in sliceNum

	for i := 0; i < inputCount; i++ {
		fmt.Scan(&inputNum)
		sliceNum = append(sliceNum, inputNum)
	}

	// Removing num from slice

	fmt.Println("\nChoose index to delete: ")

	fmt.Scan(&inputCount)

	sliceNum = append(sliceNum[0:inputCount], sliceNum[inputCount+1:]...)

	// Print result

	fmt.Printf("\nSlice without your number: %d", sliceNum)
}
