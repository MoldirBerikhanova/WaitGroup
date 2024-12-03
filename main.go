package main

import (
	"fmt"
	"sync"
	"time"
)

func runTask(id int, seconds int, wg *sync.WaitGroup) {

	defer wg.Done()
	fmt.Printf("Task %d started, will run for %d seconds\n", id, seconds)
	time.Sleep(time.Duration(seconds) * time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	var wg sync.WaitGroup

	wg.Add(5)

	go runTask(1, 1, &wg)
	go runTask(2, 5, &wg)
	go runTask(3, 10, &wg)
	go runTask(4, 3, &wg)
	go runTask(5, 7, &wg)

	// for i := 1; i <= 5; i++ {

	// 	go func() {
	// 		defer wg.Done()
	// 		runTask(i, j)
	// 	}()
	// }

	wg.Wait()

}

//func runTaskFixed(taskId int, seconds int) {

// go runTask(1, 1)
// go runTask(2, 5)
// go runTask(3, 10)
// go runTask(4, 3)
// go runTask(5, 7)

// Таск #5 выполняется 7 секунд...
// Таск #3 выполняется 10 секунд...
// Таск #4 выполняется 3 секунд...
// Таск #1 выполняется 1 секунд...
// Таск #2 выполняется 5 секунд...
// Таск #1 завершился
// Таск #4 завершился
// Таск #2 завершился
// Таск #5 завершился
// Таск #3 завершился
// Все таски выполнены

// fmt.Printf("Таск #%d выполняется %d секунд...\n", taskId, seconds)
// time.Sleep(time.Duration(seconds) * time.Second)
// fmt.Printf("Таск #%d завершился\n", taskId)
