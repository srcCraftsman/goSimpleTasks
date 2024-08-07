package cSlicesTasks

// 6. Перевернуть слайс
//    Напишите программу, которая переворачивает элементы слайса (реверс).
import (
	"fmt"
	"goSimpleTasks/slicesTasks"
)

func cSlice6() {

	// var sliceNum []int = slicesTasks.SliceGen() //slice generator
	sliceNum := slicesTasks.SliceGen() // Generate the slice
	l := len(sliceNum)
	// var l int = len(sliceNum) - 1

	// Slice reverse
	// чуть меняем подход, чтобы избежать ошибок out of range при необычной длинне
	for i := 0; i < l/2; i++ {
		sliceNum[i], sliceNum[l-1-i] = sliceNum[l-1-i], sliceNum[i]
	}

	// for i := 0; i < l; i++ {

	// 	sliceNum[i], sliceNum[l] = sliceNum[l], sliceNum[i]
	// 	l--

	// }

	// Print result

	fmt.Println(sliceNum)

}
