package readWriteTasks

//	3. Копирование файла
//	Напишите программу, которая копирует содержимое одного файла в другой.
//	Обработайте ошибки, которые могут возникнуть при открытии, чтении и записи файлов.

import (
	"bufio"
	"fmt"
	"os"
)

func readWrite3() {
	fileRead, err := os.Open(root + "/text.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScan := bufio.NewReader(fileRead)
	defer fileRead.Close()

	fileCopy, err := os.Create(root + "/textCopy.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScan.WriteTo(fileCopy)
	fileCopy.Close()

}
