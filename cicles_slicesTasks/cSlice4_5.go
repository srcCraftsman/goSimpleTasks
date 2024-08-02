package cSlicesTasks

// 4. Сдвиг элементов слайса вправо
// 5. Сдвиг элементов слайса влево
//    Напишите программу, которая сдвигает элементы слайса вправо или влево на одну позицию.

import (
	"fmt"
	"goSimpleTasks/slicesTasks" //sliceShift new func in slicesTasks folder
)

func cSlice4_5() {

	var sliceNum []int = slicesTasks.SliceGen()
	var inputLR int
	fmt.Print(`
Where do you want to move slice:
1. Right
2. Left
	`)
	fmt.Scan(&inputLR)

	if inputLR == 1 {

		// Move to right

		sliceNum = slicesTasks.ShiftSlice(sliceNum, inputLR)

		// Move to left
	} else if inputLR == 2 {
		inputLR = -1
		sliceNum = slicesTasks.ShiftSlice(sliceNum, inputLR)
	}

	// Print result
	fmt.Println("\n", sliceNum)
}
