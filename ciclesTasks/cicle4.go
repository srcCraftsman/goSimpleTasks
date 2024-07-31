package ciclesTasks

import (
	"fmt"
)

func cicle4() {

	var n int
	fmt.Scan(&n)
	fmt.Println("\n")

	for i := 1; i <= n; i++ {
		if i%2 == 0 {

			fmt.Println(i)
		}
	}
}
