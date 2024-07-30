package TaskSet

import (
	"fmt"
	"goSimpleTasks/ciclesTasks"
)

func SetCicles() {
	var val int = 99
	for val != 0 {
		fmt.Println(`
Set task:

1. Вывод чисел от 1 до 100
2. Сумма чисел от 1 до N
3. Произведение чисел от K до N
4. Четные числа от 1 до N
5. Факториал числа N
6. Фибоначчи до N
7. Обратный порядок чисел от N
8. Таблица умножения для числа N
9. Сумма четных чисел от 1 до N
10.Сумма нечетных чисел от 1 до N
0. Stop
`)

		fmt.Scan(&val)
		switch {
		case val == 1:
			ciclesTasks.Cicle1()
			fmt.Println("\n")
		case val == 2:

			fmt.Println("\nSet N")
			ciclesTasks.Cicle2()
			fmt.Println("\n")

		case val == 3:

			ciclesTasks.Cicle3()
			fmt.Println("\n")

		case val == 4:

			fmt.Println("\nSet N")
			ciclesTasks.Cicle4()
			fmt.Println("\n")

		case val == 5:

			fmt.Println("\nSet N")
			ciclesTasks.Cicle5()
			fmt.Println("\n")

		case val == 6:

			fmt.Println("\nSet N")
			ciclesTasks.Cicle6()
			fmt.Println("\n")

		case val == 7:

			fmt.Println("\nSet N")
			ciclesTasks.Cicle7()
			fmt.Println("\n")

		case val == 8:

			fmt.Println("\nSet N")
			ciclesTasks.Cicle8()
			fmt.Println("\n")

		case val == 9:

			fmt.Println("\nSet N")
			ciclesTasks.Cicle9()
			fmt.Println("\n")

		case val == 10:

			fmt.Println("\nSet N")
			ciclesTasks.Cicle10()
			fmt.Println("\n")
		}
	}
}
