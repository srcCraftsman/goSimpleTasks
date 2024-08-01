package slicesTasks

// 8. Добавление элемента в слайс
//    Напишите программу, которая добавляет элемент в слайс на заданную позицию.

import (
	"fmt"
)

func slice8() {
	var sliceNum []int
	var sliceNumIn []int         // copy first slice
	var inputCount, inputNum int // inputCount = N; inputNum = numbers added in slice
	fmt.Println("\nSet N:")
	fmt.Scan(&inputCount)
	fmt.Println("\nEnter the numbers:")

	// Add numbers in sliceNum

	for i := 0; i < inputCount; i++ {
		fmt.Scan(&inputNum)
		sliceNum = append(sliceNum, inputNum)
		sliceNumIn = append(sliceNumIn, inputNum)
	}

	// Add num to slice

	fmt.Println("\nChoose index to add: ")
	fmt.Scan(&inputCount)

	fmt.Println("\nSet number to add: ")
	fmt.Scan(&inputNum)

	sliceNumIn = append(sliceNumIn[0:inputCount], inputNum)
	sliceNumIn = append(sliceNumIn, sliceNum[inputCount:]...)

	// Print result

	fmt.Printf("\nSlice with your number: %d", sliceNumIn)
}
