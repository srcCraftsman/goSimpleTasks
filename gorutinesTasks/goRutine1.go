package gorutinesTasks

//	1. Запуск нескольких Go-рутин
//	Напишите программу, которая запускает три Go-рутины, каждая из которых выводит свой уникальный идентификатор
//  и завершает работу через случайный промежуток времени.

import (
	"fmt"
	"time"
)

func goRutine1() {

	go func() {
		fmt.Println("Gorutine 1 stop")

	}()
	go func() {
		fmt.Println("Gorutine 2 stop")

	}()
	go func() {
		fmt.Println("Gorutine 3 stop")

	}()
	time.Sleep(5 * time.Second)
}
