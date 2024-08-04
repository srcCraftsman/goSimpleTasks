package cSlicesTasks

// 6. Перевернуть слайс
//    Напишите программу, которая переворачивает элементы слайса (реверс).
import (
	"fmt"
	"goSimpleTasks/slicesTasks"
)

func cSlice6() {

	var sliceNum []int = slicesTasks.SliceGen() //slice generator
	var l int = len(sliceNum) - 1
	//slice revers

	for i := 0; i < l; i++ {

		sliceNum[i], sliceNum[l] = sliceNum[l], sliceNum[i]
		l--

	}

	// Print result

	fmt.Println(sliceNum)

}
