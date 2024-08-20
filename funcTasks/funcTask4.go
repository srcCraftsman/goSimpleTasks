package funcTasks

//	4. Функция с возвратом нескольких значений
//	Создайте функцию MinMax, которая принимает срез int и возвращает два значения:
//	минимальное и максимальное значения в срезе. Напишите код, который вызывает эту функцию и выводит результаты.

import (
	"fmt"
	"goSimpleTasks/slicesTasks"
	"slices"
)

func MinMax(z []int) (min, max int) {
	max = slices.Max(z)
	min = slices.Min(z)
	return
}
func funcTask4() {
	s := slicesTasks.SliceGen()
	min, max := MinMax(s)
	fmt.Printf("\nSlice: %v\nMin: %v\nMax: %v\n", s, min, max)
}
