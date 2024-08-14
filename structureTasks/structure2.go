package structureTasks

import "fmt"

//  2. Создание и инициализация структуры
//  Создайте структуру Book, которая будет содержать поля Title, Author (оба типа string) и Pages (тип int).
//  Инициализируйте экземпляр структуры с использованием именованных полей и выведите его на экран.

type Book struct {
	Title, Author string
	Pages         int
}

func structure2() {

	var b = Book{Title: "Piknik na obo4ine", Author: "Strugatskie brothers", Pages: 250}

	fmt.Println(b)

}
