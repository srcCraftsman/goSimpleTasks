package structureTasks

import "fmt"

//  3. Изменение значений полей структуры:
//  Создайте структуру Car с полями Brand, Model и Year (все поля типа string).
//  Создайте экземпляр этой структуры и измените значение поля Year. Выведите обновленный экземпляр на экран.

type Car struct {
	Brand, Model, Year string
}

func structure3() {

	ta4ka := Car{"Dodge", "Challenger", "1960"}
	ta4ka.Year = "2024"
	fmt.Println(ta4ka)

}
