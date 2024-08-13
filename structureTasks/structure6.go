package structureTasks

import "fmt"

//  6. Использование анонимных полей
//  Создайте структуру Employee, которая будет включать анонимные поля структур Person и Position
//  (где Position будет включать поля Title и Salary типа string и int соответственно).
//  Создайте экземпляр Employee и заполните все поля. Выведите результат на экран.

type Position struct {
	Title  string
	Salary int
}
type Employee struct {
	Person
	Position
}

func structure6() {

	empD := Employee{

		Person{
			Name: "Danya",
			Age:  26},

		Position{
			Title:  "Engineer",
			Salary: 1000000}}

	fmt.Printf("\nPerson: %v\nAge: %v\nPos: %v\nSalary: %v", empD.Name, empD.Age, empD.Title, empD.Salary)
}
