package slicesTasks

// 5. Сумма элементов слайса
//   Напишите программу, которая считает сумму всех элементов в слайсе.

import (
	"fmt"
)

func slice5() {
	var sliceNum []int
	var inputCount, inputNum int // inputCount = N; inputNum = numbers added in slice
	var numSum int               // sum numbers from sliceNum
	fmt.Println("\nSet N:")
	fmt.Scan(&inputCount)
	fmt.Println("\nEnter the numbers:")

	// Add numbers in sliceNum

	for i := 0; i < inputCount; i++ {
		fmt.Scan(&inputNum)
		sliceNum = append(sliceNum, inputNum)
	}

	// Search sum

	for i := 0; i < inputCount; i++ {
		numSum = sliceNum[i] + numSum
	}

	// Print result

	fmt.Printf("Sum numbers from slice: %d", numSum)
}
