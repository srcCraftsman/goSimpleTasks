package ciclesTasks

import (
	"fmt"
)

func cicle6() {

	var n int
	var x1, x2 int = 0, 1

	fmt.Scan(&n)
	fmt.Println("\n")
	fmt.Println("0\n1") //  я по другому не смог)))   I couldn't do anything else )))

	for x3 := 0; x3 < n; {

		x3 = x1 + x2
		x1 = x2
		x2 = x3

		fmt.Println(x3)
	}
}
