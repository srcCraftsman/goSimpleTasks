package mapsTasks

import "fmt"

//  8. Карта с вложенными картами:
//  Создайте map, где ключом будет название курса (string), а значением — другая карта,
//  которая будет хранить имя студента и его оценку за курс. Заполните её несколькими курсами и студентами,
//  затем выведите на экран.

func map8() {

	kurs := make(map[string]map[string]int)

	kurs["files"] = map[string]int{
		"Danya": 2,
		"Aleg":  5,
		"Myrza": 4,
	}
	kurs["cycles"] = map[string]int{
		"Danya": 5,
		"Aleg":  5,
		"Myrza": 4,
	}
	kurs["math"] = map[string]int{
		"Danya": 3,
		"Aleg":  5,
		"Myrza": 3,
	}
	for name, students := range kurs {

		fmt.Printf("\nKurs: %v", name)

		for student, val := range students {

			fmt.Printf("\nStudent: %v. Grade: %v.", student, val)
		}
	}
}
