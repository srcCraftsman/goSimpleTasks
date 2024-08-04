package cSlicesTasks

// 10. Разность двух слайсов
//     Напишите программу, которая находит разность двух слайсов (элементы, которые есть в первом слайсе, но нет во втором).
import (
	"fmt"
	"goSimpleTasks/slicesTasks"
)

func cSlice10() {

	var sliceNum []int = slicesTasks.SliceGen()
	var sliceNumSec []int = slicesTasks.SliceGen()
	var sliceKek1 []int // сюда записывается число каждый раз, когда оно не повторяется
	var sliceKek2 []int //slice for result

	for i := 0; i < len(sliceNum); i++ {

		for _, num := range sliceNumSec {
			if num != sliceNum[i] {
				sliceKek1 = append(sliceKek1, sliceNum[i])
			}
		}
	}
	// Здесь цикл смотрит, сколько раз повторяется число.

	// Если количество повторов числа равняется длине второго слайса,
	// то число вносится в слайс sliceKek2 для вывода результата.
	// Пока что я не додумался, как это сделать по другому, возможно позже переделаю.

	for i := 0; i < len(sliceNum); i++ {
		quantityNum := 0

		for _, num := range sliceKek1 {

			if num == sliceNum[i] {
				quantityNum++
			}
		}

		if quantityNum == len(sliceNumSec) {
			sliceKek2 = append(sliceKek2, sliceNum[i])
		}

	}

	// Print result
	fmt.Printf("\n%d", sliceKek2)

}
