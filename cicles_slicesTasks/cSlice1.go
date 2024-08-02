package cSlicesTasks

// 1. Удаление всех вхождений элемента из слайса
//	  Напишите программу, которая удаляет все вхождения заданного элемента из слайса.

import (
	"fmt"
	"goSimpleTasks/slicesTasks"
)

func cSlice1() {

	var sliceNum []int = slicesTasks.SliceGen() //sliceNum used slice generator
	var improperNum int

	fmt.Println("Set unacceptable number:\n")
	fmt.Scan(&improperNum)

	//Search and delete improperNum in slice

	for i, num := range sliceNum {

		if num == improperNum {
			sliceNum = append(sliceNum[0:i], sliceNum[i+1:]...)
		}

		// second cycle to check and delete

		for i, num := range sliceNum {

			if num == improperNum {
				sliceNum = append(sliceNum[0:i], sliceNum[i+1:]...)
			}
		}

	}

	// Print result

	fmt.Println(sliceNum)
}
