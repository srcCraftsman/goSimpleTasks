package slicesTasks

// Генератор слайсов для использования в дальнейшем

import (
	"fmt"
)

func SliceGen() []int {

	var sliceNum []int
	var inputCount, inputNum int // inputCount = N; inputNum = numbers added in slice

	fmt.Println("\nSet N:")
	fmt.Scan(&inputCount)
	fmt.Println("\nEnter the numbers:")

	// Add numbers in sliceNum

	for ; 0 < inputCount; inputCount-- {
		fmt.Scan(&inputNum)
		sliceNum = append(sliceNum, inputNum)

	}
	return sliceNum
}
