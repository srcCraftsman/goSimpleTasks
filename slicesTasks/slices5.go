package slicesTasks

// 5. Сумма элементов слайса
//   Напишите программу, которая считает сумму всех элементов в слайсе.

import (
	"fmt"
)

func slice5() {
	var sliceNum []int = SliceGen()

	var numSum int // sum numbers from sliceNum

	// Search sum

	for i := 0; i < len(sliceNum); i++ {
		numSum = sliceNum[i] + numSum
	}

	// Print result

	fmt.Printf("Sum numbers from slice: %d", numSum)
}
