package ciclesTasks

import (
	"fmt"
)

func cicle1() {
	var n int
	i := 0
	s := 1
	fmt.Scan(&n)
	for ; s <= n; i++ {
		s += i
	}
	fmt.Println(s)
}
