package goroutinesTasks

//	10. Параллельная обработка HTTP-запросов
//	Создайте программу, которая отправляет несколько HTTP-запросов параллельно, используя Go-рутины.
//	Результаты запросов должны быть собраны и выведены на экран.
//	Реализуйте обработку ошибок и таймаутов для каждого запроса.

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func httpReq(wg *sync.WaitGroup, s string, ch chan string) {
	defer wg.Done()
	resp, err := http.Get(s)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	ch <- string(body)
}
func goRoutine10() {

	page1 := "https://jsonplaceholder.typicode.com/posts/1"
	page2 := "https://jsonplaceholder.typicode.com/comments/15"
	page3 := "https://jsonplaceholder.typicode.com/users/8"

	wg := new(sync.WaitGroup)
	wg.Add(3)
	ch := make(chan string, 3)

	go httpReq(wg, page1, ch)
	go httpReq(wg, page2, ch)
	go httpReq(wg, page3, ch)

	wg.Wait()
	close(ch)
	for val := range ch {
		fmt.Println(val)
	}
}
