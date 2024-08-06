package cSlicesTasks

// 9. Поиск пересечения двух слайсов
//   Напишите программу, которая находит пересечение двух слайсов (общие элементы).
import (
	"fmt"
	"goSimpleTasks/slicesTasks"
)

func cSlice9() {

	// var sliceNum []int = slicesTasks.SliceGen()
	// var sliceNumSec []int = slicesTasks.SliceGen()
	sliceNum := slicesTasks.SliceGen()
	sliceNumSec := slicesTasks.SliceGen()
	// снова проще создать мапу для перебора, так мы будем знать были такие элементы уже или нет
	elementMap := make(map[int]bool)
	var sliceKek []int //slice for result

	// наполняем мапу элементами из первого слайса, дубли автоматом будут опущены в этом случае
	for _, num := range sliceNum {
		elementMap[num] = true
	}

	// перебираем второй слайс и сравниваем с элементами в мапе,
	for _, num := range sliceNumSec {
		if elementMap[num] {
			//  если элемет в мапе есть мы его добавляем
			sliceKek = append(sliceKek, num)
			// и сразу удаляем из мапы чтобы избежать дублей
			delete(elementMap, num)
		}
	}

	// for i := 0; i < len(sliceNum); i++ {

	// 	for _, num := range sliceNumSec {
	// 		if num == sliceNum[i] {
	// 			sliceKek = append(sliceKek, num)
	// 		}
	// 	}
	// }

	// Print result

	fmt.Println(sliceKek)
}
