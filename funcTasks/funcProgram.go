package funcTasks

import (
	"fmt"
)

func SetFunc() {

	var val int = 99
	for val != 0 {

		fmt.Print(`
1.  Функция с несколькими аргументами
2.  Функция с именованными возвращаемыми значениями
3.  Функция с переменным числом аргументов
4.  Функция с возвратом нескольких значений
5.  Функция, возвращающая ошибку
6.  Обработка ошибки в вызывающей функции
7.  Использование defer в функции
8.  Функция с указателем в аргументе
9.  Функция высшего порядка
10. Рекурсивная функция

0. Exit

`)

		fmt.Scan(&val)
		switch {

		case val == 1:

			funcTask1()
			fmt.Printf("\n")

		case val == 2:

			funcTask2()
			fmt.Printf("\n")

		case val == 3:

			funcTask3()
			fmt.Printf("\n")

			//case val == 4:
			//
			//	funcTask4()
			//	fmt.Printf("\n")
			//
			//case val == 5:
			//
			//	funcTask5()
			//	fmt.Printf("\n")
			//
			//case val == 6:
			//
			//	funcTask6()
			//	fmt.Printf("\n")
			//
			//case val == 7:
			//
			//	funcTask7()
			//	fmt.Printf("\n")
			//
			//case val == 8:
			//
			//	funcTask8()
			//	fmt.Printf("\n")
			//
			//case val == 9:
			//
			//	funcTask9()
			//	fmt.Printf("\n")
			//
			//case val == 10:
			//
			//	funcTask10()
			//	fmt.Printf("\n")
		}
	}
}
