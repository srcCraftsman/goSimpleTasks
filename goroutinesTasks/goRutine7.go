package goroutinesTasks

//	7. Реализация пула Go-рутин
//	Напишите программу, которая создает пул Go-рутин для выполнения множества задач
//  (например, вычисление квадратов чисел). Каждая задача должна обрабатываться одной Go-рутиной из пула.

import (
	"fmt"
	"sync"
	"time"
)

func pullWorker(slice []int) {

	wg := new(sync.WaitGroup)
	var resultSlice []int
	start := time.Now
	wSize := len(slice)
	cha := make(chan int, wSize)

	for i := 1; i <= wSize; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for data := range cha {
				data = data * data
				resultSlice = append(resultSlice, data)
			}
			fmt.Printf("Worker %v done\n", i)
		}()
	}
	for i := range slice {
		cha <- slice[i]
	}
	close(cha)

	wg.Wait()

	fmt.Println(time.Since(start()))
	fmt.Println(resultSlice)

}

func goRoutine7() {
	var slice = []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	pullWorker(slice)
}
