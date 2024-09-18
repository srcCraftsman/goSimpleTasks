package goroutinesTasks

//	6. Семафоры с использованием sync.Mutex
//	Создайте программу, в которой несколько Go-рутин изменяют общий счетчик.
//	Используйте sync.Mutex, чтобы обеспечить корректность изменения значения счетчика.

import (
	"fmt"
	"sync"
)

func goRoutine6() {
	mu := new(sync.Mutex)
	wg := new(sync.WaitGroup)
	var num int
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			num++
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(num)
}
