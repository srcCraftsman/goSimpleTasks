package interfacesTasks

import "fmt"

//	8.  Пустой интерфейс interface{}
//	Напишите функцию, которая принимает пустой интерфейс interface{} и выводит на экран
//  информацию о типе и значении переданного параметра с использованием fmt.Printf.

func printType(t interface{}) {
	fmt.Printf("\nValue: %v. Type: %T", t, t)
}
func interfaceTask8() {

	t := "Kek"
	z := 3
	b := true

	printType(t)
	printType(z)
	printType(b)

}
