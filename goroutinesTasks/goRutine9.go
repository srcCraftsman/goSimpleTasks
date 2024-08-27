package goroutinesTasks

//	9. Использование таймера и канала time.After
//	Напишите программу, которая запускает Go-рутину для выполнения длительной задачи.
//	Если задача не завершается в течение заданного времени, программа должна вывести сообщение о таймауте и завершиться.

import (
	"fmt"
	"sync"
	"time"
)

func goRoutine9() {
	wg := new(sync.WaitGroup)
	stop := time.After(5 * time.Second)
	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 1; ; i++ {
			select {
			case <-stop:
				return
			default:
				fmt.Println(i)
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()
	wg.Wait()
}
