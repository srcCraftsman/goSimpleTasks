package interfacesTasks

import "fmt"

//  4. Использование type switch
//  Создайте функцию, которая принимает interface{} и использует type switch для выполнения разных действий
//  в зависимости от типа переданного значения (например, int, string, bool).
//  В каждой ветке switch выведите сообщение с указанием типа и значением переменной.

func CheckCheck(v interface{}) {
	switch v.(type) {
	case string:
		fmt.Println("Type string: ", v)
	case int, float64:
		fmt.Println("Type int or float64: ", v)
	case bool:
		fmt.Println("Type bool: ", v)
	}
}

func interfaceTask4() {
	input := []interface{}{
		"dialogi tet-a-tet",
		543,
		3.14,
		"1" == "one",
	}
	for _, v := range input {
		CheckCheck(v)
	}
}
