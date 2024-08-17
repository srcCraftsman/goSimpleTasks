package funcTasks

//	1. Функция с несколькими аргументами
//  Напишите функцию Add, которая принимает два аргумента типа int и возвращает их сумму.
//  Проверьте работу функции на нескольких примерах.

import (
	"fmt"
)

func sumTwo(a, b int) int {
	return a + b
}
func funcTask1() {

	fmt.Printf("\na = %v\nb = %v\nsum = %v\n", 5, 2, sumTwo(5, 2))

}
