package cSlicesTasks

import (
	"fmt"
)

func cSlice1() {

	var sliceNum []int
	var improperNum int
	fmt.Println("Set unacceptable number:\n")
	fmt.Scan(&improperNum)

	var inputCount, inputNum int // inputCount = N; inputNum = numbers added in slice

	fmt.Println("\nSet lenght:")
	fmt.Scan(&inputCount)
	fmt.Println("\nEnter the numbers:")

	for ; inputCount > 0; inputCount-- {
		fmt.Scan(&inputNum)
		if inputNum != improperNum {
			sliceNum = append(sliceNum, inputNum)
		}
	}
	fmt.Println(sliceNum)

}
