package ciclesTasks

import (
	"fmt"
)

func Cicle2() {
	var n int
	i := 0
	s := 0
	fmt.Scan(&n)
	fmt.Println("\n")
	for ; i <= n; i++ {
		s += i
	}

	fmt.Println(s)
}
