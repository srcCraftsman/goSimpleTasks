package slicesTasks

import (
	"fmt"
)

func SetSlices() {
	var val int = 99
	for val != 0 {
		fmt.Println(`
Set task:

1.  Напишите программу, которая создает слайс из первых N целых чисел и выводит его.
2.  Напишите программу, которая находит максимальный элемент в заданном слайсе.
3.  Напишите программу, которая находит минимальный элемент в заданном слайсе.
4.  Напишите программу, которая вычисляет среднее значение элементов в слайсе.
5.  Напишите программу, которая считает сумму всех элементов в слайсе.
6.  Напишите программу, которая сортирует слайс по возрастанию.
7.  Напишите программу, которая удаляет элемент из слайса по заданному индексу.
8.  Напишите программу, которая добавляет элемент в слайс на заданную позицию.
9.  Напишите программу, которая объединяет два слайса в один.
10. Напишите программу, которая копирует один слайс в другой.
0.  Stop
`)

		fmt.Scan(&val)
		switch {
		case val == 1:
			slice1()
			fmt.Println("\n")
		case val == 2:
			slice2()
			fmt.Println("\n")

		case val == 3:
			slice3()
			fmt.Println("\n")
			//
			//	case val == 4:
			//		slice4()
			//		fmt.Println("\n")
			//
			//	case val == 5:
			//
			//		fmt.Println("\nSet N")
			//		slice5()
			//		fmt.Println("\n")
			//
			//	case val == 6:
			//
			//		fmt.Println("\nSet N")
			//		slice6()
			//		fmt.Println("\n")
			//
			//	case val == 7:
			//
			//		fmt.Println("\nSet N")
			//		slice7()
			//		fmt.Println("\n")
			//
			//	case val == 8:
			//
			//		fmt.Println("\nSet N")
			//		slice8()
			//		fmt.Println("\n")
			//
			//	case val == 9:
			//
			//		fmt.Println("\nSet N")
			//		slice9()
			//		fmt.Println("\n")
			//
			//	case val == 10:
			//
			//		fmt.Println("\nSet N")
			//		slice10()
			//		fmt.Println("\n")
		}
	}
}
