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
	userPass  string
	Connected bool
	Result    string
}

func (m *MySQL) Connect() {
	fmt.Println("connecting...")
	if m.userPass == "danya:@qwerty123@/dbname" {
		m.Connected = true
		fmt.Println("Connected successfull")
	} else {
		fmt.Println("not connected")
	}
}
func (m *MySQL) Disconnect() {
	fmt.Println("disconnecting...")
	m.Connected = false
	fmt.Println("disconnected")
}

// я ничего не понимаю по этому заданию, потому что я не знаю какие должные быть значения у структуры
// и что должно вообще происходить(((
// я базы данных еще не проходил, сейчас сильно в это углубляться не хочу, пока не закрепил то что уже прошел

func (m *MySQL) Query(query string) string {
	if m.Connected {
		m.Result = fmt.Sprintf("Result: %v", query)
		return m.Result
	}
	return "NO Result"
}

func dbIn(d Database, query string) {
	d.Connect()
	result := d.Query(query)
	fmt.Println(result)
	d.Disconnect()
}

func interfaceTask10() {
	mySql := &MySQL{userPass: "danya:@qwerty123@/dbname"}
	dbIn(mySql, "123")
}

// погуглил, у чатГПТ спросил(он дичь какую-то выдал), что-то намутил, вроде работает, но хз так это должно быть или нет.
