package interfacesTasks

import "fmt"

//	6. Интерфейсы и структуры
//	Создайте структуру Car с методами Start() и Stop(). Определите интерфейс Vehicle, включающий эти методы.
//  Создайте функцию, которая принимает Vehicle, и проверьте её работу с экземпляром Car.

type Car struct {
	Model string
	sound string
	power string
}
type Vehicle interface {
	Start()
	Stop()
}

func (c Car) Start() {
	c.sound = "wroom wroom"
	c.power = "ON"
	fmt.Printf("\nModel: %v. Power: %v. %v", c.Model, c.power, c.sound)
}
func (c Car) Stop() {
	c.sound = "silence"
	c.power = "OFF"
	fmt.Printf("\nModel: %v. Power: %v. %v", c.Model, c.power, c.sound)
}
func carTest(v Vehicle) {
	v.Start()
	v.Stop()
}
func interfaceTask6() {
	car := Car{Model: "Nexia"}
	carTest(car)
}
