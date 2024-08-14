package mapsTasks

//  1. Создание простой карты:
//  Создайте map для хранения имён студентов (ключи типа string) и их возрастов (значения типа int).
//  Инициализируйте карту с тремя студентами, выведите её на экран.

var Students = map[string]int{
	"Zohan":            18,
	"Myrza":            19,
	"Marat Myrzy brat": 20,
}

func map1() map[string]int {

	return Students
}
