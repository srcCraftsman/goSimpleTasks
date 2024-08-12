package front

import (
	"fmt"
	"goSimpleTasks/ciclesTasks"
	cSlicesTasks "goSimpleTasks/cicles_slicesTasks"
	"goSimpleTasks/mapsTasks"
	"goSimpleTasks/slicesTasks"
)

func SetTasks() {

	var val int = 99
	for val != 0 {

		fmt.Print(`
1.  Cycles
2.  Slices
3.  Cycles and slices
4.  Maps

0.  Exit
`)
    
		fmt.Scan(&val)
		switch {
		case val == 1:
			ciclesTasks.SetCicles()
			fmt.Printf("\n")

		case val == 2:
			slicesTasks.SetSlices()
			fmt.Printf("\n")
		case val == 3:
			cSlicesTasks.Set_cSlices()
			fmt.Printf("\n")
		case val == 4:
			mapsTasks.SetMaps()
			fmt.Printf("\n")
		}
	}
}
