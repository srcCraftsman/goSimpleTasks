package slicesTasks

// 7. Удаление элемента из слайса
//    Напишите программу, которая удаляет элемент из слайса по заданному индексу.

import (
	"fmt"
)

func slice7() {
	var sliceNum []int = SliceGen()
	var inputID int

	// проверка на пустоту
	if len(sliceNum) == 0 {
		fmt.Println("\nThe slice is empty.")
		return
	}

	// просто для красоты вывведем изначальный слайс
	fmt.Printf("Original slice: %v\n", sliceNum)

	// Removing num from slice

	fmt.Println("\nChoose index to delete: ")

	fmt.Scan(&inputID)

	// тут важная проверочка, чтобы тебе юзверь не выдал индекс, которого нет
	if inputID < 0 || inputID >= len(sliceNum) {
		fmt.Println("Index out of range.")
		return
	}

	sliceNum = append(sliceNum[0:inputID], sliceNum[inputID+1:]...)

	// Print result

	fmt.Printf("\nSlice without your number: %d", sliceNum)
}
