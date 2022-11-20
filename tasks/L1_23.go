package tasks

//Удалить i-ый элемент из слайса.

import "fmt"

//Возвращаем слайс из двух слайсов (до удаленного элемента и после)
func remove(a []int, i int) []int {
	return append(a[:i], a[i+1:]...)
}

func Task23() {
	arr := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(arr)
	fmt.Println(remove(arr, 3))
}
