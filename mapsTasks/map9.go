package mapsTasks

//  9. Сортировка карты по ключам:
//	Напишите функцию, которая принимает карту, сортирует её по ключам (в алфавитном порядке)
//  и выводит отсортированные ключи и значения.

import (
	"sort"
	"strconv"
)

func map9(a map[string]int) []string {
	var str []string
	var str2 []string

	for k, _ := range a {
		str = append(str, k)
	}
	sort.Strings(str)

	for _, v := range str {
		str2 = append(str2, v, ":", strconv.Itoa(a[v]), " ")
	}

	return str2
}
