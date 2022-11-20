package tasks

//Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
//По завершению программа должна выводить итоговое значение счетчика.

import (
	"fmt"
	"sync"
)

// Структура Counter с Mutex для исключения состояния гонки
type Counter struct {
	sync.Mutex
	counter int
}

// Конструктор для Counter
func NewCounter() *Counter {
	return &Counter{}
}

// Inc увеличивает Counter на один с Mutex
func (c *Counter) Inc() {
	c.Lock()
	c.counter++
	c.Unlock()
}

// String выводить значение в консоль
func (c *Counter) String() string {
	return fmt.Sprint(c.counter)
}

func Task18() {
	c := NewCounter()
	var wg sync.WaitGroup

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}

	wg.Wait()
	fmt.Println(c)
}
