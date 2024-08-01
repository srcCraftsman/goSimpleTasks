package slicesTasks

// 4. Среднее значение элементов слайса
//	  Напишите программу, которая вычисляет среднее значение элементов в слайсе.

import (
	"fmt"
	"sort"
)

func slice4() {

	var sliceNum []int
	var inputCount, inputNum int // inputCount = N; inputNum = numbers added in slice
	var average int              // average number from slice
	fmt.Println("Set N:")
	fmt.Scan(&inputCount)
	fmt.Println("\nEnter the numbers:")

	// Added numbers in sliceNum

	for i := 0; i < inputCount; i++ {
		fmt.Scan(&inputNum)
		sliceNum = append(sliceNum, inputNum)

	}

	//Sort numbers in sliceNum

	sort.Slice(sliceNum, func(i, j int) bool {
		return sliceNum[i] < sliceNum[j]
	})

	// Search average number in sliceNum

	for i := 0; i <= inputCount; i++ {

		if i != inputCount {
			inputCount--
		}

		// Или можно было сделать проще:
		//								average = sliceNum[inputCount/2]

	}

	average = sliceNum[inputCount]

	// Print result

	fmt.Printf("\nAverage number: %d", average)
}
