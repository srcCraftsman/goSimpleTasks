package cSlicesTasks

// 3. Поиск элемента в слайсе
//    Напишите программу, которая проверяет, содержится ли заданный элемент в слайсе.

import (
	"fmt"
	"goSimpleTasks/slicesTasks"
)

func cSlice3() {

	var sliceNum []int = slicesTasks.SliceGen() //generator
	var searchNum int
	var indexInSlice []int
	fmt.Println("\nSet number for search:  ")
	fmt.Scan(&searchNum)

	// search number in slice and add index to new slice for result

	for i, num := range sliceNum {
		if num == searchNum {
			indexInSlice = append(indexInSlice, i)
		}
	}
	// Print result

	if len(indexInSlice) > 0 {
		fmt.Printf("\nYour number %d based on the %d indexes.\n", searchNum, indexInSlice)
	} else {
		fmt.Println("\nThere is no such element.")
	}
}
