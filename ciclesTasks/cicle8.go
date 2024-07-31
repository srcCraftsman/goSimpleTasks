package ciclesTasks

import (
	"fmt"
)

func cicle8() {

	var N int
	fmt.Scan(&N)
	fmt.Println("\n")
	for i := 1; i <= 10; {

		a := (N * i)

		fmt.Printf("%d x %d = %d\n", N, i, a)

		i++
	}
}
