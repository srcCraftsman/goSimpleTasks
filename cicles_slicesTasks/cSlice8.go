package cSlicesTasks

// 8. Слияние и сортировка двух слайсов
//   Напишите программу, которая объединяет два слайса и сортирует результат.
import (
	"fmt"
	"goSimpleTasks/slicesTasks"
	"sort"
)

func cSlice8() {
	// можно упростить
	var sliceNum []int = slicesTasks.SliceGen()
	var sliceNumSec []int = slicesTasks.SliceGen()

	// правильное слияние
	sliceNum = append(sliceNum, sliceNumSec...)

	// Sort slice
	// тут можем заменить на sort.Ints(sliceNum) чтобы автоматом отсортировать, раз не указан метод сортировки, то нам подойдет любой

	// for i := 0; i < len(sliceNum)-1; i++ {
	// 	for i := 0; i < len(sliceNum)-1; i++ {
	// 		if sliceNum[i] > sliceNum[i+1] {
	// 			sliceNum[i], sliceNum[i+1] = sliceNum[i+1], sliceNum[i]
	// 		}
	// 	}
	// }
	sort.Ints(sliceNum)
	// Print result

	fmt.Println(sliceNum)
}
