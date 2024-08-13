package structureTasks

import (
	"fmt"
)

func SetStructure() {

	var val int = 99
	for val != 0 {

		fmt.Print(`
1.  Создание простой структуры
2.  Создание и инициализация структуры
3.  Изменение значений полей структуры
4.  Передача структуры в функцию
5.  Изменение структуры внутри функции
6.  Использование анонимных полей
7.  Вложенные структуры
8.  Сравнение структур
9.  Срез структур
10. Методы для структур

0. Exit

`)

		fmt.Scan(&val)
		switch {
		case val == 1:

			structure1()
			fmt.Printf("\n")

		case val == 2:

			structure2()
			fmt.Printf("\n")

		case val == 3:

			structure3()
			fmt.Printf("\n")

		case val == 4:

			structure4()
			fmt.Printf("\n")

		case val == 5:

			structure5()
			fmt.Printf("\n")

		case val == 6:

			structure6()
			fmt.Printf("\n")

		case val == 7:

			structure7()
			fmt.Printf("\n")

		case val == 8:

			structure8()
			fmt.Printf("\n")

		case val == 9:

			structure9()
			fmt.Printf("\n")

		case val == 10:

			structure10()
			fmt.Printf("\n")
		}
	}
}
