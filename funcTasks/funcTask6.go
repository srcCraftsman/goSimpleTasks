package funcTasks

//	6. Обработка ошибки в вызывающей функции
//  Напишите функцию OpenFile, которая пытается открыть файл по заданному пути и возвращает либо указатель на файл,
//  либо ошибку. Напишите код, который вызывает эту функцию, обрабатывает возможные ошибки
//	и выводит соответствующее сообщение.

import (
	"fmt"
	"os"
)

func OpenFile(f string) (*os.File, error) {
	file, err := os.Open(f)
	return file, err
}

func funcTask6() {
	file, err := OpenFile("/home/blacknoise/kek.go")

	info, _ := file.Stat()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Path: ", file.Name(), " Size: ", info.Size(), "bytes")
	}
	defer file.Close()
}
