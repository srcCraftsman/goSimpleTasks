package readWriteTasks

//	10. Удаление строки из файла
//	Создайте программу, которая читает файл, удаляет все строки, содержащие определенное слово,
//  и записывает результат в новый файл. Обработайте ошибки, которые могут возникнуть при работе с файлами.

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readWrite10() {
	deletedWord := "Когда"

	file, err := os.Open(root + "/textCopyTask9.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	newFile, err := os.Create(root + "/textCopyTask10.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer newFile.Close()

	sc := bufio.NewScanner(file)
	wr := bufio.NewWriter(newFile)

	for sc.Scan() {
		str := sc.Text()
		mod := strings.ReplaceAll(str, deletedWord, " ")
		_, err = wr.WriteString(mod + "\n")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	if err := wr.Flush(); err != nil {
		fmt.Println(err)
		return
	}

}
