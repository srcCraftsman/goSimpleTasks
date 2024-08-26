package gorutinesTasks

//	4. Обработка ошибок в Go-рутинах
//	Создайте программу, которая запускает несколько Go-рутин, каждая из которых может завершиться с ошибкой.
//  Используйте канал для передачи ошибок в главную Go-рутину и выводите их на экран.

import (
	"fmt"
	"os"
	"sync"
)

func goRutine4() {
	chErr := make(chan string)

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		for eRR := range chErr {
			fmt.Println(eRR)
		}
	}()

	go func() {
		defer wg.Done()
		file, err := os.Open("/e/rrr.txt")
		if err != nil {
			chErr <- fmt.Sprint(err)
		}
		file.Close()

	}()
	go func() {
		defer wg.Done()
		file, err := os.Open("/tekst.txt")
		if err != nil {
			chErr <- fmt.Sprint(err)
		}
		file.Close()

	}()

	wg.Wait()
}
