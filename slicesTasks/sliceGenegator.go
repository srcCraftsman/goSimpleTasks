package slicesTasks

// Генератор слайсов для использования в дальнейшем

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

// функция SliceGen запрашивает и пользователя размерность слайса и задает границы для выеличины элемента
// затем используя пакет rand наполняет слайс случайными числами
func SliceGen() []int {

	var resultSlice []int
	var elementsCoun, min, max int

	fmt.Println("\nSet N (number of elements):")
	fmt.Scan(&elementsCoun)
	fmt.Println("\nEnter the minimum value:")
	fmt.Scan(&min)
	fmt.Println("\nEnter the maximum value:")
	fmt.Scan(&max)

	// Тут задаем рандомайзер для генерации чисел в рамках наших пределов
	rand.New(rand.NewSource(uint64(time.Now().UnixNano())))

	for i := 0; i < elementsCoun; i++ {
		// генерим рандомное число для нашего элемента
		randomNum := rand.Intn(max-min+1) + min
		// добавляем его в массив
		resultSlice = append(resultSlice, randomNum)
	}
	// возращаем результат
	return resultSlice
}
