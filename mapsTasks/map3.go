package mapsTasks

//  3. Поиск элемента в карте:
//  Проверьте, существует ли студент с именем "John" в карте.
//  Если существует, выведите его возраст, если нет — выведите соответствующее сообщение.

import (
	"fmt"
)

func map3() {

	name := "John"

	value, ok := Students[name]
	if ok {
		fmt.Println(value, name)
	} else if !ok {
		fmt.Printf("\nИмени %s нет в карте, но есть Марат Мырзы брат.", name)
	}

}
