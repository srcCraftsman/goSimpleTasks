package slicesTasks

// Функция сдвига элементов слайса на N для использования в дальнейшем

func ShiftSlice(sliceName []int, posNum int) []int {
	var i int = posNum % len(sliceName)

	if i >= 0 {
		return append(sliceName[len(sliceName)-i:], sliceName[0:len(sliceName)-1]...)
	} else {
		i = i * -1
		return append(sliceName[i:], sliceName[0:i]...)
	}
}
