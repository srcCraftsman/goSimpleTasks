package structureTasks

import "fmt"

//  5. Изменение структуры внутри функции:
//  Напишите функцию IncreaseAge, которая принимает указатель на структуру Person
//  и увеличивает значение возраста на 1 год. Проверьте работу функции на созданной структуре.

func (p *Person) IncreaseAge() {
	p.Age++
}

func structure5() {

	per := Person{Name: "Aik Broflowski", Age: 5}
	fmt.Printf("\nName: %v. Before age: %v", per.Name, per.Age)

	per.IncreaseAge()
	fmt.Printf("\nName: %v. After age: %v", per.Name, per.Age)

}
