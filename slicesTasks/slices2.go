package slicesTasks

// 2. Поиск максимального элемента в слайсе
//	  Напишите программу, которая находит максимальный элемент в заданном слайсе.

import (
	"fmt"
)

func slice2() {

	var sliceNum []int = SliceGen()

	// var max int // maximum number in slice

	// всегда есть вероятность что массив будет пустым, этот кейс лучше отследить
	if len(sliceNum) == 0 {
		fmt.Println("\nThe slice is empty.")
		return
	}

	// Search for maximum num in sliceNum

	// проще задать его сразу первым жлементом массива, это поможет в ситуации если слайс содержит отрицательные значения, либо состоит из одного элемента
	max := sliceNum[0]

	// Обходим массив начиная со следующего элемента
	for _, num := range sliceNum[1:] {
		// сравниваем только большие, т.к. то же значение у нас в переменной уже есть
		if num > max {
			max = num
		}
	}

	// Print result

	fmt.Printf("\nMaximum number: %d", max)
}
