package ciclesTasks

import (
	"fmt"
)

func Cicle10() {

	var n int
	fmt.Scan(&n)
	var a int

	for i := 1; i <= n; i++ {
		if i%2 != 0 {
			a = i + a
		}

	}
	fmt.Println(a)
}
