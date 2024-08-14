package mapsTasks

import "fmt"

//  10. Объединение двух карт:
//  Создайте две карты, содержащие студентов и их возрасты.
//  Напишите функцию, которая объединяет эти две карты в одну. Если ключи совпадают, выберите наибольшее значение возраста.

func map10() {

	students1 := map[string]int{
		"Oleh":  30,
		"Marat": 31,
		"Danya": 26,
	}
	fmt.Println(students1)

	students2 := map[string]int{
		"Medet":  23,
		"Adilka": 42,
		"Oleh":   29,
	}
	fmt.Println(students2)

	for k := range students2 {
		if students1[k] < students2[k] {
			students1[k] = students2[k]
		} else if students1[k] > students2[k] {
			students2[k] = students1[k]
		}
		students1[k] = students2[k]
	}
	fmt.Println(students1)
}
