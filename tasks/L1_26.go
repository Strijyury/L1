package tasks

//Разработать программу, которая проверяет, что все символы в строке уникальные
//(true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

import (
	"fmt"
	"strings"
)

//Приводим строку в нижнему регистру, реализуем множество с помощью мапы
//Если в мапе столько же элементов, сколько и в строке -> все символы уникальны

func isUnique(s string) bool {
	dict := make(map[int32]bool)
	res := false
	s = strings.ToLower(s)

	for _, str := range s {
		dict[str] = true
	}

	if len(dict) == len(s) {
		res = true
	}

	return res
}

func Task26() {
	strs := []string{"abcd", "abCdefAaf", "aabcd", "aAbcd", "AbCdE"}

	for i := range strs {
		fmt.Printf("%s is %v\n", strs[i], isUnique(strs[i]))
	}
}
