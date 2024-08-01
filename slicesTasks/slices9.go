package slicesTasks

// 9. Объединение двух слайсов
//    Напишите программу, которая объединяет два слайса в один.

import (
	"fmt"
)

func slice9() {

	var sliceNum []int = SliceGen()                    //first slice from generator
	var sliceNumSec []int = make([]int, len(sliceNum)) //second slice is a copy sliceNum
	copy(sliceNumSec, sliceNum)                        //second generated
	sliceNum = SliceGen()

	fmt.Printf("\nSlice 1: %d\nSlice 2: %d", sliceNumSec, sliceNum)

	sliceNumSec = append(sliceNumSec, sliceNum...)

	// Print result

	fmt.Printf("\nSlices joined: %d", sliceNumSec)
}
