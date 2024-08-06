package slicesTasks

// 10. Копирование слайса
//    Напишите программу, которая копирует один слайс в другой.

import (
	"fmt"
)

func slice10() {

	var sliceNum []int = SliceGen()                    //first slice from generator
	var sliceNumSec []int = make([]int, len(sliceNum)) //second slice is a copy sliceNum

	fmt.Printf("\nSlice 1: %d \nSlice 2: %d", sliceNum, sliceNumSec)

	// вот тут верно, мы потыкали копирование, объявление можно упростить как было в предыдущей задаче
	copy(sliceNumSec, sliceNum)

	// Print result

	fmt.Printf("\nCopy slice 1 to slice 2...\nSlice 1: %d\nSlice 2: %d", sliceNumSec, sliceNum)
}
