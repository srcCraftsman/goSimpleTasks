package ciclesTasks

import (
	"fmt"
)

func cicle8() {

	var n int
	fmt.Scan(&n)
	fmt.Println("\n")
	for i := 1; i <= 10; i++ {

		a := (n * i)

		fmt.Printf("%d x %d = %d\n", n, i, a)

	}
}
