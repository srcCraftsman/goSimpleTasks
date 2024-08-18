package funcTasks

import "fmt"

//	9. Функция высшего порядка
//	Напишите функцию ApplyOperation, которая принимает срез int и функцию, применяет эту функцию к каждому элементу
//  среза и возвращает новый срез с результатами. Напишите несколько различных операций, таких как удвоение
//  или вычисление квадрата, и проверьте работу ApplyOperation с ними.

func ApplyOperation(s []int, kek func(int) int) (ns []int) {
	for i := 0; i < len(s); i++ {
		kek(s[i])
		ns = append(ns, kek(s[i]))
	}
	return ns
}
func Mult(i int) int {
	return i * 2
}
func quaD(i int) int {
	return i * i
}

func funcTask9() {
	var a = []int{5, 4, 3, 6, 7, 10, 2, 1}

	fmt.Printf("\nNum: %v\nMult2: %v\nQuad: %v", a, ApplyOperation(a, Mult), ApplyOperation(a, quaD))
}
