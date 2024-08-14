package mapsTasks

import "fmt"

//  7. Карта со сложными типами:
//  Создайте map, где ключами будут строки (имена студентов),
//  а значениями — структуры, содержащие возраст и группу студента.
//  Заполните карту тремя-четырьмя записями и выведите её на экран.

func map7() {

	type InfoSt struct {
		Age   int
		Group int
	}
	studentsHard := map[string]InfoSt{
		"Arkasha":  {Age: 16, Group: 101},
		"Stepan":   {Age: 17, Group: 102},
		"Mugiwara": {Age: 18, Group: 201},
		"Viktor":   {Age: 19, Group: 103},
	}

	fmt.Printf("\nMap: %v", studentsHard)

}
