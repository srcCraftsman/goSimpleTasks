package ciclesTasks

import (
	"fmt"
)

func Cicle3() {
	var k, n int
	fmt.Println("\nSet K:")
	fmt.Scan(&k)
	fmt.Println("\nSet N:")
	fmt.Scan(&n)

	p := k
	fmt.Println("\n")

	for i := k + 1; i <= n; i++ {
		p *= i
	}

	fmt.Println(p)
}
