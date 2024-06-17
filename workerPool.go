// go run -race workerPool.go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

const (
	numOfWorkers = 5
	numOfTasks   = 20
)

var wg sync.WaitGroup

type Task struct {
	id int
}

func main() {

	tasks := make(chan Task, numOfWorkers)
	for i := 1; i <= numOfWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, &wg)
	}

	go workManager(tasks)
	wg.Wait()

	fmt.Printf("num of goroutines: %d\n", runtime.NumGoroutine())
}

func worker(id int, tasks <-chan Task, wg *sync.WaitGroup) {
	for task := range tasks {
		workItem := task.id
		fmt.Printf("worker id: %d, received: %d\n", id, workItem)
	}
	wg.Done()
}

func workManager(task chan<- Task) {
	for i := 1; i <= numOfTasks; i++ {
		task <- Task{id: i}
	}
	close(task)
}
