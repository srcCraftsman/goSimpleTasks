package slicesTasks

// 9. Объединение двух слайсов
//    Напишите программу, которая объединяет два слайса в один.

import (
	"fmt"
)

func slice9() {
	// var sliceNum []int = SliceGen()                    //first slice from generator
	// var sliceNumSec []int = make([]int, len(sliceNum)) //second slice is a copy sliceNum
	// copy(sliceNumSec, sliceNum)                        //second generated
	// sliceNum = SliceGen()

	// %d не совсем верно, т.к. ожидает цифру, а нам лучше значение %v
	// fmt.Printf("\nSlice 1: %d\nSlice 2: %d", sliceNumSec, sliceNum)

	// sliceNumSec = append(sliceNumSec, sliceNum...)

	// более простое объявление слайса
	sliceNum1 := SliceGen()

	// более простое объявление слайса
	sliceNum2 := SliceGen()

	// пачатаем исходные слайсы
	fmt.Printf("\nSlice 1: %v\nSlice 2: %v", sliceNum1, sliceNum2)

	// просто склейиваем в новую переменную
	mergedSlice := append(sliceNum1, sliceNum2...)

	// Print the result
	fmt.Printf("\nSlices joined: %v\n", mergedSlice)

}
