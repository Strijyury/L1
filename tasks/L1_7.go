package tasks

//Реализовать конкурентную запись данных в map.

import (
	"fmt"
	"sync"
	"time"
)

// Из-за того, что конкурентная запись в map не потокобезопасна, помимо WaitGroup нужно использовать
// Mutex для уникальной блокировки, чтобы избежать состояния гонки и небольшую задержку с Sleep, чтобы
// все элементы были записаны в map

func Task7() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	dict := make(map[int]int)

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			mu.Lock()
			dict[i] = i * i
			mu.Unlock()

			time.Sleep(time.Millisecond)
		}(i)
	}

	wg.Wait()
	fmt.Println(len(dict))
	fmt.Println(dict[1])
	fmt.Println(dict[9999])
}
