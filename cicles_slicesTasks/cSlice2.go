package cSlicesTasks

// 2. Подсчет количества вхождений элемента в слайсе
//	  Напишите программу, которая подсчитывает количество вхождений заданного элемента в слайсе.

import (
	"fmt"
	"goSimpleTasks/slicesTasks"
)

func cSlice2() {

	var sliceNum []int = slicesTasks.SliceGen() //sliceNum used slice generator
	var interestNum int
	var quantityNum int

	fmt.Println("Set the number of interest:\n")
	fmt.Scan(&interestNum)

	// Search quantity elements == interestNum in sliceNum

	for _, num := range sliceNum {
		if num == interestNum {
			quantityNum++
		}

	}

	// Print result
	fmt.Printf("\nYour number repeating: %d", quantityNum)
}
