package slicesTasks

// 6. Сортировка слайса
//   Напишите программу, которая сортирует слайс по возрастанию.

import (
	"fmt"
	"sort"
)

func slice6() {
	var sliceNum []int = SliceGen()

	//Sort numbers in sliceNum

	sort.Slice(sliceNum, func(i, j int) bool {
		return sliceNum[i] < sliceNum[j]
	})

	// Print result

	fmt.Printf("\nYour slice after sort: %d", sliceNum)
}
