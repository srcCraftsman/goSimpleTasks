package interfacesTasks

import "fmt"

//	9. Интерфейс с указателем на структуру
//  Создайте структуру Counter с полем Count (тип int) и методом Increment(), который увеличивает значение поля.
//	Определите интерфейс Incrementer с методом Increment(). Напишите функцию, которая принимает Incrementer
//  и вызывает его метод. Проверьте, что функция работает только с указателями на структуру Counter.

type Counter struct {
	Count int
}
type Incrementer interface {
	Increment()
}

func (c *Counter) Increment() {
	c.Count++
}
func plusCount(i Incrementer) {

	i.Increment()
}
func interfaceTask9() {

	c := &Counter{Count: 3}
	plusCount(c)
	fmt.Println(c.Count)
	plusCount(c)
	fmt.Println(c.Count)
	plusCount(c)
	fmt.Println(c.Count)

}
