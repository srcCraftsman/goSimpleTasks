package interfacesTasks

import (
	"fmt"
)

func SetInterface() {

	var val int = 99
	for val != 0 {

		fmt.Print(`
1.  Создание простого интерфейса
2.  Реализация нескольких интерфейсов
3.  Проверка типа с использованием type assertion
4.  Использование type switch
5.  Композиция интерфейсов
6.  Интерфейсы и структуры
7.  Интерфейсы и полиморфизм
8.  Пустой интерфейс interface{}
9.  Интерфейс с указателем на структуру
10. Интерфейс с несколькими методами

0. Exit

`)

		fmt.Scan(&val)
		switch {
		case val == 1:

			interfaceTask1()
			fmt.Printf("\n")

		case val == 2:

			interfaceTask2()
			fmt.Printf("\n")

		case val == 3:

			interfaceTask3()
			fmt.Printf("\n")

		case val == 4:

			interfaceTask4()
			fmt.Printf("\n")

		case val == 5:

			interfaceTask5()
			fmt.Printf("\n")

			//		case val == 6:
			//
			//			interfaceTask6()
			//			fmt.Printf("\n")
			//
			//		case val == 7:
			//
			//			interfaceTask7()
			//			fmt.Printf("\n")
			//
			//		case val == 8:
			//
			//			interfaceTask8()
			//			fmt.Printf("\n")
			//
			//		case val == 9:
			//
			//			interfaceTask9()
			//			fmt.Printf("\n")
			//
			//		case val == 10:
			//
			//			interfaceTask10()
			//			fmt.Printf("\n")
		}
	}
}
