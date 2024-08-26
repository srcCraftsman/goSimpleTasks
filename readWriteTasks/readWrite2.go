package readWriteTasks

//	2. Запись в файл
//	Создайте программу, которая принимает от пользователя строку текста и записывает её в новый файл.
//  Если файл уже существует, программа должна перезаписать его содержимое.

import (
	"bufio"
	"fmt"
	"os"
)

func readWrite2() {
	file, err := os.Create(root + "/text2.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	inString := bufio.NewScanner(os.Stdin)

	for inString.Scan() {
		file.WriteString(inString.Text() + "\n")
		break
	}

}
