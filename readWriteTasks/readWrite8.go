package readWriteTasks

//	8. Подсчет слов в файле
//	Создайте программу, которая открывает текстовый файл и подсчитывает количество слов в нем.
//  Выведите результат на экран.

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readWrite8() {
	file, err := os.Open(root + "/text2.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	wordsC := 0

	for sc.Scan() {
		line := sc.Text()
		words := strings.Fields(line)
		wordsC = wordsC + len(words)
	}
	fmt.Println("Words: ", wordsC)
}
