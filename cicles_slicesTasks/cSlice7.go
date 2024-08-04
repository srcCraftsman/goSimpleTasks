package cSlicesTasks

// 7. Уникальные элементы слайса
//    Напишите программу, которая создает новый слайс, содержащий только уникальные элементы из исходного слайса.
import (
	"fmt"
	"goSimpleTasks/slicesTasks"
)

func cSlice7() {

	var sliceNum []int = slicesTasks.SliceGen()
	fmt.Println("\nAdd unique num to new slice: ")
	var uniqNum []int = slicesTasks.SliceGen()
	var uniqSlice []int

	for i := 0; i < len(sliceNum); i++ {

		for _, num := range uniqNum {

			if num == sliceNum[i] {
				uniqSlice = append(uniqSlice, num)
			}
		}
	}

	// Print result

	fmt.Printf("\n%d", uniqSlice)
}
