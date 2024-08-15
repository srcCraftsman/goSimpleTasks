package interfacesTasks

//  1. Создание простого интерфейса
//  Определите интерфейс Speaker, который будет содержать метод Speak(). Создайте две структуры, Human и Dog,
//  и реализуйте метод Speak() для обеих. Создайте функцию, которая принимает Speaker и вызывает его метод Speak().
//  Проверьте работу функции с разными структурами.

import (
	"fmt"
)

type Speaker interface {
	Speak()
}

type Human struct {
	Name string
}
type Dog struct {
	Name string
}

func (h Human) Speak() {
	fmt.Printf("My name is %v and i say - ЭЩКЕРЕEEEE\n", h.Name)
}

func (d Dog) Speak() {
	fmt.Printf("My name is %v and i say - GAW GAW\n", d.Name)
}

func Say(s Speaker) {
	s.Speak()
}

func interfaceTask1() {

	h := Human{Name: "Vasya"}
	d := Dog{Name: "Sharik"}

	Say(h)
	Say(d)
}
