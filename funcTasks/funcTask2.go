package funcTasks

//	2. Функция с именованными возвращаемыми значениями
//	Создайте функцию Divide, которая принимает два аргумента типа int и возвращает два значения: результат деления
//  и остаток. Используйте именованные возвращаемые значения. Проверьте работу функции на нескольких примерах.

import (
	"fmt"
)

func Divide(a, b int) (quotient, remainder int) {
	quotient = a / b
	remainder = a % b
	return
}
func funcTask2() {
	a := 7
	b := 2
	q, r := Divide(a, b)
	fmt.Printf("\na = %v\nb = %v\ndelenie = %v\nostatok = %v\n", a, b, q, r)
}
