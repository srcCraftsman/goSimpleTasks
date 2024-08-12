package mapsTasks

import "fmt"

func SetMaps() {

	var val int = 99
	for val != 0 {

		fmt.Print(`
1.  Создание простой карты
2.  Добавление и удаление элементов
3.  Поиск элемента в карте
4.  Итерация по карте
5.  Подсчёт количества элементов
6.  Изменение значений в карте
7.  Карта со сложными типами
8.  Карта с вложенными картами
9.  Сортировка карты по ключам
10. Объединение двух карт

0.  Exit

`)

		fmt.Scan(&val)
		switch {
		case val == 1:
			fmt.Println("\n", map1())
			fmt.Println("\n")

		case val == 2:

			map2()
			fmt.Println("\n")

		case val == 3:

			map3()
			fmt.Println("\n")

		case val == 4:

			map4()
			fmt.Println("\n")

		case val == 5:

			fmt.Printf("\nElements in map: %v", map5(Students))
			fmt.Println("\n")

			//	case val == 6:
			//
			//		map6()
			//		fmt.Println("\n")
			//
			//	case val == 7:
			//
			//		map7()
			//		fmt.Println("\n")
			//
			//	case val == 8:
			//
			//		map8()
			//		fmt.Println("\n")
			//
			//	case val == 9:
			//
			//		map9()
			//		fmt.Println("\n")
			//
			//	case val == 10:
			//
			//		map10()
			//		fmt.Println("\n")
		}
	}
}
