package slicesTasks

// 8. Добавление элемента в слайс
//    Напишите программу, которая добавляет элемент в слайс на заданную позицию.

import (
	"fmt"
)

func slice8() {
	var inputCount, inputNum int
	var sliceNum []int = SliceGen()

	// проверка на пустоту
	if len(sliceNum) == 0 {
		fmt.Println("\nThe slice is empty.")
		return
	}
	// просто для красоты вывведем изначальный слайс
	fmt.Printf("Original slice: %v\n", sliceNum)

	fmt.Println("\nChoose index to add: ")
	fmt.Scan(&inputCount)

	fmt.Println("\nSet number to add: ")
	fmt.Scan(&inputNum)

	// не совсем понял зачем делать копию, но хорошо что извучил что такой метод есть)
	// var sliceNumIn []int = make([]int, len(sliceNum)) // copy first slice
	// copy(sliceNumIn, sliceNum)

	// Add num to slice
	sliceNum = append(sliceNum[:inputCount], append([]int{inputNum}, sliceNum[inputCount:]...)...)

	// sliceNumIn = append(sliceNumIn[0:inputCount], inputNum)
	// sliceNumIn = append(sliceNumIn, sliceNum[inputCount:]...)

	// Print result

	fmt.Printf("\nSlice with your number: %d", sliceNum)
}
