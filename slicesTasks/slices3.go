package slicesTasks

// 3. Поиск минимального элемента в слайсе
//	  Напишите программу, которая находит минимальный элемент в заданном слайсе.

import (
	"fmt"
)

func slice3() {

	var sliceNum []int = SliceGen()

	// всегда есть вероятность что массив будет пустым, этот кейс лучше отследить
	if len(sliceNum) == 0 {
		fmt.Println("\nThe slice is empty.")
		return
	}

	var min int = sliceNum[0] // minimum number in slice

	// Search for minimum num in sliceNum

	for _, num := range sliceNum[1:] {
		// мы работаем с интами, поэтому нет смысла их переводить во float
		// min = int(math.Min(float64(min), float64(num)))

		// кастуем предыдущую схему
		if num < min {
			min = num
		}
	}

	// Print result

	fmt.Printf("\nMinimum number: %d", min)
}
