package interfacesTasks

import "fmt"

//  2. Реализация нескольких интерфейсов
//  Создайте два интерфейса: Mover с методом Move() и Sayer с методом Say().
//  Создайте структуру Robot, которая реализует оба интерфейса.
//  Напишите функцию, которая принимает параметры обоих интерфейсов и вызывает соответствующие методы.

type Mover interface {
	Move()
}

type Sayer interface {
	Say()
}

type Robot struct {
	Model string
}

func (r Robot) Move() {
	fmt.Printf("\n%v move to right\n", r.Model)
}
func (r Robot) Say() {
	fmt.Printf("My model name: %v\n", r.Model)
}

func SayMo(m Mover, s Sayer) {
	m.Move()
	s.Say()
}

func interfaceTask2() {

	r := Robot{Model: "X-153"}

	SayMo(r, r)
}
