package slicesTasks

// 4. Среднее значение элементов слайса
//	  Напишите программу, которая вычисляет среднее значение элементов в слайсе.

import (
	"fmt"
	"sort"
)

func slice4() {

	// Немного не ту логику реализовал, мы ищем среднее арифмитическое(мой косяк)

	var sliceNum []int = SliceGen()

	// проверка на пустоту
	if len(sliceNum) == 0 {
		fmt.Println("\nThe slice is empty.")
		return
	}

	var sum int

	// собираем сумму
	for _, num := range sliceNum {
		sum += num
	}

	// Считаем среднее. Здесь кастуем во float, т.к. деление может быть не целым
	average := float64(sum) / float64(len(sliceNum))

	sort.Slice(sliceNum, func(i, j int) bool {
		return sliceNum[i] < sliceNum[j]
	})

	// Print result

	fmt.Printf("\nAverage number: %d", average)
}
