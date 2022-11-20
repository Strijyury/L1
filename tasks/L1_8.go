package tasks

//Дана переменная int64. Разработать программу, которая устанавливает i-й бит в 1 или 0.

import "fmt"

func Task8() {
	var (
		userNum int
		userPos uint
	)
	//Читаем данные с консоли

	fmt.Println("Введите число и i-ый бит, который хотите поменять: ")
	fmt.Scanln(&userNum, &userPos)

	//Проверяем, какой бит стоит на позиции, которую указал пользователь
	//Если 1, выполняем функцию clearBit, чтобы заменить его на 0
	//В противном случает, выполняется функция setBit
	//Все действия реализуются за счет битовых сдвигов
	if hasBit(userNum, userPos) {
		fmt.Println(clearBit(userNum, userPos))
	} else {
		fmt.Println(setBit(userNum, userPos))
	}
}

func hasBit(num int, pos uint) bool {
	val := num & (1 << pos)
	return (val > 0)
}

func clearBit(num int, pos uint) int {
	mask := ^(1 << pos)
	num &= mask
	return num
}
func setBit(num int, pos uint) int {
	num |= (1 << pos)
	return num
}
