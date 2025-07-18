/*
В чём разница между стабильной и нестабильной сортировкой?
Стабильная сортировка сохраняет относительный порядок элементов с одинаковыми ключами.

Реализуй сортировку пузырьком для слайса []int:
func BubbleSort(arr []int) {
	// ваш код
}
*/

package task1

import "fmt"

func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("befor sorting:", arr)

	BubbleSort(arr)
	fmt.Println("after sorting:", arr)
}
