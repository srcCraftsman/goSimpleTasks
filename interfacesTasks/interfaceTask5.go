package interfacesTasks

import "fmt"

//  5. Композиция интерфейсов
//  Создайте два интерфейса: Reader с методом Read() и Writer с методом Write().
//  Определите третий интерфейс ReadWriter, который будет включать оба предыдущих интерфейса.
//  Создайте структуру, которая реализует ReadWriter, и проверьте её работу в функции, принимающей ReadWriter.

type Reader interface {
	Read() string
}
type Writer interface {
	Write(string)
}
type ReadWriter interface {
	Reader
	Writer
}
type User struct {
	UserDataBuffer string
}

// пишу то что считываю в UserDataBuffer
func (u *User) Write(s string) {
	u.UserDataBuffer = s
}

// считываю из буфера и вывожу то что было записано
func (u *User) Read() string {
	return u.UserDataBuffer
}

// записываю в буффер и сразу вывожу результат
func DataReadPrint(rw ReadWriter, s string) {
	rw.Write(s)

	fmt.Println(rw.Read())
}
func interfaceTask5() {

	// тут, на сколько я понял, я назначаю указатель на структуру, которая соответствует интерфейсу ReadWriter,
	// чтобы использовать этот указатель в функции.
	uData := &User{}
	DataReadPrint(uData, "Dialogi tet-a-tet")

}
