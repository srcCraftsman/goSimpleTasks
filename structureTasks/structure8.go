package structureTasks

import "fmt"

//  8. Сравнение структур
//	Создайте две структуры Rectangle с полями Width и Height (оба типа int).
//  Напишите функцию, которая будет сравнивать два прямоугольника по площади и выводить результат сравнения.

type Rectangle struct {
	Width, Height int
}

func (r Rectangle) Square() int {
	return r.Width * r.Height
}
func TwoRec(a, b Rectangle) string {
	if a.Square() > b.Square() {
		return fmt.Sprintf("Rectangle 1 - %v", a.Square())
	} else {
		return fmt.Sprintf("Rectangle 2 - %v", b.Square())
	}
}

func structure8() {

	rec1 := Rectangle{5, 3}
	rec2 := Rectangle{7, 4}

	fmt.Printf("\nRectangle 1: %v\nRectangle 2: %v\nSquare more: %v\n", rec1, rec2, TwoRec(rec1, rec2))

}
