package slicesTasks

// 6. Сортировка слайса
//   Напишите программу, которая сортирует слайс по возрастанию.

import (
	"fmt"
	"sort"
)

func slice6() {
	var sliceNum []int = SliceGen()

	// проверка на пустоту
	if len(sliceNum) == 0 {
		fmt.Println("\nThe slice is empty.")
		return
	}

	//Sort numbers in sliceNum

	// sort.Slice(sliceNum, func(i, j int) bool {
	// 	return sliceNum[i] < sliceNum[j]
	// })

	// корректное решение, но есть функция сорта по интам и она вернет как раз по возрастанию слайс:
	sort.Ints(sliceNum)

	// Print result

	fmt.Printf("\nYour slice after sort: %d", sliceNum)
}
