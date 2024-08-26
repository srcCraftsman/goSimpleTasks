package readWriteTasks

//	7. Чтение файла с использованием буфера
//	Напишите программу, которая открывает большой файл и читает его содержимое по частям (например, по 1024 байта),
//  используя буфер. Выводите каждую прочитанную часть на экран.

import (
	"fmt"
	"io"
	"os"
)

func readWrite7() {

	buf := make([]byte, 1024)

	file, err := os.Open(root + "/text.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil && err != io.EOF {
			fmt.Println(err)
		}
		fmt.Print(string(buf[:n]))
		fmt.Print("\n\n\nPart end\n\n")
	}

}
