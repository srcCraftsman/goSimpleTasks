package goroutinesTasks

import (
	"fmt"
	"sync"
)

//	5. Параллельная обработка данных
//	Напишите программу, которая параллельно обрабатывает элементы среза данных, используя Go-рутины.
//	Результаты обработки должны быть собраны в новый срез и выведены на экран.

func sum(i int, s []int, n []int, wg *sync.WaitGroup) {
	defer wg.Done()
	n[i] = s[i] * 2
	fmt.Print(i) // я так узнаю что обрабатывается именно в горутинах
}

func goRoutine5() {
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	newSlice := make([]int, len(slice))
	wg := new(sync.WaitGroup)
	for i := range slice {
		wg.Add(1)
		go sum(i, slice, newSlice, wg)
	}
	wg.Wait()
	fmt.Println("\n", newSlice)
}
