package structureTasks

import "fmt"

//  5. Изменение структуры внутри функции:
//  Напишите функцию IncreaseAge, которая принимает указатель на структуру Person
//  и увеличивает значение возраста на 1 год. Проверьте работу функции на созданной структуре.

func (p *Person) IncreaseAge() {
	p.Age++
}

func structure5() {

	per := Person{Name: "Aik", Age: 9}
	per.IncreaseAge()
	fmt.Println(per)

}
