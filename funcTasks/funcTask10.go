package funcTasks

//	10. Рекурсивная функция
//	Напишите рекурсивную функцию Factorial, которая принимает int и возвращает его факториал.
//	Проверьте работу функции с несколькими различными значениями.

import (
	"fmt"
)

func Factorial(a int) int {
	if a == 0 || a == 1 {
		return 1
	}
	return a * Factorial(a-1)
}
func funcTask10() {
	num := 1
	num2 := 4
	fmt.Printf("\nNum = %v\nFactorial = %v\n", num, Factorial(num))
	fmt.Printf("\nNum = %v\nFactorial = %v\n", num2, Factorial(num2))
}
