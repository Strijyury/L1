package tasks

//Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Реализация быстрой сортировки по принципу "разделяй и властвуй"
// Сначала слайс рандомно вбирает опорный элемент и сравнивает все оставшиеся элементы с ним
// Меняя их положение левее, если они меньше. Такием образом, все элементы слева от опорного будут
// меньше его, а спарва - больше
// Затем происходит рекурсия, дробящая слайс на две части, не включая опорный элемент и снова сортируется
// подобным образом.
// На выходе из рекурсии возвращается отсортированный слайс

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	rand.Seed(time.Now().UnixNano())
	randNum := rand.Int()
	pivot := randNum % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}

func Task16() {
	a := []int{3, 7, 2, 9, 0, 1, 6, 8, 4}
	fmt.Println(a)
	fmt.Println(quicksort(a))
	sort.Ints(a)
	fmt.Println(a)
}
