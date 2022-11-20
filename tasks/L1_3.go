package tasks

//Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов
//с использованием конкурентных вычислений.

import (
	"fmt"
	"sync"
)

func Task3() {
	var res int

	//WaitGroup позволяет горутине main дождаться завершения всех
	//остальных горутин и только потом завершиться
	var wg sync.WaitGroup
	nums := [5]int{2, 4, 6, 8, 10}

	for _, num := range nums {

		//Добавляем go func в стек WaitGroup
		wg.Add(1)

		//Горутины выполняют вычисления конкурентно, созадвая независимые легковесные потоки
		go func(num int) {

			//Сообщаем стеку WaitGroup, что горутина выполнилась
			//Defer обеспечивает выполнение wg.Done() после вывода квадрата в консоль
			defer wg.Done()
			res += num * num
		}(num)
	}
	wg.Wait()
	fmt.Println(res)
}
