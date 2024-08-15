package interfacesTasks

import "fmt"

//	7. Интерфейсы и полиморфизм
//	Создайте интерфейс Shape с методом Area(), который будет возвращать площадь фигуры. Реализуйте этот интерфейс
//  для структур Rectangle и Circle. Напишите функцию, которая принимает Shape и выводит его площадь.

type Shape interface {
	Area() float64
}
type Rectangle struct {
	a, b float64
}
type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.a * r.b
}
func (c Circle) Area() float64 {
	return 3.14 * (c.radius * c.radius)
}
func Square(s Shape) {
	fmt.Printf("\nSquare of %T = %v", s, s.Area())
}
func interfaceTask7() {

	circ := Circle{radius: 5}
	rect := Rectangle{5, 3}

	Square(circ)
	Square(rect)

}
