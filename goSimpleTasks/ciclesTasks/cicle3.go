package ciclesTasks

import (
	"fmt"
)

func Cicle3() {
	var k, n int

	fmt.Scan(&k, &n)
	p := k

	for i := k + 1; i <= n; i++ {
		p *= i
	}
	fmt.Println(p)
}
