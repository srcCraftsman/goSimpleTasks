package cSlicesTasks

// 8. Слияние и сортировка двух слайсов
//   Напишите программу, которая объединяет два слайса и сортирует результат.
import (
	"fmt"
	"goSimpleTasks/slicesTasks"
)

func cSlice8() {

	var sliceNum []int = slicesTasks.SliceGen()
	var sliceNumSec []int = slicesTasks.SliceGen()

	sliceNum = append(sliceNum, sliceNumSec...)

	// Sort slice

	for i := 0; i < len(sliceNum)-1; i++ {
		for i := 0; i < len(sliceNum)-1; i++ {
			if sliceNum[i] > sliceNum[i+1] {
				sliceNum[i], sliceNum[i+1] = sliceNum[i+1], sliceNum[i]
			}
		}
	}
	// Print result

	fmt.Println(sliceNum)
}
