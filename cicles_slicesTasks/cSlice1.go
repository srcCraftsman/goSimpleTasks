package cSlicesTasks

// 1. Удаление всех вхождений элемента из слайса
//	  Напишите программу, которая удаляет все вхождения заданного элемента из слайса.

import (
	"fmt"
	"goSimpleTasks/slicesTasks"
)

func cSlice1() {

	// немного не верная логика в подходе. У нас уже есть слайс, нам легче через него итерироваться и создать новый слайс, без ненужных вхождений

	// var sliceNum []int = slicesTasks.SliceGen() //sliceNum used slice generator

	// упрощаем
	sliceNum := slicesTasks.SliceGen()
	var improperNum int

	fmt.Printf("\nSet unacceptable number:\n")
	fmt.Scan(&improperNum)

	// Новый слайс, через который фильтруемся
	var filteredSlice []int

	// Итерируемся и пишем в новый слайс
	for _, num := range sliceNum {
		if num != improperNum {
			filteredSlice = append(filteredSlice, num)
		}
	}
	// //Search and delete improperNum in slice

	// for i, num := range sliceNum {

	// 	if num == improperNum {
	// 		sliceNum = append(sliceNum[0:i], sliceNum[i+1:]...)
	// 	}

	// 	// second cycle to check and delete

	// 	for i, num := range sliceNum {

	// 		if num == improperNum {
	// 			sliceNum = append(sliceNum[0:i], sliceNum[i+1:]...)
	// 		}
	// 	}

	// }

	// Print result

	fmt.Println(filteredSlice)
}
