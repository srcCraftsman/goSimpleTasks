package cSlicesTasks

// 9. Поиск пересечения двух слайсов
//   Напишите программу, которая находит пересечение двух слайсов (общие элементы).
import (
	"fmt"
	"goSimpleTasks/slicesTasks"
)

func cSlice9() {

	var sliceNum []int = slicesTasks.SliceGen()
	var sliceNumSec []int = slicesTasks.SliceGen()
	var sliceKek []int //slice for result

	for i := 0; i < len(sliceNum); i++ {

		for _, num := range sliceNumSec {
			if num == sliceNum[i] {
				sliceKek = append(sliceKek, num)
			}
		}
	}

	// Print result

	fmt.Println(sliceKek)
}
