package mapsTasks

import "fmt"

//  6. Изменение значений в карте:
//  Для каждого студента в карте увеличьте возраст на 1 год, а затем выведите обновленную карту.

func map6() {

	fmt.Printf("\nYear 2024: %v", Students)

	for k, v := range Students {

		Students[k] = v + 1

	}

	fmt.Printf("\nYear 2025: %v", Students)

}
