package gorutinesTasks

import (
	"fmt"
)

func SetGorutine() {

	var val int = 99
	for val != 0 {
		fmt.Printf(`
Set task:

1.  Запуск нескольких Go-рутин.
2.  Синхронизация с использованием sync.WaitGroup.
3.  Использование канала для передачи данных между Go-рутинами.
4.  Обработка ошибок в Go-рутинах.
5.  Параллельная обработка данных.
6.  Семафоры с использованием sync.Mutex.
7.  Реализация пула Go-рутин.
8.  Параллельное чтение файлов.
9.  Использование таймера и канала time.After.
10. Параллельная обработка HTTP-запросов.

0.  Stop
`)

		fmt.Scan(&val)
		switch {
		case val == 1:
			goRutine1()
			fmt.Printf("\n")

		case val == 2:
			goRutine2()
			fmt.Printf("\n")

		case val == 3:
			goRutine3()
			fmt.Printf("\n")

		case val == 4:
			goRutine4()
			fmt.Printf("\n")

			//case val == 5:
			//	goRutine5()
			//	fmt.Printf("\n")
			//
			//case val == 6:
			//	goRutine6()
			//	fmt.Printf("\n")
			//
			//case val == 7:
			//	goRutine7()
			//	fmt.Printf("\n")
			//
			//case val == 8:
			//	goRutine8()
			//	fmt.Printf("\n")
			//
			//case val == 9:
			//	goRutine9()
			//	fmt.Printf("\n")
			//
			//case val == 10:
			//	goRutine10()
			//	fmt.Printf("\n")
		}
	}
}
