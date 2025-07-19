/*
Зачем использовать worker pool вместо запуска N горутин на каждый таск?
1. Балансировка и контроль ресурсов.


Реализуй worker pool из 5 воркеров, обрабатывающих задания из канала:
func StartWorkerPool(jobs &lt;-chan int, workers int) {
	// ваш код
}
*/

package task5

import (
	"fmt"
	"sync"
	"time"
)

func StartWorkerPool(jobs <-chan int, workers int) {
	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs {
				processJob(job, workerID)
			}
		}(i + 1)
	}

	wg.Wait()
}

func processJob(job int, workerID int) {
	fmt.Printf("Воркер %d начал задание %d\n", workerID, job)
	time.Sleep(time.Second)
	fmt.Printf("Воркер %d завершил задание %d\n", workerID, job)
}

func main() {
	jobs := make(chan int, 10)
	go func() {
		for i := 1; i <= 10; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	StartWorkerPool(jobs, 5)
}
