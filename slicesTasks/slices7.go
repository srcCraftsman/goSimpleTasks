package slicesTasks

// 7. Удаление элемента из слайса
//    Напишите программу, которая удаляет элемент из слайса по заданному индексу.

import (
	"fmt"
)

func slice7() {
	var sliceNum []int = SliceGen()
	var inputID int

	// Removing num from slice

	fmt.Println("\nChoose index to delete: ")

	fmt.Scan(&inputID)

	sliceNum = append(sliceNum[0:inputID], sliceNum[inputID+1:]...)

	// Print result

	fmt.Printf("\nSlice without your number: %d", sliceNum)
}
