package funcTasks

//	7. Использование defer в функции
//	Напишите функцию CloseFile, которая открывает файл и использует defer для его закрытия.
//	Проверьте, что файл закрывается даже в случае возникновения ошибки при его чтении.

import (
	"fmt"
	"io"
	"os"
)

func CloseFile(f string) (s string) {

	buf := make([]byte, 1024) // make bufer for printing file
	// open file
	file, err := os.Open(f)
	if err != nil {
		fmt.Println("\n", err)
	}
	// srazu defer close after open
	defer file.Close()

	// file stats for size info
	fi, err := file.Stat()
	if err != nil {
		return
	}
	// check file size
	size := fi.Size()
	if size == 0 {
		fmt.Printf("File %v is empty", f)
	}

	// reading file to bufer
	for {
		if file == nil {

			break
		}
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			break
		}
		s = fmt.Sprintf("\nFile %v size: %vbyte\nFile open:\n\n%v", f, size, string(buf[:n]))
	}
	return //return result s
}

// main func for check
func funcTask7() {

	file := CloseFile("/home/blacknoise/test.json")
	fmt.Println(file)

}
