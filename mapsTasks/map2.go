package mapsTasks

//  2. Добавление и удаление элементов:
//  В уже созданную карту добавьте еще одного студента.
//  Затем удалите одного из студентов из карты и снова выведите карту на экран.

import (
	"fmt"
)

func map2() {

	fmt.Println("\n", Students)

	Students["Arkasha"] = 19

	delete(Students, "Zohan")

	fmt.Println("\n", Students)

}
