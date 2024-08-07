package slicesTasks

// 5. Сумма элементов слайса
//   Напишите программу, которая считает сумму всех элементов в слайсе.

import (
	"fmt"
)

func slice5() {
	var sliceNum []int = SliceGen()

	var numSum int // sum numbers from sliceNum

	// проверка на пустоту
	if len(sliceNum) == 0 {
		fmt.Println("\nThe slice is empty.")
		return
	}

	// Search sum
	// for i := 0; i < len(sliceNum); i++ {
	// 	numSum = sliceNum[i] + numSum
	// }

	// в целом верное решение, но можно использовать range:
	// так чуть короче

	for _, v := range sliceNum {
		numSum += v
	}

	// Print result

	fmt.Printf("Sum numbers from slice: %d", numSum)
}
