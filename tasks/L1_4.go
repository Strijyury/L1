package tasks

//Реализовать постоянную запись данных в канал (главный поток).
//Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
//Необходима возможность выбора количества воркеров при старте.
//Программа должна завершаться по нажатию Ctrl+C.
//Выбрать и обосновать способ завершения работы всех воркеров.

import (
	"fmt"
	"time"
)

func Task4() {
	var workersCap int
	fmt.Printf("Введите количество воркеров : \n")

	//Считваем данные с консоли
	fmt.Scan(&workersCap)

	//Создаем канал
	writer := make(chan int)

	//Запускаем указанное количество воркеров
	for i := 0; i < workersCap; i++ {
		go worker(i, writer)
	}

	//Бесконечным циклом записываем в канал информацию
	for {
		writer <- time.Now().Second()
		time.Sleep(time.Second)
	}

}

//Читаем инфу из канала с помощью worker
func worker(name int, chanel <-chan int) {
	for {
		msg := <-chanel
		fmt.Printf("Воркер под номером %d обработал сообщение %d\n", name, msg)
	}
}
