package readWriteTasks

//	9. Поиск и замена текста в файле
//	Напишите программу, которая читает текстовый файл, заменяет в нем все вхождения
//  одного слова на другое и записывает результат в новый файл.

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readWrite9() {
	newWord := " kek "
	oldWord := " в "

	file, err := os.Open(root + "/textCopy.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	newFile, err := os.Create(root + "/textCopyTask9.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer newFile.Close()

	sc := bufio.NewScanner(file)
	wr := bufio.NewWriter(newFile)

	for sc.Scan() {
		line := sc.Text()
		wordsNew := strings.ReplaceAll(line, oldWord, newWord)
		_, err = wr.WriteString(wordsNew + "\n")
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
