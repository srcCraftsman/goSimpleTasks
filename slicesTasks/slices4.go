package slicesTasks

// 4. Среднее значение элементов слайса
//	  Напишите программу, которая вычисляет среднее значение элементов в слайсе.

import (
	"fmt"
	"sort"
)

func slice4() {

	var sliceNum []int = SliceGen()
	var average int // average number from slice

	//Sort numbers in sliceNum

	sort.Slice(sliceNum, func(i, j int) bool {
		return sliceNum[i] < sliceNum[j]
	})

	// Search average number in sliceNum

	average = sliceNum[len(sliceNum)/2]

	// Print result

	fmt.Printf("\nAverage number: %d", average)
}
