package tasks

//Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
//Символы могут быть unicode.

import "fmt"

//Путем парного переприсваивания

func reverseString(s string) string {
	runes := []rune(s)
	left, right := 0, len(runes)-1

	for left <= right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	return string(runes)
}

//Реализация с defer

func reverseString2(s string) {
	for _, symb := range s {
		defer fmt.Print(string(symb))
	}
}

func Task19() {
	a := "главрыба"
	fmt.Println(a)
	fmt.Println(reverseString(a))
	reverseString2(a)
}
