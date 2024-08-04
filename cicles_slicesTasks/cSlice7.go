package cSlicesTasks

// 7. Уникальные элементы слайса
//    Напишите программу, которая создает новый слайс, содержащий только уникальные элементы из исходного слайса.
import (
	"fmt"
	"goSimpleTasks/slicesTasks"
)

func cSlice7() {

	var sliceNum []int = slicesTasks.SliceGen()
	var uniqSlice []int
	var uniqNum int

	for i := 0; i < len(sliceNum); i++ {
		fmt.Println("\nSet unique num: ")
		fmt.Scan(&uniqNum)

		for _, num := range sliceNum {

			if num == uniqNum {
				uniqSlice = append(uniqSlice, num)
			}
		}
	}

	// Print result

	fmt.Printf("\n%d", uniqSlice)
}
