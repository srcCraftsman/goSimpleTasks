package interfacesTasks

import "fmt"

//  3. Проверка типа с использованием type assertion
//  Напишите функцию, которая принимает параметр типа interface{} и проверяет, является ли он строкой,
//  числом или булевым значением. Для каждого типа выведите соответствующее сообщение.

func typeCheck(i interface{}) {
	if s, ok := i.(string); ok {
		fmt.Println(s, "= type string")
	} else if s, ok := i.(int); ok {
		fmt.Println(s, "= type int")
	} else if s, ok := i.(bool); ok {
		fmt.Println(s, "= type bool")
	} else if s, ok := i.(float64); ok {
		fmt.Println(s, "= type float64")
	}

}
func interfaceTask3() {

	input := []interface{}{
		"kek",
		3,
		5.4,
		3 == 0,
	}

	for _, v := range input {
		typeCheck(v)
	}

}
