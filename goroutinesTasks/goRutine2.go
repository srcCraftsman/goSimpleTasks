package goroutinesTasks

//	2. Синхронизация с использованием sync.WaitGroup
//	Создайте программу, которая запускает несколько Go-рутин для выполнения разных задач
//	(например, чтение файлов или вычисления). Используйте sync.WaitGroup,
//	чтобы дождаться завершения всех Go-рутин перед завершением программы.

import (
	"fmt"
	"os"
	"sync"
)

func goRoutine2() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go readPrint(wg)
	go fact(6, wg)

	wg.Wait()
}

func fact(a int, wg *sync.WaitGroup) {
	defer wg.Done()
	p := 1
	for i := 1; i <= a; i++ {
		p *= i
	}
	fmt.Println("Factorial: ", p)
}
func readPrint(wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open("/home/blacknoise/goSimpleTasks/readWriteTasks/files/text2.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	buf := make([]byte, 1024)
	file.Read(buf)
	fmt.Println("From File:\n", string(buf))

}
