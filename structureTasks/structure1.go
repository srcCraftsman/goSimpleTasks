package structureTasks

import "fmt"

//  1. Создание простой структуры:
//  Определите структуру Person, которая будет содержать поля Name (тип string) и Age (тип int).
//  Создайте экземпляр этой структуры, заполните его значениями и выведите на экран.

type Person struct {
	Name string
	Age  int
}

func structure1() {

	var p = Person{"Jack", 21}

	fmt.Println(p)
}
