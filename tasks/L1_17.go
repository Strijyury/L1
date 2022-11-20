package tasks

//Реализовать бинарный поиск встроенными методами языка.

import (
	"fmt"
	"log"
)

func Task17() {
	nums := []int{11, 42, 45, 66, 75, 77, 88, 90, 234}
	var target int
	fmt.Printf("Введите число из списка %v: ", nums)
	fmt.Scanln(&target)
	res := binarySearch(nums, target)
	fmt.Printf("Число %d находится по индексу %d\n", target, res)
}

//Опередяем средний элемент в слайсе. Сравниваем загаданное число со средним элементом.
//Если он меньше, рекурсивно вызываем функцию бинарного поиска на слайс до среднего значения.
//Делаем обратную операцию, если загаданное число больше
func binarySearch(nums []int, target int) int {
	flag := true

	for _, num := range nums {
		if target == num {
			flag = false
		}
	}

	if flag {
		log.Fatalf("Загаданного числа нет в списке! Выберите другое число...")
	}

	mid := len(nums) / 2
	var res int

	switch {
	case nums[mid] < target:
		res = binarySearch(nums[mid+1:], target)
		res += mid + 1
	case nums[mid] > target:
		res = binarySearch(nums[:mid], target)
	default:
		res = mid
	}

	return res
}
