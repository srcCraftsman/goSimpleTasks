package readWriteTasks

//	1. Чтение файла построчно
//	Напишите программу, которая открывает текстовый файл и читает его содержимое построчно,
//  выводя каждую строку на экран. Реализуйте обработку ошибок при открытии и чтении файла.

import (
	"bufio"
	"fmt"
	"os"
)

var root string = "/home/blacknoise/goSimpleTasks/readWriteTasks/files"

func readWrite1() {
	file, err := os.Open(root + "/text.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
