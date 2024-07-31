package ciclesTasks

import (
	"fmt"
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
			cicle1()
			fmt.Println("\n")
		case val == 2:

			fmt.Println("\nSet N")
			cicle2()
			fmt.Println("\n")

		case val == 3:

			cicle3()
			fmt.Println("\n")

		case val == 4:

			fmt.Println("\nSet N")
			cicle4()
			fmt.Println("\n")

		case val == 5:

			fmt.Println("\nSet N")
			cicle5()
			fmt.Println("\n")

		case val == 6:

			fmt.Println("\nSet N")
			cicle6()
			fmt.Println("\n")

		case val == 7:

			fmt.Println("\nSet N")
			cicle7()
			fmt.Println("\n")

		case val == 8:

			fmt.Println("\nSet N")
			cicle8()
			fmt.Println("\n")

		case val == 9:

			fmt.Println("\nSet N")
			cicle9()
			fmt.Println("\n")

		case val == 10:

			fmt.Println("\nSet N")
			cicle10()
			fmt.Println("\n")
		}
	}
}
