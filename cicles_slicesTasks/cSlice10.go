package cSlicesTasks

// 10. Разность двух слайсов
//     Напишите программу, которая находит разность двух слайсов (элементы, которые есть в первом слайсе, но нет во втором).
import (
	"fmt"
	"goSimpleTasks/slicesTasks"
)

func cSlice10() {

	// здесь задачу также надо решать через мапу, как и предыдущую, оставлю тебе на самостоятельное изыскание)

	sliceNum := slicesTasks.SliceGen()
	sliceNumSec := slicesTasks.SliceGen()
	var sliceKek []int
	elementMap := make(map[int]bool)

	for _, num := range sliceNumSec {
		elementMap[num] = true
	}

	for _, num := range sliceNum {
		if elementMap[num] != true {

			sliceKek = append(sliceKek, num)
			delete(elementMap, num)
		}
	}

	fmt.Printf("First slice :%d\nSecond slice: %d\nResult: %d\n", sliceNum, sliceNumSec, sliceKek)

}
