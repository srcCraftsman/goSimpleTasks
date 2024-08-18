package funcTasks

//	5. Функция, возвращающая ошибку
//	Напишите функцию Sqrt, которая принимает float64 и возвращает значение квадратного корня и ошибку.
//  Ошибка должна возникать, если передано отрицательное число.
//  Проверьте работу функции с положительным и отрицательным значениями.

import (
	"errors"
	"fmt"
	"math"
)

func Sqrt(a float64) (float64, error) {

	if a <= 0 {
		return 0.0, errors.New("Your num < 0")
	} else {
		return math.Sqrt(a), nil
	}
}

func funcTask5() {

	// Num 9
	if s, err := Sqrt(9); err != nil {
		fmt.Printf("\n%v", err)
	} else {
		fmt.Printf("\n9 = %.3f", s)
	}
	// Num 25
	if s, err := Sqrt(23); err != nil {
		fmt.Printf("\n%v", err)
	} else {
		fmt.Printf("\n23 = %.3f", s)
	}
	// Num -9
	if s, err := Sqrt(-9); err != nil {
		fmt.Printf("\n-9 %v", err)
	} else {
		fmt.Printf("\n-9 = %.3f", s)
	}
	// Num -4
	if s, err := Sqrt(-4); err != nil {
		fmt.Printf("\n-4 %v", err)
	} else {
		fmt.Printf("\n-4 = %.3f", s)
	}
}
