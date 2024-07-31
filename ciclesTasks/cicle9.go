package ciclesTasks

import (
	"fmt"
)

func cicle9() {

	var n int
	fmt.Scan(&n)
	var a int
	fmt.Println("\n")
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			a = i + a
		}

	}

	fmt.Println(a)
}
