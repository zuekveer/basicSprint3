/*
Почему бинарный поиск работает только на отсортированных массивах?
Алгоритм делит массив 50/50, справа значения должны быть больше. Для этого массив д.б. отсортирован.

Реализуй функцию BinarySearch, возвращающую индекс элемента в отсортированном массиве или -1, если не найден:
go
КопироватьРедактировать
func BinarySearch(arr []int, target int) int {
	// ваш код
}
*/

package task1

import "fmt"

func BinarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15}

	target := 7

	index := BinarySearch(arr, target)

	if index != -1 {
		fmt.Printf("Элемент %d найден на позиции %d\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден в массиве\n", target)
	}

	target = 8
	index = BinarySearch(arr, target)

	if index != -1 {
		fmt.Printf("Элемент %d найден на позиции %d\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден в массиве\n", target)
	}
}
