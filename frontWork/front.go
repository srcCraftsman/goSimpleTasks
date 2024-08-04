package front

import (
	"fmt"
	"goSimpleTasks/ciclesTasks"
	cSlicesTasks "goSimpleTasks/cicles_slicesTasks"
	"goSimpleTasks/slicesTasks"
)

func SetTasks() {
	var val int = 99
	for val != 0 {
		fmt.Println(`
Set task:
1. Cicles
2. Slices
3. Cicles and Slices
0. Stop
`)
		fmt.Scan(&val)
		switch {
		case val == 1:
			ciclesTasks.SetCicles()
			fmt.Printf("\n")

		case val == 2:
			slicesTasks.SetSlices()

		case val == 3:
			cSlicesTasks.Set_cSlices()

		}
	}
}
