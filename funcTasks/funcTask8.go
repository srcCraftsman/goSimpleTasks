package funcTasks

//	8. Функция с указателем в аргументе
//  Создайте функцию Increment, которая принимает указатель на int и увеличивает его значение на 1.
//	Проверьте работу функции, вызвав её для переменной типа int.

import (
	"fmt"
)

func Increment(a *int) {
	*a++
}

func funcTask8() {
	num := 5
	fmt.Println("Before:", num)
	Increment(&num)
	fmt.Println("After: ", num)
	Increment(&num)
	fmt.Println("After: ", num)
	Increment(&num)
	fmt.Println("After: ", num)
}
