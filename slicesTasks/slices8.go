package slicesTasks

// 8. Добавление элемента в слайс
//    Напишите программу, которая добавляет элемент в слайс на заданную позицию.

import (
	"fmt"
)

func slice8() {
	var inputCount, inputNum int
	var sliceNum []int = SliceGen()
	var sliceNumIn []int = make([]int, len(sliceNum)) // copy first slice
	copy(sliceNumIn, sliceNum)

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
