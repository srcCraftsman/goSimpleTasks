package readWriteTasks

import (
	"fmt"
)

func SetReadWrite() {

	var val int = 99
	for val != 0 {
		fmt.Printf(`
Set task:

1.  Чтение файла построчно.
2.  Запись в файл.
3.  Копирование файла.
4.  Чтение и обработка JSON из файла.
5.  Запись структуры в файл в формате JSON.
6.  Добавление данных в существующий файл.
7.  Чтение файла с использованием буфера.
8.  Подсчет слов в файле.
9.  Поиск и замена текста в файле.
10. Удаление строки из файла.

0.  Stop
`)

		fmt.Scan(&val)
		switch {
		case val == 1:
			readWrite1()
			fmt.Printf("\n")

		case val == 2:
			readWrite2()
			fmt.Printf("\n")

		case val == 3:
			readWrite3()
			fmt.Printf("\n")

		case val == 4:
			readWrite4()
			fmt.Printf("\n")

		case val == 5:
			readWrite5()
			fmt.Printf("\n")

		case val == 6:
			readWrite6()
			fmt.Printf("\n")

		case val == 7:
			readWrite7()
			fmt.Printf("\n")

		case val == 8:
			readWrite8()
			fmt.Printf("\n")

		case val == 9:
			readWrite9()
			fmt.Printf("\n")

			//case val == 10:
			//	readWrite10()
			//	fmt.Printf("\n")
		}
	}
}
