package tasks

// Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из
//массива (2,4,6,8,10) и выведет их квадраты в stdout.

import (
	"fmt"
	"sync"
)

func Task2() {
	nums := [5]int{2, 4, 6, 8, 10}

	//WaitGroup позволяет горутине main дождаться завершения всех
	//остальных горутин и только потом завершиться
	var wg sync.WaitGroup

	for _, num := range nums {

		//Добавляем go func в стек WaitGroup
		wg.Add(1)

		//Горутины выполняют вычисления конкурентно, созадвая независимые легковесные потоки
		go func(num int) {

			//Сообщаем стеку WaitGroup, что горутина выполнилась
			//Defer обеспечивает выполнение wg.Done() после вывода квадрата в консоль
			defer wg.Done()
			fmt.Println(num * num)
		}(num)
	}

	wg.Wait()
}
