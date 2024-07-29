package ciclesTasks

import (
	"fmt"
)

func cicle3() {
	var n int

	p := 1
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		p *= i
	}
	fmt.Println(p)
}
