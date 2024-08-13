package structureTasks

import "fmt"

//  10. Методы для структур:
//	Создайте структуру Circle с полем Radius (тип float64). Напишите метод для этой структуры, который будет вычислять
//  и возвращать площадь круга. Проверьте работу метода на нескольких экземплярах структуры.

type Circle struct {
	Radius float64
}

func (c Circle) Square() float64 {
	return 3.14 * (c.Radius * c.Radius)
}
func structure10() {

	circ := Circle{5.5}

	fmt.Printf("\nSquare circle = %.2f", circ.Square())
}
