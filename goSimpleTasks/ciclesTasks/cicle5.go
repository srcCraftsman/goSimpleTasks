package ciclesTasks

import (
	"fmt"
)

func cicle5() {

	var n int
	p := 1
	fmt.Scan(&n)
	fmt.Println("\n")
	for i := 1; i <= n; i++ {
		p *= i
	}

	fmt.Println(p)
}
