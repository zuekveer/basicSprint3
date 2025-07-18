/*
Какая временная сложность у этого кода?
for i := 0; i < n; i++ {
 for j := 0; j < n; j++ {
	  ... 
	} 
}
Сложность: О(n^2)

Дано: массив nums []int. Найди пару элементов, сумма которых равна target. Напиши эффективное решение (O(n)) и прокомментируй его сложность.
func TwoSum(nums []int, target int) (int, int, bool) {
	// ваш код
}
*/

package task2

import "fmt"

func TwoSum(nums []int, target int) (int, int, bool) {
    seen := make(map[int]int) 

    for i, num := range nums {
        complement := target - num
        if idx, ok := seen[complement]; ok {
            return idx, i, true 
        }
        seen[num] = i 
    }

    return 0, 0, false 
}

func main() {
    nums := []int{2, 7, 11, 15}
    target := 9

    i, j, found := TwoSum(nums, target)
    if found {
        fmt.Printf("Пара найдена: индексы %d и %d (числа %d и %d)\n", i, j, nums[i], nums[j])
    } else {
        fmt.Println("Пара не найдена")
    }
}