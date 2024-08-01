package slicesTasks

// 1. Создание слайса
//	  Напишите программу, которая создает слайс из первых N целых чисел и выводит его.

import (
	"fmt"
)

func slice1() {

	sliceNum := SliceGen()

	// Print result

	fmt.Println("\nSlice: ", sliceNum)

}
