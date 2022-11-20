package tasks

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

import (
	"fmt"
	"strings"
)

//Переобразовал строку в слайс и путем парного переприсваивания изменил порядок слов
//Затем снова переобразовал слайс в строку

func reverseWords(str string) string {
	words := strings.Fields(str)

	left, right := 0, len(words)-1

	for left <= right {
		words[left], words[right] = words[right], words[left]
		left++
		right--
	}

	return strings.Join(words, " ")
}

//С помощью string.Builder записал в новую переменную слова в обратном порядке

func reverseWords2(str string) string {
	words := strings.Fields(str)

	var res strings.Builder

	for i := len(words) - 1; i >= 0; i-- {
		res.WriteString(words[i])
		res.WriteString(" ")
	}

	return strings.TrimSpace(res.String())
}

func Task20() {
	s := "snow dog sun"

	fmt.Println(s)
	fmt.Println(reverseWords(s))
	fmt.Println(reverseWords2(s))
}
