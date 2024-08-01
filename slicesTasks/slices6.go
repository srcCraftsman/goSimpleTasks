package slicesTasks

// 6. Сортировка слайса
//   Напишите программу, которая сортирует слайс по возрастанию.

import (
	"fmt"
	"sort"
)

func slice6() {
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

	//Sort numbers in sliceNum

	sort.Slice(sliceNum, func(i, j int) bool {
		return sliceNum[i] < sliceNum[j]
	})

	// Print result

	fmt.Printf("\nYour slice after sort: %d", sliceNum)
}
