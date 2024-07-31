package ciclesTasks

import (
	"fmt"
)

func cicle3() {
	var k, n int
	fmt.Println("\nSet K:")
	fmt.Scan(&k)
	fmt.Println("\nSet N:")
	fmt.Scan(&n)

	proizvedenieChisel := k
	fmt.Println("\n")

	// for i := k + 1; i <= n; i++ {

	// }

	for ; k <= n; k++ {
		proizvedenieChisel *= k
	}

	fmt.Println(proizvedenieChisel)
}
