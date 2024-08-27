package goroutinesTasks

import (
	"fmt"
)

func SetGoroutine() {

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
			goRoutine1()
			fmt.Printf("\n")

		case val == 2:
			goRoutine2()
			fmt.Printf("\n")

		case val == 3:
			goRoutine3()
			fmt.Printf("\n")

		case val == 4:
			goRoutine4()
			fmt.Printf("\n")

		case val == 5:
			goRoutine5()
			fmt.Printf("\n")

		case val == 6:
			goRoutine6()
			fmt.Printf("\n")

		case val == 7:
			goRoutine7()
			fmt.Printf("\n")

		case val == 8:
			goRoutine8()
			fmt.Printf("\n")

		case val == 9:
			goRoutine9()
			fmt.Printf("\n")

		case val == 10:
			goRoutine10()
			fmt.Printf("\n")
		}
	}
}
