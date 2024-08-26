package readWriteTasks

//	5. Запись структуры в файл в формате JSON
//	Напишите программу, которая создает экземпляр структуры Go и сериализует его в JSON.
//  Полученный JSON должен быть записан в файл. Если файл уже существует, программа должна перезаписать его содержимое.

import (
	"encoding/json"
	"fmt"
	"os"
)

type Human struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Weight int    `json:"weight"`
	Gender bool   `json:"gender"`
}

func readWrite5() {
	humanFirst := Human{"Apolinary Morgenshrek", 16, 90, false}
	file, err := os.Create(root + "/task5.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	data, err := json.MarshalIndent(humanFirst, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	if _, err = file.Write(data); err != nil {
		fmt.Println(err)
	}

}
