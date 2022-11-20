package tasks

//Реализовать пересечение двух неупорядоченных множеств

import (
	"fmt"
)

// Реализуем множества через map
// Если данные из слайса будут совпадать, то в мапе останется только один ключ со значением true
// Далее мы добавляем в пустой слайс те элементы, которые есть в обоих мапах
func intersection(slice1, slice2 []int) []int {
	set1 := make(map[int]bool, len(slice1))
	set2 := make(map[int]bool, len(slice2))

	var res []int

	for _, val := range slice1 {
		set1[val] = true
	}

	for _, val := range slice2 {
		set2[val] = true
	}

	for key := range set2 {
		if set1[key] {
			res = append(res, key)
		}
	}

	return res
}

func Task11() {
	set1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 5}
	set2 := []int{4, 5, 6, 8, 1, 10, 11, 0, 5}

	fmt.Println(intersection(set1, set2))
}
