package tasks

//Реализовать собственную функцию sleep.

import (
	"fmt"
	"time"
)

//Как только считаются данные с канала time.After, программа заснет на заданное время

var sec time.Duration

func sleep(d time.Duration) {
	<-time.After(d)
}

func Task25() {
	fmt.Println("Enter the number of seconds the program will sleep for: ")
	fmt.Scan(&sec)

	start := time.Now()

	fmt.Println("Program started")
	fmt.Printf("Sleep for %d seconds\n", sec)

	sleep(sec * time.Second)

	fmt.Println("Program fineshed with time: ", time.Since(start))
}
