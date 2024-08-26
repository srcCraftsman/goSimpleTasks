package goroutinesTasks

//	3. Использование канала для передачи данных между Go-рутинами
//	Напишите программу, в которой одна Go-рутина генерирует числа и отправляет их в канал,
//	а другая Go-рутина читает числа из канала и выводит их на экран.

import (
	"fmt"
	"sync"
)

func goRoutine3() {
	ch := make(chan int)
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		defer wg.Done()
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	go func() {
		defer wg.Done()
		for i := range ch {
			fmt.Print(i, " ")
		}
	}()
	wg.Wait()
}
