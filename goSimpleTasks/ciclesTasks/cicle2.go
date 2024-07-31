package ciclesTasks

import (
	"fmt"
)

func cicle2() {
	var n int
	s := 0
	fmt.Scan(&n)
	fmt.Println("\n")
	for i := 0; i <= n; i++ {
		s += i
	}

	fmt.Printf("Sum of all digits: %v", s)
}
