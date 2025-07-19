/*
Как с помощью семафора ограничить количество одновременно выполняющихся горутин?
Исп.буферизированный канал.

Запусти все задачи, но не более limit одновременно:
func RunLimited(tasks []func(), limit int) {
	// ваш код
}
*/

package task5

import (
	"fmt"
	"sync"
	"time"
)

func RunLimited(tasks []func(), limit int) {
	semaphore := make(chan struct{}, limit)
	var wg sync.WaitGroup

	for _, task := range tasks {
		wg.Add(1)
		go func(f func()) {
			defer wg.Done()
			semaphore <- struct{}{}
			defer func() { <-semaphore }()
			f()
		}(task)
	}

	wg.Wait()
}

func main() {
	tasks := []func(){
		func() { time.Sleep(1 * time.Second); fmt.Println("Task 1") },
		func() { time.Sleep(2 * time.Second); fmt.Println("Task 2") },
		func() { time.Sleep(1 * time.Second); fmt.Println("Task 3") },
		func() { time.Sleep(3 * time.Second); fmt.Println("Task 4") },
	}

	RunLimited(tasks, 2)
}
