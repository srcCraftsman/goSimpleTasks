package cSlicesTasks

// 7. Уникальные элементы слайса
//    Напишите программу, которая создает новый слайс, содержащий только уникальные элементы из исходного слайса.
import (
	"fmt"
	"goSimpleTasks/slicesTasks"
)

func cSlice7() {

	// не совсем понятно, зачем делать две генерации слайса тут, нам надо итерировать через первый и писать во новый
	sliceNum := slicesTasks.SliceGen()
	uniqueMap := make(map[int]bool) // создаем мапу, где будем хранитьуникальность элемента
	var uniqSlice []int

	// Итерируем через первый и обновлем нашу мапу значениями, которые в процессе чекаем
	for _, num := range sliceNum {
		if !uniqueMap[num] {
			uniqueMap[num] = true
			uniqSlice = append(uniqSlice, num)
		}
	}

	// var sliceNum []int = slicesTasks.SliceGen()
	// fmt.Println("\nAdd unique num to new slice: ")
	// var uniqNum []int = slicesTasks.SliceGen()
	// var uniqSlice []int

	// for i := 0; i < len(sliceNum); i++ {

	// 	for _, num := range uniqNum {

	// 		if num == sliceNum[i] {
	// 			uniqSlice = append(uniqSlice, num)
	// 		}
	// 	}
	// }

	// Print result

	fmt.Printf("\n%d", uniqSlice)
}
