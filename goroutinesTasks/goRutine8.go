package goroutinesTasks

//	8. Параллельное чтение файлов
//	Создайте программу, которая параллельно читает содержимое нескольких файлов с использованием Go-рутин.
//  Результаты чтения должны быть собраны и выведены на экран.

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

func read(wg *sync.WaitGroup, filePath string) {

	defer wg.Done()

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	read, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\nGoRoutine read start\n%v\nGoRoutine read done\n", string(read))
}

func goRoutine8() {
	start := time.Now()
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go read(wg, "/home/blacknoise/goSimpleTasks/readWriteTasks/files/text2.txt")
	go read(wg, "/home/blacknoise/goSimpleTasks/readWriteTasks/files/text.txt")
	wg.Wait()
	fmt.Println("\nTime: ", time.Since(start))
}
