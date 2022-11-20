package tasks

import (
	"fmt"
	"sync"
)

//Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
//во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

func Task9() {
	const n = 10

	//Создаем два канала: принимающий числа из массива и принимающий квадраты этих чисел
	numsChannel := make(chan int)
	squaresChannel := make(chan int)
	array := [10]int{2, 4, 5, 8, 11, 6, 10, 56, 75, 32}
	wg := &sync.WaitGroup{}

	//Создаем 10 воркеров
	for i := 0; i < n; i++ {
		wg.Add(1)
		go communicator(wg, numsChannel, squaresChannel)
	}

	//Записываем в канал числа из массива и закрываем его
	go func() {
		for _, num := range array {
			numsChannel <- num
		}
		close(numsChannel)
	}()

	//Здесь программа ждет, пока канал squaresChannel заполнится квадратами чисел и после закрывает его
	go func() {
		wg.Wait()
		close(squaresChannel)
	}()

	//Выводим квадраты в консоль
	for resultValue := range squaresChannel {
		fmt.Println(resultValue)
	}
}

//Реализация конвейера
func communicator(wg *sync.WaitGroup, numsChannel <-chan int, squaresChannel chan<- int) {
	defer wg.Done()
	for {
		v, ok := <-numsChannel
		if !ok {
			return
		}
		squaresChannel <- v * v
	}
}
