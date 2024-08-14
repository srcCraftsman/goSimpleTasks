package mapsTasks

//  4. Итерация по карте:
//	Пройдитесь по карте с использованием цикла for range и выведите на экран имена студентов и их возраст.

import (
	"fmt"
)

func map4() {

	students := map1()

	for name, age := range students {

		fmt.Printf("\nStudent: %v. Age: %v", name, age)

	}

}
