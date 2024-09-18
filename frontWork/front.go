package front

import (
	"fmt"
	"goSimpleTasks/ciclesTasks"
	cSlicesTasks "goSimpleTasks/cicles_slicesTasks"
	"goSimpleTasks/funcTasks"
	"goSimpleTasks/goroutinesTasks"
	"goSimpleTasks/interfacesTasks"
	"goSimpleTasks/mapsTasks"
	"goSimpleTasks/readWriteTasks"
	"goSimpleTasks/slicesTasks"
	"goSimpleTasks/structureTasks"
)

func SetTasks() {

	var val int = 99
	for val != 0 {

		fmt.Print(`
1.  Cycles
2.  Slices
3.  Cycles and slices
4.  Maps
5.  Structure
6.  Interface
7.  Functions
8.  ReadWrite
9.  Go-rutines

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
		case val == 5:
			structureTasks.SetStructure()
			fmt.Printf("\n")
		case val == 6:
			interfacesTasks.SetInterface()
			fmt.Printf("\n")
		case val == 7:
			funcTasks.SetFunc()
			fmt.Printf("\n")
		case val == 8:
			readWriteTasks.SetReadWrite()
			fmt.Printf("\n")
		case val == 9:
			goroutinesTasks.SetGoroutine()
			fmt.Printf("\n")
		}
	}
}
