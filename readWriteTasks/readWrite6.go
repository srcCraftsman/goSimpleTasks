package readWriteTasks

//	6. Добавление данных в существующий файл
//	Создайте программу, которая открывает файл и добавляет новую строку текста в конец файла.
//	Если файла не существует, программа должна его создать и записать строку.

import (
	"fmt"
	"os"
)

func readWrite6() {
	file, err := os.OpenFile(root+"/text2.txt", os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.WriteString("\nDo utra za zhili byli")
}
