package interfacesTasks

import "fmt"

// 10. Интерфейс с несколькими методами
// Определите интерфейс Database с методами Connect(), Disconnect() и Query(query string).
// Создайте структуру MySQL, которая реализует этот интерфейс.
// Напишите функцию, которая принимает Database и выполняет последовательность операций подключения, запроса и отключения.

type Database interface {
	Connect()
	Disconnect()
	Query(query string) string
}
type MySQL struct {
	ConnectionString string
	Connected        bool
	Result           string
}

func (m MySQL) Connect() {
	fmt.Println("connecting...")
	m.Connected = true
	fmt.Println("Connected successfull")
}
func (m MySQL) Disconnect() {
	fmt.Println("disconnecting...")
	m.Connected = false
	fmt.Println("disconnected")
}
func (m MySQL) Query(query string) {

}

// я ничего не понимаю по этому заданию. потому что я не знаю какие должные быть значения у структуры
// и что должно вообще происходить(((
