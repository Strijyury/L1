package tasks

//Реализовать все возможные способы остановки выполнения горутины.

import (
	"context"
	"fmt"
	"time"
)

//Остановка с помощью закрытия канала
func stopByClosing() {
	ch := make(chan int)

	go func() {
		for {
			v, ok := <-ch

			//При закрытом канале отсутствует параметр ok
			if !ok {

				//Выводим в консоль сигнал и выходим из бесконечного цикла
				fmt.Println("Stopped by closing...")
				return
			}
			fmt.Println(v)
		}
	}()
	for i := 0; i < 3; i++ {
		ch <- i
		time.Sleep(time.Second)
	}

	//Закрываем канал на том же уровне, на котором записывали в него данные
	close(ch)
}

//Остановка с помощью другого канала
func stopByChannel() {
	ch := make(chan int)
	doneCh := make(chan struct{})

	//В первый канал данные записываются сразу же каждую секунду.
	//Во второй канал данные начнут поступать через 3 секунды.
	//Благодаря Select первые 3 секунды будет выпоняться первый case
	//Затем начнется запись во второй канал. Как выполнится второй case,
	//закроется первый канал и выйдет из цикла
	go func() {
		for {
			select {
			case ch <- 1:
			case <-doneCh:
				close(ch)
				return
			}
			time.Sleep(time.Second)
		}
	}()

	go func() {
		time.Sleep(3 * time.Second)
		doneCh <- struct{}{}
	}()

	for i := range ch {
		fmt.Println("Received: ", i)
	}

	fmt.Println("Stopped by channel...")
}

//Остановка с помощью context.WithCancel
func stoppedByWithCancel() {
	ch := make(chan struct{})

	//Контекст с отменой позволит завершить программу сразу же,
	//как реализуется функция cancel
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			//по дефолту в консоль выводится строка foo каждую секунду
			//в другой горутине спустя 3 секунды запустится функция cancel
			//которая передаст данные в case о завершении контекста и программа завершится
			select {
			case <-ctx.Done():
				//В канал помещается любая пустая структура, чтобы позже прочитать из канала данные
				//это делается для того, чтобы основная функция дождалась выполнения горутин в ней
				//и после завершилась
				ch <- struct{}{}
				return
			default:
				fmt.Println("foo")
			}

			time.Sleep(time.Second)
		}
	}(ctx)

	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()

	<-ch
	fmt.Println("Stopped by WithCancel...")
}

func stoppedByWithTimeout() {
	ch := make(chan int)

	//WithTimeout задает время до завершения программы (3 секунды)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	//В канал бесконечно записывается инфа
	go func() {
		for {
			ch <- time.Now().Second()
			time.Sleep(time.Second)
		}
	}()

	//До активации остановки спустя 3 секунды, в консоль будут выводиться данные из канала
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stopped with TimeOut...")
			return
		case msg := <-ch:
			fmt.Println(msg)
		}
	}
}

//Остановка с помощью Deadline
func stoppedByWithDeadline() {
	ch := make(chan int)
	d := time.Now().Add(3 * time.Second)

	//Реализация схожа с WithTimeout: программа завершиться спустя 3 секунды после запуска программы
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	go func() {
		for {
			ch <- time.Now().Second()
			time.Sleep(time.Second)
		}
	}()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stopped by Deadline...")
			return
		case msg := <-ch:
			fmt.Println(msg)
		}
	}
}

func Task6() {
	stopByClosing()
	stopByChannel()
	stoppedByWithCancel()
	stoppedByWithTimeout()
	stoppedByWithDeadline()
}
