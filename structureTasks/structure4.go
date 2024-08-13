package structureTasks

import (
	"fmt"
)

//  4. Передача структуры в функцию
//  Напишите функцию PrintPerson, которая принимает структуру Person (из задания 41)
//  и выводит её поля в отформатированном виде. Создайте экземпляр структуры и передайте его в функцию.

func (p Person) PrintPerson() string {

	return fmt.Sprintf("Name: %v. Age: %v", p.Name, p.Age)
}

func structure4() {

	p := Person{Name: "Jack", Age: 21}

	fmt.Printf("\n%v", p.PrintPerson())
}
