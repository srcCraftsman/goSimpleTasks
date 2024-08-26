package readWriteTasks

//	4. Чтение и обработка JSON из файла
//	Создайте программу, которая читает JSON-файл, десериализует его в структуру Go,
//	а затем выводит данные этой структуры на экран. Реализуйте обработку ошибок при чтении и декодировании JSON.

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type UserJ struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Twitter string `json:"twitter"`
	Github  string `json:"github"`
}

func readWrite4() {
	var jS UserJ

	file, err := os.Open(root + "/test.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	reader, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(reader, &jS)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(jS)

}
