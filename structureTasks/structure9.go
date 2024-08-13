package structureTasks

import "fmt"

//  9. Срез структур
//	Создайте структуру Student с полями Name и Grade (оба типа string).
//	Создайте срез из 5 студентов, заполните его данными и выведите на экран.

type Student struct {
	Name, Grade string
}

func structure9() {

	slice := []Student{
		{Name: "Oleg", Grade: "ZBS"},
		{Name: "Emilio", Grade: "Norm"},
		{Name: "Danil", Grade: "4otko"}}

	fmt.Printf("\nStudent 1: %v\nStudent 2: %v\nStudent 3: %v", slice[0], slice[1], slice[2])

}
